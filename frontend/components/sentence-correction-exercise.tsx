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
        {correctAnswer != null && answeredCorrectly && <span> ✅</span>}
        {correctAnswer != null && !answeredCorrectly && <span> ❌</span>}
      </div>
      <div>
        <input
          className={`border`}
          type="text"
          value={answer ?? ""}
          onChange={handleChange}
          disabled={disabled}
          required
        />
        {correctAnswer != null && !answeredCorrectly && (
          <div className="text-red-500">Correct answer: {correctAnswer}</div>
        )}
      </div>
      {feedback && !answeredCorrectly && <Feedback feedback={feedback} />}
    </div>
  );
};

export default SentenceCorrectionExercise;