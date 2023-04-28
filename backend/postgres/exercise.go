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
	`, id, "multipleChoice", e.Question, e.Options, e.CorrectOption)

	entity, err := mapToEntity(row)
	if err != nil {
		return nil, fmt.Errorf("failed to map row to entity: %w", err)
	}

	return mapToMultipleChoiceExercise(*entity), nil
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
		entity.Question,
		entity.Options,
		entity.CorrectOption,
	)
	return &e
}

func mapToExercises(entities []Exercise) ([]any, error) {
	exercises := make([]any, 0)
	for _, entity := range entities {
		var exercise any
		switch entity.Type {
		case "multipleChoice":
			exercise = *mapToMultipleChoiceExercise(entity)
		default:
			return nil, fmt.Errorf("unknown exercise type: %s", entity.Type)
		}
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

type Exercise struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Type      string

	// multiple choice fields
	Question      string
	Options       []string
	CorrectOption string
}
