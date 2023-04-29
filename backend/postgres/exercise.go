package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lvanderveekens/testmaker/exercise"
)

type ExerciseStorage struct {
	dbpool *pgxpool.Pool
}

func NewExerciseStorage(conn *pgxpool.Pool) *ExerciseStorage {
	return &ExerciseStorage{dbpool: conn}
}

func (s *ExerciseStorage) CreateMultipleChoiceExercise(
	e exercise.CreateMultipleChoiceExerciseCommand,
) (*exercise.MultipleChoiceExercise, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
	}

	row := s.dbpool.QueryRow(context.Background(), `
		INSERT INTO exercise (id, type, question, options, correct_option) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING *
	`, id, exercise.TypeMultipleChoice, e.Question, e.Options, e.CorrectOption)

	entity, err := mapToEntity(row)
	if err != nil {
		return nil, fmt.Errorf("failed to map row to entity: %w", err)
	}

	return mapToMultipleChoiceExercise(*entity), nil
}

func (s *ExerciseStorage) CreateCompleteTheSentenceExercise(
	e exercise.CreateCompleteTheSentenceExerciseCommand,
) (*exercise.CompleteTheSentenceExercise, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
	}

	row := s.dbpool.QueryRow(context.Background(), `
		INSERT INTO exercise (id, type, before_gap, gap, after_gap) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING *
	`, id, exercise.TypeCompleteTheSentence, e.BeforeGap, e.Gap, e.AfterGap)

	entity, err := mapToEntity(row)
	if err != nil {
		return nil, fmt.Errorf("failed to map row to entity: %w", err)
	}

	return mapToCompleteTheSentenceExercise(*entity), nil
}

func mapToEntity(row pgx.Row) (*Exercise, error) {
	var entity Exercise
	err := row.Scan(
		&entity.ID,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.Type,
		&entity.Question,
		&entity.Options,
		&entity.CorrectOption,
		&entity.BeforeGap,
		&entity.Gap,
		&entity.AfterGap,
	)
	return &entity, err
}

func (s *ExerciseStorage) Find() ([]any, error) {
	rows, err := s.dbpool.Query(context.Background(), `
		SELECT *
		FROM exercise
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query exercise table: %w", err)
	}
	defer rows.Close()

	entities := make([]Exercise, 0)
	for rows.Next() {
		entity, err := mapToEntity(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to map row to entity: %w", err)
		}
		entities = append(entities, *entity)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read exercise: %w", err)
	}

	return mapToExercises(entities)
}

func mapToMultipleChoiceExercise(entity Exercise) *exercise.MultipleChoiceExercise {
	e := exercise.NewMultipleChoiceExercise(
		exercise.New(entity.ID.String(), entity.CreatedAt, entity.UpdatedAt),
		*entity.Question,
		*entity.Options,
		*entity.CorrectOption,
	)
	return &e
}

func mapToCompleteTheSentenceExercise(entity Exercise) *exercise.CompleteTheSentenceExercise {
	e := exercise.NewCompleteTheSentenceExercise(
		exercise.New(entity.ID.String(), entity.CreatedAt, entity.UpdatedAt),
		*entity.BeforeGap,
		*entity.Gap,
		*entity.AfterGap,
	)
	return &e
}

func mapToExercises(entities []Exercise) ([]any, error) {
	exercises := make([]any, 0)
	for _, entity := range entities {
		exercise, err := mapToExercise(entity)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

func mapToExercise(entity Exercise) (any, error) {
	switch entity.Type {
	case exercise.TypeMultipleChoice:
		return *mapToMultipleChoiceExercise(entity), nil
	case exercise.TypeCompleteTheSentence:
		return *mapToCompleteTheSentenceExercise(entity), nil
	default:
		return nil, fmt.Errorf("unknown exercise type: %s", entity.Type)
	}
}

type Exercise struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Type      string

	// multiple choice fields
	Question      *string
	Options       *[]string
	CorrectOption *string

	// complete the sentence fields
	BeforeGap *string
	Gap       *string
	AfterGap  *string
}
