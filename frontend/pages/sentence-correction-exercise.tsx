import React, { ChangeEvent } from 'react';

type Props = {
  sentence: string
  answer?: string
  setAnswer: (answer: string) => void
};

const SentenceCorrectionExercise: React.FC<Props> = ({ sentence, answer, setAnswer }) => {

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setAnswer(event.target.value);
  };

  return (
    <div className="border border-black">
      <div className='font-bold'>Sentence correction</div>
      <div>{sentence}</div>
      <input
        className="border border-black"
        type="text"
        value={answer ?? ""}
        onChange={handleChange}
      />
    </div>
  );
};

export default SentenceCorrectionExercise;