package postgres

// type ExerciseStorage struct {
// 	dbpool *pgxpool.Pool
// }

// func NewExerciseStorage(conn *pgxpool.Pool) *ExerciseStorage {
// 	return &ExerciseStorage{dbpool: conn}
// }

// func (s *ExerciseStorage) CreateMultipleChoiceExercise(
// 	e exercise.CreateMultipleChoiceExerciseCommand,
// ) (*exercise.MultipleChoiceExercise, error) {
// 	id, err := uuid.NewRandom()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
// 	}

// 	row := s.dbpool.QueryRow(context.Background(), `
// 		INSERT INTO exercise (id, type, question, choices, answer)
// 		VALUES ($1, $2, $3, $4, $5)
// 		RETURNING *
// 	// `, id, exercise.TypeMultipleChoice, e.Question, e.Choices, e.Answer)

// 	entity, err := mapToEntity(row)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to map row to entity: %w", err)
// 	}

// 	return mapToMultipleChoiceExercise(*entity), nil
// }

// func (s *ExerciseStorage) CreateFillInTheBlankExercise(
// 	e exercise.CreateFillInTheBlankExerciseCommand,
// ) (*exercise.FillInTheBlankExercise, error) {
// 	id, err := uuid.NewRandom()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
// 	}

// 	row := s.dbpool.QueryRow(context.Background(), `
// 		INSERT INTO exercise (id, type, question, answer)
// 		VALUES ($1, $2, $3, $4)
// 		RETURNING *
// 	`, id, exercise.TypeFillInTheBlank, e.Question, e.Answer)

// 	entity, err := mapToEntity(row)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to map row to entity: %w", err)
// 	}

// 	return mapToFillInTheBlankExercise(*entity), nil
// }

// func (s *ExerciseStorage) CreateSentenceCorrectionExercise(
// 	e exercise.CreateSentenceCorrectionExerciseCommand,
// ) (*exercise.SentenceCorrectionExercise, error) {
// 	id, err := uuid.NewRandom()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
// 	}

// 	row := s.dbpool.QueryRow(context.Background(), `
// 		INSERT INTO exercise (id, type, sentence, corrected_sentence)
// 		VALUES ($1, $2, $3, $4)
// 		RETURNING *
// 	`, id, exercise.TypeSentenceCorrection, e.Sentence, e.CorrectedSentence)

// 	entity, err := mapToEntity(row)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to map row to entity: %w", err)
// 	}

// 	return mapToSentenceCorrectionExercise(*entity), nil
// }

// func mapToEntity(row pgx.Row) (*ExerciseEntity, error) {
// 	var entity ExerciseEntity
// 	err := row.Scan(
// 		&entity.ID,
// 		&entity.CreatedAt,
// 		&entity.UpdatedAt,
// 		&entity.Type,
// 		&entity.Question,
// 		&entity.Choices,
// 		&entity.Answer,
// 		&entity.Sentence,
// 		&entity.CorrectedSentence,
// 	)
// 	return &entity, err
// }

// func (s *ExerciseStorage) Find() ([]exercise.Exercise, error) {
// 	rows, err := s.dbpool.Query(context.Background(), `
// 		SELECT *
// 		FROM exercise
// 	`)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to query exercise table: %w", err)
// 	}
// 	defer rows.Close()

// 	entities := make([]ExerciseEntity, 0)
// 	for rows.Next() {
// 		entity, err := mapToEntity(rows)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to map row to entity: %w", err)
// 		}
// 		entities = append(entities, *entity)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, fmt.Errorf("failed to read exercise: %w", err)
// 	}

// 	return mapToExercises(entities)
// }

// func (s *ExerciseStorage) FindByID(id string) (exercise.Exercise, error) {
// 	uuid, err := uuid.Parse(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse id as uuid: %w", err)
// 	}

// 	row := s.dbpool.QueryRow(context.Background(), `
// 		SELECT *
// 		FROM exercise
// 		WHERE id = $1
// 	`, uuid)

// 	entity, err := mapToEntity(row)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to map row to entity: %w", err)
// 	}

// 	return mapToExercise(*entity)
// }

// func mapToMultipleChoiceExercise(entity ExerciseEntity) *exercise.MultipleChoiceExercise {
// 	e := exercise.NewMultipleChoiceExercise(
// 		entity.ID.String(),
// 		entity.CreatedAt,
// 		entity.UpdatedAt,
// 		*entity.Question,
// 		*entity.Choices,
// 		*entity.Answer,
// 	)
// 	return &e
// }

// func mapToFillInTheBlankExercise(entity ExerciseEntity) *exercise.FillInTheBlankExercise {
// 	e := exercise.NewFillInTheBlankExercise(
// 		entity.ID.String(),
// 		entity.CreatedAt,
// 		entity.UpdatedAt,
// 		*entity.Question,
// 		*entity.Answer,
// 	)
// 	return &e
// }

// func mapToSentenceCorrectionExercise(entity ExerciseEntity) *exercise.SentenceCorrectionExercise {
// 	e := exercise.NewSentenceCorrectionExercise(
// 		entity.ID.String(),
// 		entity.CreatedAt,
// 		entity.UpdatedAt,
// 		*entity.Sentence,
// 		*entity.CorrectedSentence,
// 	)
// 	return &e
// }

// func mapToExercises(entities []ExerciseEntity) ([]exercise.Exercise, error) {
// 	exercises := make([]exercise.Exercise, 0)
// 	for _, entity := range entities {
// 		exercise, err := mapToExercise(entity)
// 		if err != nil {
// 			return nil, err
// 		}
// 		exercises = append(exercises, exercise)
// 	}
// 	return exercises, nil
// }

// func mapToExercise(entity ExerciseEntity) (exercise.Exercise, error) {
// 	switch entity.Type {
// 	case exercise.TypeMultipleChoice:
// 		return mapToMultipleChoiceExercise(entity), nil
// 	case exercise.TypeFillInTheBlank:
// 		return mapToFillInTheBlankExercise(entity), nil
// 	case exercise.TypeSentenceCorrection:
// 		return mapToSentenceCorrectionExercise(entity), nil
// 	default:
// 		return nil, fmt.Errorf("unknown exercise type: %s", entity.Type)
// 	}
// }

// type ExerciseEntity struct {
// 	ID        uuid.UUID
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	Type      string

// 	Question *string
// 	Choices  *[]string
// 	Answer   *string

// 	Sentence          *string
// 	CorrectedSentence *string
// }
