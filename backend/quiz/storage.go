package quiz

type Storage interface {
	CreateQuiz(cmd CreateQuizCommand) (*Quiz, error)
}
