package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lvanderveekens/language-resources/exercise"
)

type ExerciseStorage struct {
	dbpool *pgxpool.Pool
}

func NewExerciseStorage(conn *pgxpool.Pool) *ExerciseStorage {
	return &ExerciseStorage{dbpool: conn}
}

func (es *ExerciseStorage) CreateExercise() (*exercise.Exercise, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	var e Exercise
	err = es.dbpool.QueryRow(context.Background(), `
		INSERT INTO "exercise" ("id") 
		VALUES ($1) 
		RETURNING *
	`, id).Scan(&e)
	if err != nil {
		return nil, err
	}

	return mapToDomainObject(e), nil
}

func mapToDomainObject(e Exercise) *exercise.Exercise {
	return nil
}

type Exercise struct {
}
