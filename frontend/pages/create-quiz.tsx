import { useState } from "react";
import { v4 as uuidv4 } from 'uuid';

export default function CreateQuizPage() {
  const [formValues, setFormValues] = useState<QuizFormValues>({})

  const handleNameChange = (event: any) => {
    setFormValues({ ...formValues, name: event.target.value });
  };

  const handleSubmit = async (event: any) => {
    event.preventDefault();

    console.log(`Form submitted with values: ${JSON.stringify(formValues)}`)
  };

  const handleSectionNameChange = (sectionIndex: any) => (sectionName: string) => {
    const updatedSections = [...(formValues.sections ?? [])]
    updatedSections[sectionIndex].name = sectionName

    const updatedFormValues = {...formValues, sections: updatedSections}
    setFormValues(updatedFormValues)
  }

  const handleExercisesChange = (sectionIndex: any) => (exercises: ExerciseFormValues[]) => {
    const updatedSections = [...(formValues.sections ?? [])]
    updatedSections[sectionIndex].exercises = exercises

    const updatedFormValues = {...formValues, sections: updatedSections}
    setFormValues(updatedFormValues)
  }

  const handleAddSectionClick = () => {
    const newSection = {_key: uuidv4()}
    setFormValues({ ...formValues, sections: [...(formValues.sections ?? []), newSection] });
  };

  return (
    <div className="container mx-auto">
      Create quiz page
      <form onSubmit={handleSubmit}>
        <div>
          <label>
            <span className="mr-3">Name</span>
            <input
              className="border border-black"
              type="text"
              value={formValues.name ?? ""}
              onChange={handleNameChange}
            />
          </label>
        </div>

        {formValues.sections &&
          formValues.sections.map((formValues: QuizSectionFormValues, i) => (
            <QuizSectionInput
              key={formValues._key}
              name={formValues.name}
              onNameChange={handleSectionNameChange(i)}
              exercises={formValues.exercises}
              onExercisesChange={handleExercisesChange(i)}
            />
          ))}

        <div>
          <button
            type="button"
            className="border border-black px-3"
            onClick={handleAddSectionClick}
          >
            Add section
          </button>
        </div>
        <div>
          <button className="border border-black px-3" type="submit">
            Submit
          </button>
        </div>
      </form>
    </div>
  );
}

type QuizSectionInputProps = {
  name?: string
  onNameChange: (name: string) => void
  exercises?: ExerciseFormValues[]
  onExercisesChange: (exercises: ExerciseFormValues[]) => void
};

const QuizSectionInput: React.FC<QuizSectionInputProps> = ({
  name,
  onNameChange,
  exercises,
  onExercisesChange,
}) => {

  const handleAddExerciseClick = () => {
    const newExercise = { _key: uuidv4() };
    onExercisesChange([...(exercises ?? []), newExercise]);
  };

  const handleExerciseChange =
    (index: number) => (value: ExerciseFormValues) => {
      const updatedExercises = [...(exercises ?? [])];
      updatedExercises[index] = value;

      onExercisesChange(updatedExercises);
    };

  return (
    <div className="border border-black">
      Section
      <div>
        <label>
          <span className="mr-3">Name</span>
          <input
            className="border border-black"
            type="text"
            value={name ?? ""}
            onChange={(e) => onNameChange(e.target.value)}
          />
        </label>
      </div>
      {exercises &&
        exercises.map((formValues: ExerciseFormValues, i) => (
          <ExerciseInput key={formValues._key} value={exercises[i]} onChange={handleExerciseChange(i)} />
        ))}
      <div>
        <button
          type="button"
          className="border border-black px-3"
          onClick={handleAddExerciseClick}
        >
          Add exercise
        </button>
      </div>
    </div>
  );
}

type ExerciseInputProps = {
  value: ExerciseFormValues
  onChange: (value: ExerciseFormValues) => void
};

