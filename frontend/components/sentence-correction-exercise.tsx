import React, { ChangeEvent } from 'react';
import Feedback from './feedback';

type Props = {
  index: number
  sentence: string
  answer?: string
  setAnswer: (answer: string) => void
  correctAnswer?: string
  feedback?: string
  disabled?: boolean
};

const SentenceCorrectionExercise: React.FC<Props> = ({
  index,
  sentence,
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

  return (
    <div className="">
      <div>
        {index + 1}. {sentence}
      </div>
      <div>
        <input
          className={`border
            ${correctAnswer != null && answeredCorrectly ? "text-green-500 border-green-500" : ""}
            ${correctAnswer != null && !answeredCorrectly ? "text-red-500 border-red-500" : ""}
          `}
          type="text"
          value={answer ?? ""}
          onChange={handleChange}
          disabled={disabled}
        />
      </div>
      {feedback && !answeredCorrectly && <Feedback feedback={feedback} />}
    </div>
  );
};

export default SentenceCorrectionExercise;