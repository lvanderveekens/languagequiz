import React, { ChangeEvent } from 'react';
import ExerciseFeedback from './exercise-feedback';

type Props = {
  index: number
  sentence: string
  answer?: string
  setAnswer: (answer: string) => void
  correct?: boolean
  correctAnswer?: string
  feedback?: string
  disabled?: boolean
};

const SentenceCorrectionExercise: React.FC<Props> = ({
  index,
  sentence,
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

  return (
    <div className="">
      <div>
        {index + 1}. {sentence}
      </div>
      <div>
        <input
          className={`
            border
            ${correct != null && correct ? "text-green-500" : ""}
            ${correct != null && !correct ? "text-red-500" : ""}
            disabled:opacity-100
          `}
          type="text"
          value={answer ?? ""}
          onChange={handleChange}
          disabled={disabled}
          required
        />
        {correctAnswer != null && !correct && (
          <div className="text-red-500">Correct answer: {correctAnswer}</div>
        )}
      </div>
      {feedback && !correct && <ExerciseFeedback text={feedback} />}
    </div>
  );
};

export default SentenceCorrectionExercise;