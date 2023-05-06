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

  const parts = question.split("______");
  const before = parts[0]; 
  const after = parts[1]; 

  return (
    <div className="border border-black">
      <div className='font-bold'>Fill in the blank</div>
      <div>
        {before}
        <input
          className="border border-black"
          type="text"
          value={answer ?? ""}
          onChange={handleChange}
        />
        {after}
      </div>
    </div>
  );
};

export default FillInTheBlankExercise;