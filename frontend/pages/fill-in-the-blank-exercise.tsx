import React, { ChangeEvent } from 'react';

type Props = {
  question: string
  answer?: string
  setAnswer: (answer: string) => void
};

const FillInTheBlankExercise: React.FC<Props> = ({ question, answer, setAnswer }) => {

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setAnswer(event.target.value);
  };

  const parts = question.split("{0}");
  const before = parts[0]; 
  const after = parts[1]; 

  return (
    <div className="border border-black">
      {before}
      <input
        className="border border-black"
        type="text"
        value={answer ?? ''}
        onChange={handleChange}
      />
      {after}
    </div>
  );
};

export default FillInTheBlankExercise;