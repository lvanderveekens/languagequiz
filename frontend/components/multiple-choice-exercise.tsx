import React, { useState } from 'react';
import { ChangeEvent } from 'react';

type Props = {
  index: number
  question: string
  choices: string[]
  answer?: string
  setAnswer: (answer: string) => void
};

const MultipleChoiceExercise: React.FC<Props> = ({
  index,
  question,
  choices,
  answer,
  setAnswer,
}) => {
  return (
    <div className="">
      <div>{index + 1}. {question}</div>
      <div>
        {choices.map((choice: string) => (
          <div key={choice}>
            <label>
              <input
              className='mr-2'
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