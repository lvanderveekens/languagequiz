import React from 'react';
import { FaComment, FaRegComment } from 'react-icons/fa';

type Props = {
  index: number
  question: string
  choices: string[]
  answer?: string
  setAnswer: (answer: string) => void
  correctAnswer?: string
  feedback?: string
};

const MultipleChoiceExercise: React.FC<Props> = ({
  index,
  question,
  choices,
  answer,
  setAnswer,
  correctAnswer,
  feedback,
}) => {
  return (
    <div className="">
      <div>
        {index + 1}. {question}
      </div>
      <div>
        {choices.map((choice: string) => (
          <div>
            <label
              key={choice}
              className={`
                ${correctAnswer != null && correctAnswer === choice ? "text-green-500" : ""}
                ${correctAnswer != null && choice === answer && answer != correctAnswer ? "text-red-500" : ""}
              `}
            >
              <input
                className="mr-2"
                type="radio"
                value={choice ?? ""}
                checked={answer === choice}
                onChange={(event) => setAnswer(event.target.value)}
                name={`exercise-${index + 1}`}
                required
              />
              {choice}
            </label>
          </div>
        ))}
        {feedback && (
          <div className='pt-4'>
            <span className="p-4 border border-black inline-flex">
              <FaRegComment className="text-2xl inline mr-2" /> {feedback}
            </span>
          </div>
        )}
      </div>
    </div>
  );
};

export default MultipleChoiceExercise;