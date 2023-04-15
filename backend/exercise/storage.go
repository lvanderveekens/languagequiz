package exercise

type Storage interface {
	CreateExercise() (*Exercise, error)
}
