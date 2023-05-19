import React, { ChangeEvent } from 'react';
import Feedback from './feedback';

type Props = {
  index: number,
  question: string
  answer?: string
  setAnswer: (answer: string) => void
  correctAnswer?: string
  feedback?: string
  disabled?: boolean
};

const FillInTheBlankExercise: React.FC<Props> = ({
  index,
  question,
  answer,
  setAnswer,
  correctAnswer,
  feedback,
  disabled,
}) => {
  const answeredCorrectly = answer === correctAnswer;

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
          className={`border
            ${correctAnswer != null && answeredCorrectly ? "text-green-500 border-green-500" : ""}
            ${correctAnswer != null && !answeredCorrectly ? "text-red-500 border-red-500" : ""}
          `}
          type="text"
          value={answer ?? ""}
          onChange={handleChange}
          required
          disabled={disabled}
        />
        {after}
      </div>
      {feedback && !answeredCorrectly && <Feedback feedback={feedback} />}
    </div>
  );
};

export default FillInTheBlankExercise;