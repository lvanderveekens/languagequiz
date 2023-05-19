import React from 'react';

type Props = {
  feedback: string;
};

const Feedback: React.FC<Props> = ({ feedback }) => {
  return (
    <div className="pt-4">
      <span className="items-center px-4 py-2 border-2 border-[#ffea94] rounded-lg bg-[#fff6d2] inline-flex">
        <span className='text-4xl mr-2'>👩🏻‍🏫</span>{feedback}
      </span>
    </div>
  );
};

export default Feedback;