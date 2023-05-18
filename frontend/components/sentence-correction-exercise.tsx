import React, { ChangeEvent } from 'react';

type Props = {
  index: number
  sentence: string
  answer?: string
  setAnswer: (answer: string) => void
};

const SentenceCorrectionExercise: React.FC<Props> = ({ index, sentence, answer, setAnswer }) => {

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setAnswer(event.target.value);
  };

  return (
    <div className="">
      <div>{index + 1}. {sentence}</div>
      <input
        className="border border-black"
        type="text"
        value={answer ?? ""}
        onChange={handleChange}
        required
      />
    </div>
  );
};

export default SentenceCorrectionExercise;