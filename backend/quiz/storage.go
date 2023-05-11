package quiz

type Storage interface {
	FindQuizzes() ([]Quiz, error)
	CreateQuiz(cmd CreateQuizCommand) (*Quiz, error)
}
