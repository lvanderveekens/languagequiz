package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"languagequiz/drill"
	"languagequiz/drill/exercise"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DrillHandler struct {
	// drillStorage drill.Storage
}

func NewDrillHandler() *DrillHandler {
	return &DrillHandler{}
}

func (h *DrillHandler) CreateDrill(c *gin.Context) error {
	var req createDrillRequest
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

	fmt.Printf("Command: %+v\n", *cmd)

	c.JSON(http.StatusCreated, nil)
	return nil
}

type createDrillRequest struct {
	Name     string                 `json:"name"`
	Sections []createSectionRequest `json:"sections"`
}

func (r *createDrillRequest) validate() error {
	if r.Name == "" {
		return errors.New("required field is missing: name")
	}
	if r.Sections == nil {
		return errors.New("required field is missing: sections")
	}
	for _, createSectionRequest := range r.Sections {
		if err := createSectionRequest.validate(); err != nil {
			return fmt.Errorf("section validation error: %w", err)
		}
	}
	return nil
}

func (r *createDrillRequest) toCommand() (*drill.CreateDrillCommand, error) {
	createSectionCommands := make([]drill.CreateSectionCommand, 0)
	for _, createSectionRequest := range r.Sections {
		createSectionCommand, err := createSectionRequest.toCommand()
		if err != nil {
			return nil, err
		}
		createSectionCommands = append(createSectionCommands, *createSectionCommand)
	}

	createDrillCommand := drill.NewCreateDrillCommand(r.Name, createSectionCommands)
	return &createDrillCommand, nil
}

type createSectionRequest struct {
	Name      string            `json:"name"`
	Exercises []json.RawMessage `json:"exercises"`
}

func (r *createSectionRequest) validate() error {
	if r.Name == "" {
		return errors.New("required field is missing: name")
	}
	if r.Exercises == nil {
		return errors.New("required field is missing: exercises")
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
			// TODO: more exercise types
		default:
			return fmt.Errorf("unsupported exercise type: %v", createExerciseRequestJson["type"])
		}

	}
	return nil
}

func (r *createSectionRequest) toCommand() (*drill.CreateSectionCommand, error) {
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
			// TODO: more exercise types
		default:
			return nil, NewError(http.StatusBadRequest, fmt.Sprintf("unsupported exercise type: %v", createExerciseRequestJson["type"]))
		}

	}

	return &drill.CreateSectionCommand{
		Name:      r.Name,
		Exercises: createExerciseCommands,
	}, nil
}
