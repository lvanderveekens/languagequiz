import React from 'react';

type Props = {
  question: string
};

const FillInTheBlankExercise: React.FC<Props> = ({ question }) => {
  return (
    <div className='border border-black'>
      <div>{question}</div>
    </div>
  );
};

export default FillInTheBlankExercise;