import { ExerciseFormValues, ExerciseType, labelByExerciseType } from "@/components/models";
import { useState } from "react";
import { GrFormClose } from "react-icons/gr";
import "/node_modules/flag-icons/css/flag-icons.min.css";
import { v4 as uuidv4 } from 'uuid';
import ExerciseInput from "./exercise-input";

type QuizSectionInputProps = {
  className?: string
  name?: string
  onNameChange: (name: string) => void
  exercises?: ExerciseFormValues[]
  onExercisesChange: (exercises: ExerciseFormValues[]) => void
  onRemove?: () => void
};

const QuizSectionInput: React.FC<QuizSectionInputProps> = ({
  className,
  name,
  onNameChange,
  exercises,
  onExercisesChange,
  onRemove,
}) => {
  const [exerciseType, setExerciseType] = useState<ExerciseType | null>(null);

  const handleRemoveClick = (event: any) => {
    event.preventDefault();
    onRemove!();
  };

  const addExercise = () => {
    const newExercise = buildExerciseFormValues(exerciseType!);
    onExercisesChange([...(exercises ?? []), newExercise]);
  };

  const removeExercise = (index: number) => () => {
    const updatedExercises = [...(exercises ?? [])];
    updatedExercises.splice(index, 1);
    onExercisesChange(updatedExercises);
  };

  const handleExerciseChange = (index: number) => (value: ExerciseFormValues) => {
    const updatedExercises = [...(exercises ?? [])];
    updatedExercises[index] = value;

    onExercisesChange(updatedExercises);
  };

  const handleExerciseTypeChange = (event: any) => {
    const exerciseType = event.target.value;
    setExerciseType(exerciseType);
    const newExercise = buildExerciseFormValues(exerciseType);
    onExercisesChange([newExercise]);
  };

  function buildExerciseFormValues(type: ExerciseType): ExerciseFormValues {
    return {
      _key: uuidv4(),
      type: type,
    };
  }

  return (
    <div className={`${className} w-full`}>
      {onRemove == null && <div className="mb-4 text-xl font-bold">Section</div>}
      {onRemove != null && (
        <div className="flex justify-between items-center mb-4">
          <div className="text-xl font-bold">Section</div>
          <button className="text-3xl" onClick={handleRemoveClick}>
            <GrFormClose />
          </button>
        </div>
      )}
      <div className="mb-4">
        <label>
          <span className="mr-3">Name</span>
          <input
            className="w-full"
            placeholder="Enter a name"
            type="text"
            value={name ?? ""}
            onChange={(e) => onNameChange(e.target.value)}
            required
          />
        </label>
      </div>

      <div className="mb-4">
        <label className="" htmlFor="choice">
          <div>Exercise type</div>
          <select value={exerciseType ?? ""} onChange={handleExerciseTypeChange} required>
            <option selected disabled value="">
              Select an exercise type
            </option>
            {Object.values(ExerciseType)
              .filter((key) => isNaN(Number(key)))
              .map((exerciseType) => (
                <option key={exerciseType} value={exerciseType}>
                  {labelByExerciseType[exerciseType]}
                </option>
              ))}
          </select>
        </label>
      </div>

      {exercises &&
        exercises.map((formValues: ExerciseFormValues, i) => (
          <ExerciseInput
            key={formValues._key}
            value={exercises[i]}
            onChange={handleExerciseChange(i)}
            onRemove={i != 0 ? removeExercise(i) : undefined}
          />
        ))}
      {exerciseType && (
        <div className="mb-4">
          <button
            type="button"
            className="border p-4 w-full px-3 flex items-center justify-center hover:border-black"
            onClick={addExercise}
          >
            <span className="text-2xl mr-1">➕</span>
            Add exercise
          </button>
        </div>
      )}
    </div>
  );
}
export default QuizSectionInput;