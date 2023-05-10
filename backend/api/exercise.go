package api

import (
	"errors"
	"fmt"

	"languagequiz/quiz/exercise"
)

// func (h *ExerciseHandler) CreateExercise(c *gin.Context) error {
// 	bodyBytes, err := io.ReadAll(c.Request.Body)
// 	if err != nil {
// 		return fmt.Errorf("failed to read request body: %w", err)
// 	}

// 	var req createExerciseRequestBase
// 	err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&req)
// 	if err != nil {
// 		return fmt.Errorf("failed to decode request body as map: %w", err)
// 	}

// 	var dto any
// 	switch req.Type {
// 	case exercise.TypeMultipleChoice:
// 		var req createMultipleChoiceExerciseRequest
// 		if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&req); err != nil {
// 			return NewError(http.StatusBadRequest, fmt.Sprintf("failed to decode request body as struct: %s", err.Error()))
// 		}
// 		if err := req.Validate(); err != nil {
// 			return NewError(http.StatusBadRequest, err.Error())
// 		}

// 		_, err := req.toCommand()
// 		if err != nil {
// 			return NewError(http.StatusBadRequest, err.Error())
// 		}

// 		// exercise, err := h.exerciseStorage.CreateMultipleChoiceExercise(*cmd)
// 		// if err != nil {
// 		// 	return fmt.Errorf("failed to create exercise: %w", err)
// 		// }

// 		// dto = newMultipleChoiceExerciseDto(*exercise)
// 	case exercise.TypeFillInTheBlank:
// 		var req createFillInTheBlankExerciseRequest
// 		if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&req); err != nil {
// 			return NewError(http.StatusBadRequest, fmt.Sprintf("failed to decode request body as struct: %s", err.Error()))
// 		}
// 		if err := req.Validate(); err != nil {
// 			return NewError(http.StatusBadRequest, err.Error())
// 		}

// 		cmd, err := req.toCommand()
// 		if err != nil {
// 			return NewError(http.StatusBadRequest, err.Error())
// 		}

// 		exercise, err := h.exerciseStorage.CreateFillInTheBlankExercise(*cmd)
// 		if err != nil {
// 			return fmt.Errorf("failed to create exercise: %w", err)
// 		}

// 		dto = newFillInTheBlankExerciseDto(*exercise)
// 	case exercise.TypeSentenceCorrection:
// 		var req createSentenceCorrectionExerciseRequest
// 		if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&req); err != nil {
// 			return NewError(http.StatusBadRequest, fmt.Sprintf("failed to decode request body as struct: %s", err.Error()))
// 		}
// 		if err := req.Validate(); err != nil {
// 			return NewError(http.StatusBadRequest, err.Error())
// 		}

// 		cmd, err := req.toCommand()
// 		if err != nil {
// 			return NewError(http.StatusBadRequest, err.Error())
// 		}

// 		exercise, err := h.exerciseStorage.CreateSentenceCorrectionExercise(*cmd)
// 		if err != nil {
// 			return fmt.Errorf("failed to create exercise: %w", err)
// 		}

// 		dto = newSentenceCorrectionExerciseDto(*exercise)
// 	default:
// 		return NewError(http.StatusBadRequest, fmt.Sprintf("Unsupported exercise type: %v", req.Type))
// 	}

// 	c.JSON(http.StatusCreated, dto)
// 	return nil
// }

// func (h *ExerciseHandler) GetExercises(c *gin.Context) error {
// 	exercises, err := h.exerciseStorage.Find()
// 	if err != nil {
// 		return fmt.Errorf("failed to find exercises: %w", err)
// 	}

// 	dtos := make([]any, 0)
// 	for _, exercise := range exercises {
// 		dto, err := mapExerciseToDto(exercise)
// 		if err != nil {
// 			return fmt.Errorf("failed to map exercise to dto: %w", err)
// 		}
// 		dtos = append(dtos, dto)
// 	}

// 	c.JSON(http.StatusOK, dtos)
// 	return nil
// }

// func (h *ExerciseHandler) SubmitAnswers(c *gin.Context) error {
// 	// TODO: validate request
// 	var req submitAnswersRequest
// 	err := json.NewDecoder(c.Request.Body).Decode(&req)
// 	if err != nil {
// 		return fmt.Errorf("failed to decode request body: %w", err)
// 	}

// 	results := make([]exerciseResult, 0)
// 	for _, submission := range req.Submissions {
// 		e, err := h.exerciseStorage.FindByID(submission.ExerciseID)
// 		if err != nil {
// 			return fmt.Errorf("failed to find exercise by id: %w", err)
// 		}

// 		correct := e.CheckAnswer(submission.Answer)
// 		result := newExerciseResult(submission.ExerciseID, e.GetType(), correct, e.GetAnswer())
// 		results = append(results, result)
// 	}

// 	c.JSON(http.StatusOK, newSubmitAnswersResponse(results))
// 	return nil
// }

func mapExerciseToDTO(e exercise.Exercise) (any, error) {
	switch e := e.(type) {
	case *exercise.MultipleChoiceExercise:
		return mapMultipleChoiceExerciseToDTO(*e), nil
	case *exercise.FillInTheBlankExercise:
		return mapFillInTheBlankExerciseToDTO(*e), nil
	case *exercise.SentenceCorrectionExercise:
		return mapSentenceCorrectionExerciseToDTO(*e), nil
	default:
		return nil, fmt.Errorf("unknown exercise type: %T", e)
	}
}

