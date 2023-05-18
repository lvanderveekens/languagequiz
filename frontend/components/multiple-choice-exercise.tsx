import React, { useState } from 'react';
import { ChangeEvent } from 'react';

type Props = {
  index: number
  question: string
  choices: string[]
  answer?: string
  setAnswer: (answer: string) => void
  correctAnswer?: string
};

const MultipleChoiceExercise: React.FC<Props> = ({
  index,
  question,
  choices,
  answer,
  setAnswer,
  correctAnswer,
}) => {
  return (
    <div className="">
      <div>
        {index + 1}. {question}
      </div>
      <div>
        {choices.map((choice: string) => (
          <div>
            <label
              key={choice}
              className={`${choice === correctAnswer ? "text-green-500" : ""}`}
            >
              <input
                className="mr-2"
                type="radio"
                value={choice ?? ""}
                checked={answer === choice}
                onChange={(event) => setAnswer(event.target.value)}
                name={`exercise-${index + 1}`}
                required
              />
              {choice}
            </label>
          </div>
        ))}
      </div>
    </div>
  );
};

export default MultipleChoiceExercise;