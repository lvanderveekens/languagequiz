import React, { useState } from 'react';
import FeedbackModal from './feedback-modal';
import { useRouter } from 'next/router';

type Props = {
};

const FeedbackButton: React.FC<Props> = ({}) => {
  const router = useRouter();

  const [showModal, setShowModal] = useState<boolean>(false)

  const handleClick = (event: any) => {
    event.stopPropagation();
    setShowModal(!showModal);
  };

  const handleSubmit = async (text: string) => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/v1/feedback`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          text: text,
          pagePath: router.asPath,
        }),
      });
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <div className="">
      {showModal && <FeedbackModal onSubmit={handleSubmit} onClose={() => setShowModal(false)} />}
      <div className="fixed top-3/4 right-0 transform origin-bottom-left translate-x-full -rotate-90 z-10">
        <button className="text-white font-bold bg-black border-black border-x-2 border-t-2 
          hover:text-black hover:bg-white rounded-t-lg px-4 py-1 " onClick={handleClick}>
          Feedback
        </button>
      </div>
    </div>
  );
};

export default FeedbackButton;