type exerciseDTOBase struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func newExerciseDTOBase(id string, exerciseType string) exerciseDTOBase {
	return exerciseDTOBase{
		ID:   id,
		Type: exerciseType,
	}
}

type multipleChoiceExerciseDTO struct {
	exerciseDTOBase
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
	Answer   string   `json:"answer"`
}

func newMultipleChoiceExerciseDTO(id, question string, choices []string, answer string) multipleChoiceExerciseDTO {
	return multipleChoiceExerciseDTO{
		exerciseDTOBase: newExerciseDTOBase(id, exercise.TypeMultipleChoice),
		Question:        question,
		Choices:         choices,
		Answer:          answer,
	}
}

func mapMultipleChoiceExerciseToDTO(e exercise.MultipleChoiceExercise) multipleChoiceExerciseDTO {
	return newMultipleChoiceExerciseDTO(e.ID, e.Question, e.Choices, e.Answer)
}

type fillInTheBlankExerciseDTO struct {
	exerciseDTOBase
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func newFillInTheBlankExerciseDTO(id, question, answer string) fillInTheBlankExerciseDTO {
	return fillInTheBlankExerciseDTO{
		exerciseDTOBase: newExerciseDTOBase(id, exercise.TypeFillInTheBlank),
		Question:        question,
		Answer:          answer,
	}
}

func mapFillInTheBlankExerciseToDTO(e exercise.FillInTheBlankExercise) fillInTheBlankExerciseDTO {
	return newFillInTheBlankExerciseDTO(e.ID, e.Question, e.Answer)
}

type sentenceCorrectionExerciseDTO struct {
	exerciseDTOBase
	Sentence          string `json:"sentence"`
	CorrectedSentence string `json:"correctedSentence"`
}

func newSentenceCorrectionExerciseDTO(id, sentence, correctedSentence string) sentenceCorrectionExerciseDTO {
	return sentenceCorrectionExerciseDTO{
		exerciseDTOBase:   newExerciseDTOBase(id, exercise.TypeSentenceCorrection),
		Sentence:          sentence,
		CorrectedSentence: correctedSentence,
	}
}

func mapSentenceCorrectionExerciseToDTO(e exercise.SentenceCorrectionExercise) sentenceCorrectionExerciseDTO {
	return newSentenceCorrectionExerciseDTO(e.ID, e.Sentence, e.CorrectedSentence)
}

type createExerciseRequestBase struct {
	Type string `json:"type"`
}

type createMultipleChoiceExerciseRequest struct {
	createExerciseRequestBase
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
	Answer   string   `json:"answer"`
}

func (r *createMultipleChoiceExerciseRequest) toCommand() (*exercise.CreateMultipleChoiceExerciseCommand, error) {
	return exercise.NewCreateMultipleChoiceExerciseCommand(
		r.Question,
		r.Choices,
		r.Answer,
	)
}

func (r *createMultipleChoiceExerciseRequest) Validate() error {
	if r.Question == "" {
		return errors.New("required field is missing: question")
	}
	if r.Choices == nil {
		return errors.New("required field is missing: choices")
	}
	if r.Answer == "" {
		return errors.New("required field is missing: answer")
	}
	return nil
}

type createFillInTheBlankExerciseRequest struct {
	createExerciseRequestBase
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (r *createFillInTheBlankExerciseRequest) toCommand() (*exercise.CreateFillInTheBlankExerciseCommand, error) {
	return exercise.NewCreateFillInTheBlankExerciseCommand(
		r.Question,
		r.Answer,
	)
}

func (r *createFillInTheBlankExerciseRequest) Validate() error {
	if r.Question == "" {
		return errors.New("required field is missing: question")
	}
	if r.Answer == "" {
		return errors.New("required field is missing: answer")
	}
	return nil
}

type createSentenceCorrectionExerciseRequest struct {
	createExerciseRequestBase
	Sentence          string `json:"sentence"`
	CorrectedSentence string `json:"correctedSentence"`
}

func (r *createSentenceCorrectionExerciseRequest) toCommand() (*exercise.CreateSentenceCorrectionExerciseCommand, error) {
	return exercise.NewCreateSentenceCorrectionExerciseCommand(
		r.Sentence,
		r.CorrectedSentence,
	)
}

func (r *createSentenceCorrectionExerciseRequest) Validate() error {
	if r.Sentence == "" {
		return errors.New("required field is missing: sentence")
	}
	if r.CorrectedSentence == "" {
		return errors.New("required field is missing: correctedSentence")
	}
	return nil
}

// type submitAnswersRequest struct {
// 	Submissions []exerciseSubmission `json:"submissions"`
// }

// type exerciseSubmission struct {
// 	ExerciseID string `json:"exerciseId"`
// 	Answer     any    `json:"answer"`
// }

// type submitAnswersResponse struct {
// 	Results []exerciseResult `json:"results"`
// }

// func newSubmitAnswersResponse(results []exerciseResult) submitAnswersResponse {
// 	return submitAnswersResponse{
// 		Results: results,
// 	}
// }

// type exerciseResult struct {
// 	ExerciseID   string `json:"exerciseId"`
// 	ExerciseType string `json:"exerciseType"`
// 	Correct      bool   `json:"correct"`
// 	Answer       any    `json:"answer"`
// }

// func newExerciseResult(exerciseID string, exerciseType string, correct bool, answer any) exerciseResult {
// 	return exerciseResult{
// 		ExerciseID:   exerciseID,
// 		ExerciseType: exerciseType,
// 		Correct:      correct,
// 		Answer:       answer,
// 	}
// }
