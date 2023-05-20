import React from 'react';
import Feedback from './feedback';

type Props = {
  index: number
  question: string
  choices: string[]
  answer?: string
  setAnswer: (answer: string) => void
  correctAnswer?: string
  feedback?: string
  disabled?: boolean
};

const MultipleChoiceExercise: React.FC<Props> = ({
  index,
  question,
  choices,
  answer,
  setAnswer,
  correctAnswer,
  feedback,
  disabled,
}) => {
  const answeredCorrectly = answer === correctAnswer;

  return (
    <div className="">
      <div>
        {index + 1}. {question}
        {correctAnswer != null && answeredCorrectly && <span> ✅</span>}
        {correctAnswer != null && !answeredCorrectly && <span> ❌</span>}
      </div>
      <div>
        {choices.map((choice: string) => (
          <div>
            <label
              key={choice}
              className={``}
            >
              <input
                className="mr-2"
                type="radio"
                value={choice ?? ""}
                checked={answer === choice}
                onChange={(event) => setAnswer(event.target.value)}
                name={`exercise-${index + 1}`}
                required
                disabled={disabled}
              />
              {choice}
            </label>
          </div>
        ))}
        {correctAnswer != null && !answeredCorrectly && (
          <div className="text-red-500">Correct answer: {correctAnswer}</div>
        )}
        {feedback && !answeredCorrectly && <Feedback feedback={feedback} />}
      </div>
    </div>
  );
};

export default MultipleChoiceExercise;