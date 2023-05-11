package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"languagequiz/quiz"
	"languagequiz/quiz/exercise"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuizHandler struct {
	quizStorage quiz.Storage
}

func NewQuizHandler(quizStorage quiz.Storage) *QuizHandler {
	return &QuizHandler{quizStorage: quizStorage}
}

func (h *QuizHandler) FindQuizzes(c *gin.Context) error {
	quizzes, err := h.quizStorage.FindQuizzes()
	if err != nil {
		return fmt.Errorf("failed to find quizzes: %w", err)
	}

	dtos, err := mapToQuizDTOs(quizzes)
	if err != nil {
		return fmt.Errorf("failed to map quizzes to dtos: %w", err)
	}

	c.JSON(http.StatusOK, dtos)
	return nil
}

func (h *QuizHandler) CreateQuiz(c *gin.Context) error {
	var req createQuizRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return fmt.Errorf("failed to decode request body: %w", err)
	}

	if err := req.validate(); err != nil {
		return NewError(http.StatusBadRequest, err.Error())
	}

	cmd, err := req.toCommand()
	if err != nil {
		return err
	}

	quiz, err := h.quizStorage.CreateQuiz(*cmd)
	if err != nil {
		return fmt.Errorf("failed to create quiz: %w", err)
	}

	dto, err := mapToQuizDTO(*quiz)
	if err != nil {
		return fmt.Errorf("failed to map quiz to dto: %w", err)
	}

	c.JSON(http.StatusCreated, *dto)
	return nil
}

type createQuizRequest struct {
	Name     string                     `json:"name"`
	Sections []createQuizSectionRequest `json:"sections"`
}

func (r *createQuizRequest) validate() error {
	if r.Name == "" {
		return errors.New("field 'name' is missing")
	}
	if r.Sections == nil {
		return errors.New("field 'sections' is missing")
	}
	for _, createSectionRequest := range r.Sections {
		if err := createSectionRequest.validate(); err != nil {
			return fmt.Errorf("section validation error: %w", err)
		}
	}
	return nil
}

func (r *createQuizRequest) toCommand() (*quiz.CreateQuizCommand, error) {
	createSectionCommands := make([]quiz.CreateQuizSectionCommand, 0)
	for _, createSectionRequest := range r.Sections {
		createSectionCommand, err := createSectionRequest.toCommand()
		if err != nil {
			return nil, err
		}
		createSectionCommands = append(createSectionCommands, *createSectionCommand)
	}

	createQuizCommand := quiz.NewCreateQuizCommand(r.Name, createSectionCommands)
	return &createQuizCommand, nil
}

type createQuizSectionRequest struct {
	Name      string            `json:"name"`
	Exercises []json.RawMessage `json:"exercises"`
}

func (r *createQuizSectionRequest) validate() error {
	if r.Name == "" {
		return errors.New("field 'name' is missing")
	}
	if r.Exercises == nil {
		return errors.New("field 'exercises' is missing")
	}
	for _, createExerciseRequestRaw := range r.Exercises {
		var createExerciseRequestJson map[string]any
		if err := json.Unmarshal(createExerciseRequestRaw, &createExerciseRequestJson); err != nil {
			return fmt.Errorf("failed to decode exercise: %w", err)
		}

		switch createExerciseRequestJson["type"] {
		case exercise.TypeMultipleChoice:
			var createExerciseRequest createMultipleChoiceExerciseRequest
			if err := json.Unmarshal(createExerciseRequestRaw, &createExerciseRequest); err != nil {
				return fmt.Errorf("failed to decode exercise: %w", err)
			}

			if err := createExerciseRequest.Validate(); err != nil {
				return fmt.Errorf("exercise validation error: %w", err)
			}
		case exercise.TypeFillInTheBlank:
			var createExerciseRequest createFillInTheBlankExerciseRequest
			if err := json.Unmarshal(createExerciseRequestRaw, &createExerciseRequest); err != nil {
				return fmt.Errorf("failed to decode exercise: %w", err)
			}

			if err := createExerciseRequest.Validate(); err != nil {
				return fmt.Errorf("exercise validation error: %w", err)
			}
		case exercise.TypeSentenceCorrection:
			var createExerciseRequest createSentenceCorrectionExerciseRequest
			if err := json.Unmarshal(createExerciseRequestRaw, &createExerciseRequest); err != nil {
				return fmt.Errorf("failed to decode exercise: %w", err)
			}

			if err := createExerciseRequest.Validate(); err != nil {
				return fmt.Errorf("exercise validation error: %w", err)
			}
		default:
			return fmt.Errorf("unsupported exercise type: %v", createExerciseRequestJson["type"])
		}

	}
	return nil
}

