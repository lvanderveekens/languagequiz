package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"languagequiz/quiz"
	"languagequiz/quiz/exercise"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

type QuizHandler struct {
	quizStorage quiz.Storage
}

func NewQuizHandler(quizStorage quiz.Storage) *QuizHandler {
	return &QuizHandler{quizStorage: quizStorage}
}

func (h *QuizHandler) GetQuizByID(c *gin.Context) error {
	id := c.Param("id")

	quiz, err := h.quizStorage.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find quiz: %w", err)
	}

	dto, err := mapToQuizDTO(*quiz)
	if err != nil {
		return fmt.Errorf("failed to map quiz to dto: %w", err)
	}

	c.JSON(http.StatusOK, *dto)
	return nil
}

func (h *QuizHandler) GetQuizzes(c *gin.Context) error {
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
	Name        string                     `json:"name"`
	LanguageTag string                     `json:"languageTag"`
	Sections    []createQuizSectionRequest `json:"sections"`
}

func (r *createQuizRequest) validate() error {
	if r.Name == "" {
		return errors.New("field 'name' is missing")
	}
	if r.LanguageTag == "" {
		return errors.New("field 'languageTag' is missing")
	}
	if r.Sections == nil {
		return errors.New("field 'sections' is missing")
	}
	if len(r.Sections) == 0 {
		return errors.New("field 'sections' is empty")
	}
	for _, createSectionRequest := range r.Sections {
		if err := createSectionRequest.validate(); err != nil {
			return fmt.Errorf("section validation error: %w", err)
		}
	}
	return nil
}

func (r *createQuizRequest) toCommand() (*quiz.CreateQuizCommand, error) {
	languageTag, err := language.Parse(r.LanguageTag)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, err.Error())
	}

	createSectionCommands := make([]quiz.CreateSectionCommand, 0)
	for _, createSectionRequest := range r.Sections {
		createSectionCommand, err := createSectionRequest.toCommand()
		if err != nil {
			return nil, err
		}
		createSectionCommands = append(createSectionCommands, *createSectionCommand)
	}

	createQuizCommand := quiz.NewCreateQuizCommand(r.Name, languageTag, createSectionCommands)
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
	if len(r.Exercises) == 0 {
		return errors.New("field 'exercises' is empty")
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

func (r *createQuizSectionRequest) toCommand() (*quiz.CreateSectionCommand, error) {
	createExerciseCommands := make([]exercise.CreateExerciseCommand, 0)
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

			createExerciseCommands = append(createExerciseCommands, createExerciseCommand)
		case exercise.TypeFillInTheBlank:
			var createExerciseRequest createFillInTheBlankExerciseRequest
			if err := json.Unmarshal(createExerciseRequestRaw, &createExerciseRequest); err != nil {
				return nil, fmt.Errorf("failed to decode exercise: %w", err)
			}

			createExerciseCommand, err := createExerciseRequest.toCommand()
			if err != nil {
				return nil, NewError(http.StatusBadRequest, err.Error())
			}

			createExerciseCommands = append(createExerciseCommands, createExerciseCommand)
		case exercise.TypeSentenceCorrection:
			var createExerciseRequest createSentenceCorrectionExerciseRequest
			if err := json.Unmarshal(createExerciseRequestRaw, &createExerciseRequest); err != nil {
				return nil, fmt.Errorf("failed to decode exercise: %w", err)
			}

			createExerciseCommand, err := createExerciseRequest.toCommand()
			if err != nil {
				return nil, NewError(http.StatusBadRequest, err.Error())
			}

			createExerciseCommands = append(createExerciseCommands, createExerciseCommand)
		default:
			return nil, NewError(http.StatusBadRequest, fmt.Sprintf("unsupported exercise type: %v", createExerciseRequestJson["type"]))
		}
	}

	createSectionCommand, err := quiz.NewCreateSectionCommand(r.Name, createExerciseCommands)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, err.Error())
	}
	return createSectionCommand, nil
}

func mapToQuizDTO(q quiz.Quiz) (*QuizDTO, error) {
	quizSectionDTOs, err := mapToQuizSectionDTOs(q.Sections)
	if err != nil {
		return nil, fmt.Errorf("failed to map quiz sections to dtos: %w", err)
	}
	quizDTO := newQuizDTO(q.ID, q.CreatedAt, q.Name, q.LanguageTag.String(), quizSectionDTOs)
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
	ID          string           `json:"id"`
	CreatedAt   time.Time        `json:"createdAt"`
	Name        string           `json:"name"`
	LanguageTag string           `json:"languageTag"`
	Sections    []QuizSectionDTO `json:"sections"`
}

func newQuizDTO(id string, createdAt time.Time, name, languageTag string, sections []QuizSectionDTO) QuizDTO {
	return QuizDTO{
		ID:          id,
		CreatedAt:   createdAt,
		Name:        name,
		LanguageTag: languageTag,
		Sections:    sections,
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

type submitAnswersRequest struct {
	UserAnswers []any `json:"userAnswers"`
}

func (r *submitAnswersRequest) validate() error {
	if r.UserAnswers == nil {
		return errors.New("field 'userAnswers' is missing")
	}
	return nil
}

type submitAnswersResponse struct {
	Results []submitAnswerResult `json:"results"`
}

func newSubmitAnswersResponse(results []submitAnswerResult) submitAnswersResponse {
	return submitAnswersResponse{
		Results: results,
	}
}

type submitAnswerResult struct {
	Correct  bool    `json:"correct"`
	Answer   any     `json:"answer"`
	Feedback *string `json:"feedback,omitempty"`
}

func newSubmitAnswerResult(correct bool, answer any, feedback *string) submitAnswerResult {
	return submitAnswerResult{
		Correct:  correct,
		Answer:   answer,
		Feedback: feedback,
	}
}

func (h *QuizHandler) SubmitAnswers(c *gin.Context) error {
	id := c.Param("id")

	var req submitAnswersRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return fmt.Errorf("failed to decode request body: %w", err)
	}

	if err := req.validate(); err != nil {
		return NewError(http.StatusBadRequest, err.Error())
	}

	quiz, err := h.quizStorage.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find quiz by id: %s, cause: %w", id, err)
	}

	exercises := quiz.GetExercises()
	results := make([]submitAnswerResult, 0)
	for i, userAnswer := range req.UserAnswers {
		exercise := exercises[i]
		correct := exercise.CheckAnswer(userAnswer)
		results = append(results, newSubmitAnswerResult(correct, exercise.Answer(), exercise.Feedback()))
	}

	c.JSON(http.StatusOK, newSubmitAnswersResponse(results))
	return nil
}
