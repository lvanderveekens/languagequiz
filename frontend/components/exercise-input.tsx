import { ExerciseFormValues, ExerciseType } from "@/components/models";
import { FaInfoCircle } from "react-icons/fa";
import { GrFormClose } from "react-icons/gr";
import "/node_modules/flag-icons/css/flag-icons.min.css";

type ExerciseInputProps = {
  value: ExerciseFormValues;
  onChange: (value: ExerciseFormValues) => void;
  onRemove?: () => void;
  exerciseNumber: number;
};

const ExerciseInput: React.FC<ExerciseInputProps> = ({
  value,
  onChange,
  onRemove,
  exerciseNumber,
}) => {
  const handleRemoveClick = (event: any) => {
    event.preventDefault();
    onRemove!();
  };

  let exerciseInputComponent = null
  switch (value.type) {
    case ExerciseType.MultipleChoice:
      exerciseInputComponent = (
        <MultipleChoiceExerciseInput
          question={value.question}
          choices={value.choices}
          answer={value.answer}
          feedback={value.feedback}
          onQuestionChange={(question: string) => onChange({ ...value, question: question })}
          onChoicesChange={(choices: string[]) => onChange({ ...value, choices: choices })}
          onAnswerChange={(answer: string) => onChange({ ...value, answer: answer })}
          onFeedbackChange={(feedback: string) => onChange({ ...value, feedback: feedback })}
        />
      );
      break;
    case ExerciseType.FillInTheBlank:
      exerciseInputComponent = (
        <FillInTheBlankExerciseInput
          question={value.question}
          answer={value.answer}
          feedback={value.feedback}
          onQuestionChange={(question: string) => onChange({ ...value, question: question })}
          onAnswerChange={(answer: string) => onChange({ ...value, answer: answer })}
          onFeedbackChange={(feedback: string) => onChange({ ...value, feedback: feedback })}
        />
      );
      break;
    case ExerciseType.SentenceCorrection:
        exerciseInputComponent = (
          <SentenceCorrectionExerciseInput
            sentence={value.sentence}
            correctedSentence={value.correctedSentence}
            feedback={value.feedback}
            onSentenceChange={(sentence: string) => onChange({ ...value, sentence: sentence })}
            onCorrectedSentenceChange={(correctedSentence: string) =>
              onChange({ ...value, correctedSentence: correctedSentence })
            }
            onFeedbackChange={(feedback: string) => onChange({ ...value, feedback: feedback })}
          />
        );
        break;
    default:
      return <p>Unknown exercise type</p>;
  }

  return (
    <div className="mb-4 p-4 border">
      {onRemove == null && <div className="mb-4 font-bold">Exercise {exerciseNumber}</div>}
      {onRemove != null && (
        <div className="flex justify-between items-center mb-4">
          <div className="font-bold">Exercise {exerciseNumber}</div>
          <button className="text-3xl" onClick={handleRemoveClick}>
            <GrFormClose />
          </button>
        </div>
      )}
      <div>{exerciseInputComponent}</div>
    </div>
  );

}

type MultipleChoiceExerciseInputProps = {
  question?: string;
  choices?: string[];
  answer?: string;
  feedback?: string;
  onQuestionChange: (question: string) => void;
  onChoicesChange: (choices: string[]) => void;
  onAnswerChange: (answer: string) => void;
  onFeedbackChange: (feedback: string) => void;
};

