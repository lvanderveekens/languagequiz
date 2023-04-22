package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lvanderveekens/language-resources/exercise"
)

type ExerciseStorage struct {
	dbpool *pgxpool.Pool
}

func NewExerciseStorage(conn *pgxpool.Pool) *ExerciseStorage {
	return &ExerciseStorage{dbpool: conn}
}

func (s *ExerciseStorage) Create() (*exercise.Exercise, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
	}

	row := s.dbpool.QueryRow(context.Background(), `
		INSERT INTO exercise (id) 
		VALUES ($1) 
		RETURNING *
	`, id)

	entity, err := s.mapToEntity(row)
	if err != nil {
		return nil, fmt.Errorf("failed to map row to entity: %w", err)
	}

	return s.mapToDomainObject(*entity), nil
}

func (s *ExerciseStorage) mapToEntity(row pgx.Row) (*Exercise, error) {
	var entity Exercise
	err := row.Scan(&entity.ID, &entity.CreatedAt, &entity.UpdatedAt)
	return &entity, err
}

func (s *ExerciseStorage) Find() ([]exercise.Exercise, error) {
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
		entity, err := s.mapToEntity(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to map row to entity: %w", err)
		}
		entities = append(entities, *entity)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read exercise: %w", err)
	}

	return s.mapToDomainObjects(entities), nil
}

func (s *ExerciseStorage) mapToDomainObject(entity Exercise) *exercise.Exercise {
	domainObject := exercise.New(entity.ID.String(), entity.CreatedAt, entity.UpdatedAt)
	return &domainObject
}

func (s *ExerciseStorage) mapToDomainObjects(entities []Exercise) []exercise.Exercise {
	domainObjects := make([]exercise.Exercise, 0)
	for _, entity := range entities {
		domainObjects = append(domainObjects, *s.mapToDomainObject(entity))
	}
	return domainObjects
}

type Exercise struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
