package quiz

type Storage interface {
	CreateQuiz(c CreateQuizCommand) (*Quiz, error)
}
