import React, { useState } from 'react';
import FeedbackForm from './feedback-form';

type Props = {
  // text: string;
};

const FeedbackButton: React.FC<Props> = ({}) => {

  const [showForm, setShowForm] = useState<boolean>(false)

  const handleClick = () => {
    setShowForm(true)
  }

  const handleSubmit = async (text: string) => {
    try {
      const res = await fetch(`http://localhost:8888/v1/feedback`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          "text": text
        }),
      });
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <div className="fixed bottom-0 right-4 z-50 flex flex-col">
      {showForm && (
        <FeedbackForm
          className="mb-4"
          onSubmit={handleSubmit}
          onClose={() => setShowForm(false)}
        />
      )}
      <button
        className="text-white font-bold bg-[#003259] border-x-2 border-t-2 border-[#003259] 
          hover:text-[#003259] hover:bg-white rounded-t-lg px-4 py-2 ml-auto"
        onClick={handleClick}
      >
        Feedback
      </button>
    </div>
  );
};

export default FeedbackButton;