const ExerciseInput: React.FC<ExerciseInputProps> = ({
  value,
  onChange,
}) => {

  // const [exerciseType, setExerciseType] = useState<ExerciseType | undefined>()

  const handleSelectChange = (event: any) => {
    onChange({...value, type: event.target.value})
  };

  if (!value.type) {
    return (
      <div className="border border-black">
        Exercise
        <div>
          <label htmlFor="choice">Choose an exercise type:</label>
          <select
            id="choice"
            name="choice"
            value={value.type}
            onChange={handleSelectChange}
          >
            <option value="">Select an option</option>
            {Object.values(ExerciseType)
              .filter((key) => isNaN(Number(key)))
              .map((exerciseType) => (
                <option key={exerciseType} value={exerciseType}>{exerciseType}</option>
              ))}
          </select>
        </div>
      </div>
    );
  }

  switch (value.type) {
    case ExerciseType.MultipleChoice:
      return (
        <MultipleChoiceExerciseInput
          question={value.question}
          choices={value.choices}
          answer={value.answer}
          onQuestionChange={(question: string) => onChange({...value, question: question})}
          onChoicesChange={(choices: string[]) => onChange({...value, choices: choices})}
          onAnswerChange={(answer: string) => onChange({...value, answer: answer})}
        />
      );
    case ExerciseType.FillInTheBlank:
      return <FillInTheBlankExerciseInput />
    case ExerciseType.SentenceCorrection:
      return <SentenceCorrectionExerciseInput />
    default:
      return <p>Unknown exercise type</p>
  }
}

type MultipleChoiceExerciseInputProps = {
  question?: string
  choices?: string[]
  answer?: string
  onQuestionChange: (question: string) => void
  onChoicesChange: (choices: string[]) => void
  onAnswerChange: (answer: string) => void
};

const MultipleChoiceExerciseInput: React.FC<MultipleChoiceExerciseInputProps> = ({
  question,
  choices,
  answer,
  onQuestionChange,
  onChoicesChange,
  onAnswerChange,
}) => {

  // TODO: 1 question
  // TODO: 4 choices
  // TODO: 1 answer

  const handleChoiceChange = (index: number) => (event: any) => {
    const { value } = event.target;
    const updatedChoices = [...(choices ?? new Array(4))]
    updatedChoices[index] = value
    onChoicesChange(updatedChoices);
  };

  return (
    <div className="border border-black">
      Multiple Choice Exercise!!
      <div>
        <label>
          <span className="mr-3">Question</span>
          <input
            className="border border-black"
            type="text"
            value={question ?? ""}
            onChange={(e) => onQuestionChange(e.target.value)}
          />
        </label>
      </div>
      <div>
        <label>
          <span className="mr-3">Choice 1</span>
          <input
            className="border border-black"
            type="text"
            value={choices?.[0] ?? ""}
            onChange={handleChoiceChange(0)}
          />
        </label>
      </div>
      <div>
        <label>
          <span className="mr-3">Answer</span>
          <input
            className="border border-black"
            type="text"
            value={answer ?? ""}
            onChange={(e) => onAnswerChange(e.target.value)}
          />
        </label>
      </div>
    </div>
  );
}

type FillInTheBlankExerciseInputProps = {
};

const FillInTheBlankExerciseInput: React.FC<FillInTheBlankExerciseInputProps> = ({
}) => {

  return (
    <div className="border border-black">
      Fill In The Blank Exercise
    </div>
  );
}

type SentenceCorrectionExerciseInputProps = {
};

const SentenceCorrectionExerciseInput: React.FC<SentenceCorrectionExerciseInputProps> = ({
}) => {

  return (
    <div className="border border-black">
      Sentence Correction Exercise
    </div>
  );
}

interface QuizFormValues {
    name?: string
    sections?: QuizSectionFormValues[]
}

interface QuizSectionFormValues {
    _key: string
    name?: string
    exercises?: ExerciseFormValues[]
}

interface ExerciseFormValues {
  _key: string;
  type?: string;
  question?: string;
  choices?: string[];
  sentence?: string;
  answer?: string
}

enum ExerciseType {
  MultipleChoice = "multipleChoice",
  FillInTheBlank = "fillInTheBlank",
  SentenceCorrection = "sentenceCorrection",
}