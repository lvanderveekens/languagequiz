import React, { useEffect, useRef, useState } from 'react';
import { GrFormClose } from 'react-icons/gr';
import Button from './button';

type Props = {
  className?: string
  onSubmit: (text: any) => void
  onClose: () => void
};

const FeedbackModal: React.FC<Props> = ({ className = "", onSubmit, onClose }) => {
  const [text, setText] = useState<string>();
  const [submitting, setSubmitting] = useState<boolean>(false);
  const [submitted, setSubmitted] = useState<boolean>(false);

  const modalRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleKeyPress = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        onClose();
      }
    };

  const handleClickOutside = (event: MouseEvent) => {
    if (modalRef.current && !modalRef.current.contains(event.target as Node)) {
      onClose();
    }
  };

    document.addEventListener('keydown', handleKeyPress);
    document.addEventListener('click', handleClickOutside);

    return () => {
      document.removeEventListener('keydown', handleKeyPress);
      document.removeEventListener('click', handleClickOutside);
    };
  }, []);

  const handleSubmit = (event: any) => {
    event.preventDefault();
    setSubmitting(true);

    onSubmit(text);

    resetForm()
    setSubmitting(false);
    setSubmitted(true);
  };

  const resetForm = () => {
    setText("")
  }

  return (
    <div className={`${className} fixed inset-0 flex items-center justify-center z-50`}>
      <div className="fixed inset-0 bg-black opacity-50"></div> {/* Darkened background */}
      <div ref={modalRef} className="w-full h-full md:w-[400px] md:h-auto md:border-2 p-4 md:rounded-lg bg-white z-10">
        <div className="flex justify-between items-center mb-4">
          <div className="font-bold text-2xl">Feedback</div>
          <button className="text-4xl" onClick={onClose}>
            <GrFormClose />
          </button>
        </div>

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
          <Button type="submit" variant="primary-dark" className="ml-auto" disabled={submitting}>
            Submit
          </Button>
        </form>
        {submitted && <div className="mt-4">Thanks for your feedback! 😀</div>}
      </div>
    </div>
  );
};

export default FeedbackModal;