func (r *createQuizSectionRequest) toCommand() (*quiz.CreateQuizSectionCommand, error) {
	createExerciseCommands := make([]any, 0)
	for _, createExerciseRequestRaw := range r.Exercises {
		var createExerciseRequestJson map[string]any
		if err := json.Unmarshal(createExerciseRequestRaw, &createExerciseRequestJson); err != nil {
			return nil, fmt.Errorf("failed to decode exercise: %w", err)
		}

		switch createExerciseRequestJson["type"] {
		case exercise.TypeMultipleChoice:
			var createExerciseRequest createMultipleChoiceExerciseRequest
			if err := json.Unmarshal(createExerciseRequestRaw, &createExerciseRequest); err != nil {
				return nil, fmt.Errorf("failed to decode exercise: %w", err)
			}

			createExerciseCommand, err := createExerciseRequest.toCommand()
			if err != nil {
				return nil, NewError(http.StatusBadRequest, err.Error())
			}

			createExerciseCommands = append(createExerciseCommands, *createExerciseCommand)
		case exercise.TypeFillInTheBlank:
			var createExerciseRequest createFillInTheBlankExerciseRequest
			if err := json.Unmarshal(createExerciseRequestRaw, &createExerciseRequest); err != nil {
				return nil, fmt.Errorf("failed to decode exercise: %w", err)
			}

			createExerciseCommand, err := createExerciseRequest.toCommand()
			if err != nil {
				return nil, NewError(http.StatusBadRequest, err.Error())
			}

			createExerciseCommands = append(createExerciseCommands, *createExerciseCommand)
		case exercise.TypeSentenceCorrection:
			var createExerciseRequest createSentenceCorrectionExerciseRequest
			if err := json.Unmarshal(createExerciseRequestRaw, &createExerciseRequest); err != nil {
				return nil, fmt.Errorf("failed to decode exercise: %w", err)
			}

			createExerciseCommand, err := createExerciseRequest.toCommand()
			if err != nil {
				return nil, NewError(http.StatusBadRequest, err.Error())
			}

			createExerciseCommands = append(createExerciseCommands, *createExerciseCommand)
		default:
			return nil, NewError(http.StatusBadRequest, fmt.Sprintf("unsupported exercise type: %v", createExerciseRequestJson["type"]))
		}
	}

	return &quiz.CreateQuizSectionCommand{
		Name:      r.Name,
		Exercises: createExerciseCommands,
	}, nil
}

func mapToQuizDTO(q quiz.Quiz) (*QuizDTO, error) {
	quizSectionDTOs, err := mapToQuizSectionDTOs(q.Sections)
	if err != nil {
		return nil, fmt.Errorf("failed to map quiz sections to dtos: %w", err)
	}
	quizDTO := newQuizDTO(q.Name, quizSectionDTOs)
	return &quizDTO, nil
}

func mapToQuizDTOs(quizzes []quiz.Quiz) ([]QuizDTO, error) {
	dtos := make([]QuizDTO, 0)
	for _, quiz := range quizzes {
		dto, err := mapToQuizDTO(quiz)
		if err != nil {
			return nil, fmt.Errorf("failed to map quiz to dto: %w", err)
		}
		dtos = append(dtos, *dto)
	}
	return dtos, nil
}

func mapToQuizSectionDTOs(quizSections []quiz.Section) ([]QuizSectionDTO, error) {
	dtos := make([]QuizSectionDTO, 0)
	for _, quizSection := range quizSections {
		exerciseDTOs, err := mapToExerciseDTOs(quizSection.Exercises)
		if err != nil {
			return nil, fmt.Errorf("failed to map exercises to dtos: %w", err)
		}
		dtos = append(dtos, newQuizSectionDTO(quizSection.Name, exerciseDTOs))
	}
	return dtos, nil
}

func mapToExerciseDTOs(exercises []exercise.Exercise) ([]any, error) {
	dtos := make([]any, 0)
	for _, e := range exercises {
		dto, err := mapExerciseToDTO(e)
		if err != nil {
			return nil, fmt.Errorf("failed to map exercise to dto: %w", err)
		}
		dtos = append(dtos, dto)
	}
	return dtos, nil
}

type QuizDTO struct {
	Name     string           `json:"name"`
	Sections []QuizSectionDTO `json:"sections"`
}

func newQuizDTO(name string, sections []QuizSectionDTO) QuizDTO {
	return QuizDTO{
		Name:     name,
		Sections: sections,
	}
}

type QuizSectionDTO struct {
	Name      string `json:"name"`
	Exercises []any  `json:"exercises"`
}

func newQuizSectionDTO(name string, exercises []any) QuizSectionDTO {
	return QuizSectionDTO{
		Name:      name,
		Exercises: exercises,
	}
}