const MultipleChoiceExerciseInput: React.FC<MultipleChoiceExerciseInputProps> = ({
  question,
  choices,
  answer,
  feedback,
  onQuestionChange,
  onChoicesChange,
  onAnswerChange,
  onFeedbackChange,
}) => {

  const handleChoiceChange = (index: number) => (event: any) => {
    const { value } = event.target;
    const updatedChoices = [...(choices ?? new Array(4))]
    updatedChoices[index] = value
    onChoicesChange(updatedChoices);
  };

  return (
    <div className="">
      <div className="mb-4">
        <label>
          <div className="">Question</div>
          <input
            className="w-full border"
            placeholder="Enter a question"
            type="text"
            value={question ?? ""}
            onChange={(e) => onQuestionChange(e.target.value)}
            required
          />
        </label>
      </div>
      <div className="mb-4">
        <label>
          <div className="">Choice 1</div>
          <input
            className="w-full border"
            placeholder="Enter first choice"
            type="text"
            value={choices?.[0] ?? ""}
            onChange={handleChoiceChange(0)}
            required
          />
        </label>
      </div>
      <div className="mb-4">
        <label>
          <div className="">Choice 2</div>
          <input
            className="w-full border"
            placeholder="Enter second choice"
            type="text"
            value={choices?.[1] ?? ""}
            onChange={handleChoiceChange(1)}
            required
          />
        </label>
      </div>
      <div className="mb-4">
        <label>
          <div className="">Choice 3</div>
          <input
            className="w-full border"
            placeholder="Enter third choice"
            type="text"
            value={choices?.[2] ?? ""}
            onChange={handleChoiceChange(2)}
            required
          />
        </label>
      </div>
      <div className="mb-4">
        <label>
          <div className="">Choice 4</div>
          <input
            className="w-full border"
            placeholder="Enter fourth choice"
            type="text"
            value={choices?.[3] ?? ""}
            onChange={handleChoiceChange(3)}
            required
          />
        </label>
      </div>
      <div className="mb-4">
        <label>
          <div className="mr-3">Answer</div>
          <input
            className="w-full border"
            placeholder="Enter the answer"
            type="text"
            value={answer ?? ""}
            onChange={(e) => onAnswerChange(e.target.value)}
            required
          />
        </label>
      </div>
      <div>
        <label>
          <div className="mr-3">Feedback (optional)</div>
          <input
            className="w-full border"
            placeholder="Enter feedback"
            type="text"
            value={feedback ?? ""}
            onChange={(e) => onFeedbackChange(e.target.value)}
          />
        </label>
      </div>
    </div>
  );
}

type FillInTheBlankExerciseInputProps = {
  question?: string
  answer?: string
  feedback?: string
  onQuestionChange: (question: string) => void
  onAnswerChange: (answer: string) => void
  onFeedbackChange: (feedback: string) => void
};

const FillInTheBlankExerciseInput: React.FC<FillInTheBlankExerciseInputProps> = ({
  feedback,
  question,
  answer,
  onQuestionChange,
  onAnswerChange,
  onFeedbackChange,
}) => {

  return (
    <div className="">
      <div className="mb-4">
        <label>
          <div className="">Question</div>
          <input
            className="w-full border"
            placeholder="Enter a question"
            type="text"
            value={question ?? ""}
            onChange={(e) => onQuestionChange(e.target.value)}
            required
          />
        </label>
        <div className="mt-4">
          <span className="items-center px-4 py-2 border-2 rounded-lg inline-flex">
            <span className="mr-2 text-xl">
              <FaInfoCircle />
            </span>
            The question should include a blank (______).
          </span>
        </div>
      </div>
      <div className="mb-4">
        <label>
          <div className="">Answer</div>
          <input
            className="w-full border"
            placeholder="Enter an answer"
            type="text"
            value={answer ?? ""}
            onChange={(e) => onAnswerChange(e.target.value)}
            required
          />
        </label>
      </div>
      <div className="mb-4">
        <label>
          <div className="">Feedback (optional)</div>
          <input
            className="w-full border"
            placeholder="Enter feedback"
            type="text"
            value={feedback ?? ""}
            onChange={(e) => onFeedbackChange(e.target.value)}
          />
        </label>
      </div>
    </div>
  );
}

type SentenceCorrectionExerciseInputProps = {
  sentence?: string;
  correctedSentence?: string;
  feedback?: string;
  onSentenceChange: (sentence: string) => void;
  onCorrectedSentenceChange: (correctedSentence: string) => void;
  onFeedbackChange: (feedback: string) => void;
};

const SentenceCorrectionExerciseInput: React.FC<SentenceCorrectionExerciseInputProps> = ({
  feedback,
  sentence,
  correctedSentence,
  onSentenceChange,
  onCorrectedSentenceChange,
  onFeedbackChange,
}) => {

  return (
    <div className="">
      <div className="mb-4">
        <label>
          <span className="">Sentence</span>
          <input
            className="w-full border"
            placeholder="Enter a sentence"
            type="text"
            value={sentence ?? ""}
            onChange={(e) => onSentenceChange(e.target.value)}
            required
          />
        </label>
      </div>
      <div className="mb-4">
        <label>
          <span className="">Corrected sentence</span>
          <input
            className="w-full border"
            placeholder="Enter the corrected sentence"
            type="text"
            value={correctedSentence ?? ""}
            onChange={(e) => onCorrectedSentenceChange(e.target.value)}
            required
          />
        </label>
      </div>
      <div className="mb-4">
        <label>
          <span className="">Feedback (optional)</span>
          <input
            className="w-full border"
            placeholder="Enter feedback"
            type="text"
            value={feedback ?? ""}
            onChange={(e) => onFeedbackChange(e.target.value)}
          />
        </label>
      </div>
    </div>
  );
}

export default ExerciseInput;