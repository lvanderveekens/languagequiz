import React from 'react';

type Props = {
  question: string
  options: string[]
};

const MultipleChoiceExercise: React.FC<Props> = ({ question, options }) => {
  return (
    <div className='border border-black'>
      <div>{question}</div>
      <div>
        {options.map((option) => (
          <div>{option}</div>
        ))}
        </div>
    </div>
  );
};

export default MultipleChoiceExercise;