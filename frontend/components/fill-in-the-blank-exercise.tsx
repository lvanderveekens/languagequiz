import React, { ChangeEvent } from 'react';
import ExerciseFeedback from './exercise-feedback';

type Props = {
  index: number,
  question: string
  answer?: string
  setAnswer: (answer: string) => void
  correct?: boolean
  correctAnswer?: string
  feedback?: string
  disabled?: boolean
};

const FillInTheBlankExercise: React.FC<Props> = ({
  index,
  question,
  answer,
  setAnswer,
  correct,
  correctAnswer,
  feedback,
  disabled,
}) => {
  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setAnswer(event.target.value);
  };

  const parts = question.split("______");
  const before = parts[0];
  const after = parts[1];

  return (
    <div className="">
      <div>
        {index + 1}. {before}
        <input
          className={`
            border
            ${correct != null && correct ? "text-green-500" : ""}
            ${correct != null && !correct ? "text-red-500" : ""}
            opacity-100
          `}
          type="text"
          value={answer ?? ""}
          onChange={handleChange}
          required
          disabled={disabled}
        />
        {after}
      </div>
      {correctAnswer && !correct && <div className="text-red-500">Correct answer: {correctAnswer}</div>}
      {feedback && !correct && <ExerciseFeedback text={feedback} />}
    </div>
  );
};

export default FillInTheBlankExercise;