package exercise

type Storage interface {
	Create() (*Exercise, error)
	Find() ([]Exercise, error)
}
