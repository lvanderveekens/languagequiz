import React, { useState } from 'react';
import { GrFormClose } from 'react-icons/gr';
import Button from './button';

type Props = {
  className: string
  onSubmit: (text: any) => void
  onClose: () => void
};

const FeedbackForm: React.FC<Props> = ({ className, onSubmit, onClose }) => {

  const [text, setText] = useState<string>()
  const [submitting, setSubmitting] = useState<boolean>(false);
  const [submitted, setSubmitted] = useState<boolean>(false);

  const handleSubmit = (event: any) => {
    event.preventDefault();
    setSubmitting(true);

    onSubmit(text);

    setSubmitting(false);
    setSubmitted(true);
  }

  return (
    <div className={`${className} w-[400px] border-2 border-black bg-white rounded-lg p-4`}>
      <div className="flex justify-between items-center mb-4">
        <div className="font-bold text-2xl">Feedback</div>
        <button className="text-4xl" onClick={onClose}>
          <GrFormClose />
        </button>
      </div>

      {!submitted && (
        <form onSubmit={handleSubmit} className="flex flex-col">
          <div className="mb-4">
            <label>
              <div className="mb-4">Do you have any suggestions?</div>
              <textarea
                className="w-full border resize-none h-auto"
                placeholder="Enter your suggestions"
                value={text}
                onChange={(e) => setText(e.target.value)}
                rows={10}
                required
              />
            </label>
          </div>
          <Button type="submit" variant="primary-dark" className="ml-auto" disabled={submitting || submitted}>
            Submit
          </Button>
        </form>
      )}
      {submitted && (
        <span>Thanks for your feedback!</span>
      )}
    </div>
  );
};

export default FeedbackForm;