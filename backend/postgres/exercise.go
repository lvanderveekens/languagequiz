package postgres

import (
	"github.com/jackc/pgx/v4"
	"github.com/lvanderveekens/language-resources/exercise"
)

type ExerciseStorage struct {
	conn *pgx.Conn
}

func NewExerciseStorage(conn *pgx.Conn) *ExerciseStorage {
	return &ExerciseStorage{conn: conn}
}

func (s *ExerciseStorage) CreateExercise() (*exercise.Exercise, error) {
	panic("TODO")
}
