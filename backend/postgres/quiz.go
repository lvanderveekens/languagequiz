package postgres

import (
	"context"
	"fmt"
	"time"

	"languagequiz/quiz"
	"languagequiz/quiz/exercise"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QuizStorage struct {
	dbpool *pgxpool.Pool
}

func NewQuizStorage(conn *pgxpool.Pool) *QuizStorage {
	return &QuizStorage{dbpool: conn}
}

func (s *QuizStorage) FindQuizzes() ([]quiz.Quiz, error) {
	rows, err := s.dbpool.Query(context.Background(), `
		SELECT *
		FROM quiz
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query quiz table: %w", err)
	}
	defer rows.Close()

	quizEntities := make([]QuizEntity, 0)
	for rows.Next() {
		quizEntity, err := mapToQuizEntity(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to map row to quiz entity: %w", err)
		}
		quizEntities = append(quizEntities, *quizEntity)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read quiz table rows: %w", err)
	}

	quizzes := make([]quiz.Quiz, 0)
	for _, quizEntity := range quizEntities {
		quizSectionEntities, err := s.findQuizSectionEntities(quizEntity.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to find quiz section entities: %w", err)
		}

		quizSectionIDs := make([]uuid.UUID, 0)
		for _, quizSectionEntity := range quizSectionEntities {
			quizSectionIDs = append(quizSectionIDs, quizSectionEntity.ID)
		}

		exerciseEntitiesBySectionID, err := s.findExerciseEntitiesBySectionID(quizSectionIDs)
		if err != nil {
			return nil, fmt.Errorf("failed to find exercise entities: %w", err)
		}

		quiz, err := mapToQuiz(quizEntity, quizSectionEntities, exerciseEntitiesBySectionID)
		if err != nil {
			return nil, fmt.Errorf("failed to map entity to quiz: %w", err)
		}
		quizzes = append(quizzes, *quiz)
	}

	return quizzes, nil
}

func (s *QuizStorage) findExerciseEntitiesBySectionID(quizSectionIDs []uuid.UUID) (map[string][]ExerciseEntity, error) {
	rows, err := s.dbpool.Query(context.Background(), `
		SELECT *
		FROM exercise
		WHERE quiz_section_id = ANY ($1)
	`, quizSectionIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to query exercise table: %w", err)
	}
	defer rows.Close()

	exerciseEntities := make([]ExerciseEntity, 0)
	for rows.Next() {
		exerciseEntity, err := mapToExerciseEntity(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to map row to exercise entity: %w", err)
		}
		exerciseEntities = append(exerciseEntities, *exerciseEntity)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read quiz table rows: %w", err)
	}

	exerciseEntitiesBySectionID := make(map[string][]ExerciseEntity)
	for _, exerciseEntity := range exerciseEntities {
		exerciseEntitiesBySectionID[exerciseEntity.QuizSectionID.String()] = append(
			exerciseEntitiesBySectionID[exerciseEntity.QuizSectionID.String()],
			exerciseEntity,
		)
	}

	return exerciseEntitiesBySectionID, nil
}

func (s *QuizStorage) findQuizSectionEntities(quizID uuid.UUID) ([]QuizSectionEntity, error) {
	rows, err := s.dbpool.Query(context.Background(), `
		SELECT *
		FROM quiz_section
		WHERE quiz_id = $1
	`, quizID)
	if err != nil {
		return nil, fmt.Errorf("failed to query quiz_section table: %w", err)
	}
	defer rows.Close()

	quizSectionEntities := make([]QuizSectionEntity, 0)
	for rows.Next() {
		quizSectionEntity, err := mapToQuizSectionEntity(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to map row to quiz section entity: %w", err)
		}
		quizSectionEntities = append(quizSectionEntities, *quizSectionEntity)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read quiz table rows: %w", err)
	}

	return quizSectionEntities, nil
}

func (s *QuizStorage) CreateQuiz(cmd quiz.CreateQuizCommand) (*quiz.Quiz, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
	}

	tx, err := s.dbpool.Begin(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(context.Background())

	quizEntity, err := mapToQuizEntity(tx.QueryRow(context.Background(), `
		INSERT INTO quiz (id, name)
		VALUES ($1, $2)
		RETURNING *
	`, id, cmd.Name))
	if err != nil {
		return nil, fmt.Errorf("failed to insert quiz: %w", err)
	}

	quizSectionEntities := make([]QuizSectionEntity, 0)
	exerciseEntitiesBySectionID := make(map[string][]ExerciseEntity)
	for _, createSectionCommand := range cmd.Sections {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, fmt.Errorf("failed to generate new UUID: %w", err)
		}

		quizSectionEntity, err := mapToQuizSectionEntity(tx.QueryRow(context.Background(), `
			INSERT INTO quiz_section (id, quiz_id, name)
			VALUES ($1, $2, $3)
			RETURNING *
		`, id, quizEntity.ID, createSectionCommand.Name))
		if err != nil {
			return nil, fmt.Errorf("failed to insert quiz section: %w", err)
		}
		quizSectionEntities = append(quizSectionEntities, *quizSectionEntity)

		for _, createExerciseCommand := range createSectionCommand.Exercises {
			var exerciseEntity *ExerciseEntity
			var err error
			switch createExerciseCommand := createExerciseCommand.(type) {
			case exercise.CreateMultipleChoiceExerciseCommand:
				exerciseEntity, err = insertMultipleChoiceExercise(tx, createExerciseCommand, quizSectionEntity.ID)
			case exercise.CreateFillInTheBlankExerciseCommand:
				exerciseEntity, err = insertFillInTheBlankExercise(tx, createExerciseCommand, quizSectionEntity.ID)
			case exercise.CreateSentenceCorrectionExerciseCommand:
				exerciseEntity, err = insertSentenceCorrectionExercise(tx, createExerciseCommand, quizSectionEntity.ID)
			default:
				return nil, fmt.Errorf("unknown exercise type: %T", createExerciseCommand)
			}
			if err != nil {
				return nil, fmt.Errorf("failed to insert exercise: %w", err)
			}
			exerciseEntitiesBySectionID[quizSectionEntity.ID.String()] =
				append(exerciseEntitiesBySectionID[quizSectionEntity.ID.String()], *exerciseEntity)
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return mapToQuiz(*quizEntity, quizSectionEntities, exerciseEntitiesBySectionID)
}

func insertMultipleChoiceExercise(
	tx pgx.Tx,
	cmd exercise.CreateMultipleChoiceExerciseCommand,
	quizSectionId uuid.UUID,
) (*ExerciseEntity, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
	}

	row := tx.QueryRow(context.Background(), `
		INSERT INTO exercise (id, quiz_section_id, type, question, choices, answer) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING *
	`, id, quizSectionId, exercise.TypeMultipleChoice, cmd.Question, cmd.Choices, cmd.Answer)

	return mapToExerciseEntity(row)
}

func insertFillInTheBlankExercise(
	tx pgx.Tx,
	cmd exercise.CreateFillInTheBlankExerciseCommand,
	quizSectionId uuid.UUID,
) (*ExerciseEntity, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
	}

	row := tx.QueryRow(context.Background(), `
		INSERT INTO exercise (id, quiz_section_id, type, question, answer) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING *
	`, id, quizSectionId, exercise.TypeFillInTheBlank, cmd.Question, cmd.Answer)

	return mapToExerciseEntity(row)
}

func insertSentenceCorrectionExercise(
	tx pgx.Tx,
	cmd exercise.CreateSentenceCorrectionExerciseCommand,
	quizSectionId uuid.UUID,
) (*ExerciseEntity, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
	}

	row := tx.QueryRow(context.Background(), `
		INSERT INTO exercise (id, quiz_section_id, type, sentence, corrected_sentence) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING *
	`, id, quizSectionId, exercise.TypeSentenceCorrection, cmd.Sentence, cmd.CorrectedSentence)

	return mapToExerciseEntity(row)
}

func mapToQuizEntity(row pgx.Row) (*QuizEntity, error) {
	var entity QuizEntity
	err := row.Scan(
		&entity.ID,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.Name,
	)
	return &entity, err
}

func mapToQuizSectionEntity(row pgx.Row) (*QuizSectionEntity, error) {
	var entity QuizSectionEntity
	err := row.Scan(
		&entity.ID,
		&entity.QuizID,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.Name,
	)
	return &entity, err
}

func mapToExerciseEntity(row pgx.Row) (*ExerciseEntity, error) {
	var entity ExerciseEntity
	err := row.Scan(
		&entity.ID,
		&entity.QuizSectionID,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.Type,
		&entity.Question,
		&entity.Choices,
		&entity.Answer,
		&entity.Sentence,
		&entity.CorrectedSentence,
	)
	return &entity, err
}

func mapToQuiz(
	quizEntity QuizEntity,
	quizSectionEntities []QuizSectionEntity,
	exerciseEntitiesBySectionID map[string][]ExerciseEntity,
) (*quiz.Quiz, error) {

	sections := make([]quiz.Section, 0)
	for _, quizSectionEntity := range quizSectionEntities {
		exerciseEntities := exerciseEntitiesBySectionID[quizSectionEntity.ID.String()]

		exercises, err := mapToExercises(exerciseEntities)
		if err != nil {
			return nil, err
		}

		section := quiz.NewSection(quizSectionEntity.Name, exercises)
		sections = append(sections, section)
	}

	quiz := quiz.New(
		quizEntity.ID.String(),
		quizEntity.CreatedAt,
		quizEntity.UpdatedAt,
		quizEntity.Name,
		sections,
	)

	return &quiz, nil
}

type QuizEntity struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

type QuizSectionEntity struct {
	ID        uuid.UUID
	QuizID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

type ExerciseEntity struct {
	ID            uuid.UUID
	QuizSectionID uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Type          string

	Question *string
	Choices  *[]string
	Answer   *string

	Sentence          *string
	CorrectedSentence *string
}

func mapToExercises(entities []ExerciseEntity) ([]exercise.Exercise, error) {
	exercises := make([]exercise.Exercise, 0)
	for _, entity := range entities {
		exercise, err := mapToExercise(entity)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

func mapToExercise(entity ExerciseEntity) (exercise.Exercise, error) {
	switch entity.Type {
	case exercise.TypeMultipleChoice:
		return mapToMultipleChoiceExercise(entity), nil
	case exercise.TypeFillInTheBlank:
		return mapToFillInTheBlankExercise(entity), nil
	case exercise.TypeSentenceCorrection:
		return mapToSentenceCorrectionExercise(entity), nil
	default:
		return nil, fmt.Errorf("unknown exercise type: %s", entity.Type)
	}
}

func mapToMultipleChoiceExercise(entity ExerciseEntity) *exercise.MultipleChoiceExercise {
	e := exercise.NewMultipleChoiceExercise(
		entity.ID.String(),
		entity.CreatedAt,
		entity.UpdatedAt,
		*entity.Question,
		*entity.Choices,
		*entity.Answer,
	)
	return &e
}

func mapToFillInTheBlankExercise(entity ExerciseEntity) *exercise.FillInTheBlankExercise {
	e := exercise.NewFillInTheBlankExercise(
		entity.ID.String(),
		entity.CreatedAt,
		entity.UpdatedAt,
		*entity.Question,
		*entity.Answer,
	)
	return &e
}

func mapToSentenceCorrectionExercise(entity ExerciseEntity) *exercise.SentenceCorrectionExercise {
	e := exercise.NewSentenceCorrectionExercise(
		entity.ID.String(),
		entity.CreatedAt,
		entity.UpdatedAt,
		*entity.Sentence,
		*entity.CorrectedSentence,
	)
	return &e
}
