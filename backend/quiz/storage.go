package quiz

type Storage interface {
	FindByID(id string) (*Quiz, error)
	FindQuizzes() ([]Quiz, error)
	CreateQuiz(cmd CreateQuizCommand) (*Quiz, error)
}
