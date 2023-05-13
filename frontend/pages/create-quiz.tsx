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

type Props = {
  name?: string
  onNameChange: (name: string) => void
  exercises?: ExerciseFormValues[]
  onExercisesChange: (exercises: ExerciseFormValues[]) => void
};

const QuizSectionInput: React.FC<Props> = ({
  name,
  onNameChange: setName,
  exercises,
  onExercisesChange: setExercises
}) => {

  const handleAddExerciseClick = () => {
    const newExercise = { _key: uuidv4() };
    setExercises([...(exercises ?? []), newExercise]);
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
            onChange={(e) => setName(e.target.value)}
          />
        </label>
      </div>
      {exercises &&
        exercises.map((formValues: ExerciseFormValues, i) => (
          <p key={formValues._key}>EXERCISE YO</p>
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
}
