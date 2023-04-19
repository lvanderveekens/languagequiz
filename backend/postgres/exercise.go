package postgres

import (
	"context"

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
	var e Exercise
	err := es.dbpool.QueryRow(context.Background(), `
		INSERT INTO "exercise" ("name") 
		VALUES ('foo') 
		RETURNING *
	`).Scan(&e)
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
