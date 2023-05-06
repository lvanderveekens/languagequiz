import React, { useState } from 'react';
import { ChangeEvent } from 'react';

type Props = {
  question: string
  choices: string[]
  answer?: string
  setAnswer: (answer: string) => void
};

const MultipleChoiceExercise: React.FC<Props> = ({
  question,
  choices,
  answer,
  setAnswer,
}) => {
  return (
    <div className="border border-black">
      <div className='font-bold'>Multiple choice</div>
      <div>{question}</div>
      <div>
        {choices.map((choice: string) => (
          <div key={choice}>
            <label>
              <input
                type="radio"
                value={choice ?? ""}
                checked={answer === choice}
                onChange={(event) => setAnswer(event.target.value)}
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