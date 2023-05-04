import React, { useState } from 'react';
import { ChangeEvent } from 'react';

type Props = {
  question: string
  options: string[]
  answer?: string
  setAnswer: (answer: string) => void
};

const MultipleChoiceExercise: React.FC<Props> = ({ question, options, answer, setAnswer }) => {
  return (
    <div className="border border-black">
      <div>{question}</div>
      <div>
        {options.map((option: string) => (
          <div key={option}>
            <label>
              <input
                type="radio"
                value={option ?? ''}
                checked={answer === option}
                onChange={(event) => setAnswer(event.target.value)}
              />
              {option}
            </label>
          </div>
        ))}
      </div>
    </div>
  );
};

export default MultipleChoiceExercise;