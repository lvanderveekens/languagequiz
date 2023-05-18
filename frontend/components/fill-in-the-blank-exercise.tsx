import React, { ChangeEvent } from 'react';

type Props = {
  index: number,
  question: string
  answer?: string
  setAnswer: (answer: string) => void
};

const FillInTheBlankExercise: React.FC<Props> = ({ index, question, answer, setAnswer }) => {

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setAnswer(event.target.value);
  };

  const parts = question.split("______");
  const before = parts[0]; 
  const after = parts[1]; 

  return (
    <div className="">
      <div>
        {index + 1}. {before}
        <input
          className="border border-black"
          type="text"
          value={answer ?? ""}
          onChange={handleChange}
          required
        />
        {after}
      </div>
    </div>
  );
};

export default FillInTheBlankExercise;