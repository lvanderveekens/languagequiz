import { CreateQuizRequest, ExerciseType } from "@/components/models";
import { useState } from "react";
import { v4 as uuidv4 } from 'uuid';
import { useRouter } from 'next/router';
import languages from "@/components/languages";
import Navbar from "@/components/navbar";

const initialQuizFormValues: QuizFormValues = {
  sections: [
    {
      _key: uuidv4(),
      exercises: [
        {
          _key: uuidv4(),
        },
      ],
    },
  ],
};

export default function CreateQuizPage() {
  const [formValues, setFormValues] = useState<QuizFormValues>(initialQuizFormValues)
  const [errorMessage, setErrorMessage] = useState<string | null>(null)
  const router = useRouter();

  const handleNameChange = (event: any) => {
    setFormValues({ ...formValues, name: event.target.value });
  };

  const handleLanguageChange = (event: any) => {
    setFormValues({ ...formValues, languageTag: event.target.value });
  };

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    console.log(`Submitting form with values: ${JSON.stringify(formValues)}`);
    setErrorMessage(null);

    try {
      const req = mapToRequest(formValues);
      console.log(`Request: ${JSON.stringify(req)}`);

      const res = await fetch(`/api/quizzes`, {
        method: "POST",
        body: JSON.stringify(req),
        headers: { "Content-Type": "application/json" },
      });

      if (res.status == 201) {
        router.push("/");
      } else {
        const responseBody = await res.json();
        setErrorMessage(responseBody.error);
      }
    } catch (error) {
      console.error(error);
    }
  };

  const mapToRequest = (formValues: QuizFormValues) => {
    const req: CreateQuizRequest = {
      languageTag: formValues.languageTag!,
      name: formValues.name!,
      sections: formValues.sections!.map((section) => ({
        name: section.name!,
        exercises: section.exercises!.map((exercise) => ({
          type: exercise.type!,
          question: exercise.question,
          choices: exercise.choices,
          sentence: exercise.sentence,
          correctedSentence: exercise.correctedSentence,
          answer: exercise.answer,
          feedback: exercise.feedback,
        })),
      })),
    };
    return req;
  }

  const handleSectionNameChange = (sectionIndex: any) => (sectionName: string) => {
    const updatedSections = [...(formValues.sections ?? [])]
    updatedSections[sectionIndex].name = sectionName

    const updatedFormValues = {...formValues, sections: updatedSections}
    setFormValues(updatedFormValues)
  }

  const handleExercisesChange = (sectionIndex: any) => (exercises: ExerciseFormValues[]) => {
    const updatedSections = [...(formValues.sections ?? [])]
    updatedSections[sectionIndex].exercises = exercises

    const updatedFormValues = {...formValues, sections: updatedSections}
    setFormValues(updatedFormValues)
  }

  const handleAddSectionClick = () => {
    const newSection = {_key: uuidv4()}
    setFormValues({ ...formValues, sections: [...(formValues.sections ?? []), newSection] });
  };

  return (
    <div>
      <Navbar className="mb-8" />
      <div className="container mx-auto">
        <div className="text-2xl font-bold mb-8">
          <span className="mr-2">Create quiz</span>
        </div>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label>
              <div className="font-bold">Name</div>
              <input
                className=""
                type="text"
                placeholder="Enter a name"
                value={formValues.name ?? ""}
                onChange={handleNameChange}
                required
              />
            </label>
          </div>
          <div className="mb-4">
            <label>
              <div className="font-bold">Language</div>
              <select
                className="invalid:text-gray-400"
                placeholder="Select a language"
                value={formValues.languageTag}
                onChange={handleLanguageChange}
                required
              >
                <option selected disabled value="">
                  Select a language
                </option>
                {languages
                  .sort((a, b) => a.name.localeCompare(b.name))
                  .map((language) => (
                    <option key={language.languageTag} value={language.languageTag}>
                      {language.name}
                    </option>
                  ))}
              </select>
            </label>
          </div>
          {formValues.sections &&
            formValues.sections.map((formValues: QuizSectionFormValues, i) => (
              <QuizSectionInput
                key={formValues._key}
                name={formValues.name}
                onNameChange={handleSectionNameChange(i)}
                exercises={formValues.exercises}
                onExercisesChange={handleExercisesChange(i)}
              />
            ))}

          <div>
            <button type="button" className="border border-black px-3" onClick={handleAddSectionClick}>
              Add section
            </button>
          </div>
          <div>
            <button className="border border-black px-3" type="submit">
              Submit
            </button>
          </div>
        </form>
        {errorMessage && <div>Error: {errorMessage}</div>}
      </div>
    </div>
  );
}

type QuizSectionInputProps = {
  name?: string
  onNameChange: (name: string) => void
  exercises?: ExerciseFormValues[]
  onExercisesChange: (exercises: ExerciseFormValues[]) => void
};

const QuizSectionInput: React.FC<QuizSectionInputProps> = ({
  name,
  onNameChange,
  exercises,
  onExercisesChange,
}) => {

  const handleAddExerciseClick = () => {
    const newExercise = { _key: uuidv4() };
    onExercisesChange([...(exercises ?? []), newExercise]);
  };

  const handleExerciseChange =
    (index: number) => (value: ExerciseFormValues) => {
      const updatedExercises = [...(exercises ?? [])];
      updatedExercises[index] = value;

      onExercisesChange(updatedExercises);
    };

  return (
    <div className="border border-black">
      Section
      <div>
        <label>
          <span className="mr-3">Name</span>
          <input
            className="border border-black"
            type="text"
            value={name ?? ""}
            onChange={(e) => onNameChange(e.target.value)}
            required
          />
        </label>
      </div>
      {exercises &&
        exercises.map((formValues: ExerciseFormValues, i) => (
          <ExerciseInput key={formValues._key} value={exercises[i]} onChange={handleExerciseChange(i)} />
        ))}
      <div>
        <button
          type="button"
          className="border border-black px-3"
          onClick={handleAddExerciseClick}
        >
          Add exercise
        </button>
      </div>
    </div>
  );
}

type ExerciseInputProps = {
  value: ExerciseFormValues
  onChange: (value: ExerciseFormValues) => void
};

const ExerciseInput: React.FC<ExerciseInputProps> = ({
  value,
  onChange,
}) => {

  const handleTypeChange = (event: any) => {
    onChange({...value, type: event.target.value})
  };

  if (!value.type) {
    return (
      <div className="border border-black">
        Exercise
        <div>
          <label htmlFor="choice">Exercise type:</label>
          <select
            id="choice"
            name="choice"
            value={value.type}
            onChange={handleTypeChange}
            required
          >
            <option value="">Select an option</option>
            {Object.values(ExerciseType)
              .filter((key) => isNaN(Number(key)))
              .map((exerciseType) => (
                <option key={exerciseType} value={exerciseType}>{exerciseType}</option>
              ))}
          </select>
        </div>
      </div>
    );
  }

  switch (value.type) {
    case ExerciseType.MultipleChoice:
      return (
        <MultipleChoiceExerciseInput
          question={value.question}
          choices={value.choices}
          answer={value.answer}
          feedback={value.feedback}
          onQuestionChange={(question: string) =>
            onChange({ ...value, question: question })
          }
          onChoicesChange={(choices: string[]) =>
            onChange({ ...value, choices: choices })
          }
          onAnswerChange={(answer: string) =>
            onChange({ ...value, answer: answer })
          }
          onFeedbackChange={(feedback: string) =>
            onChange({ ...value, feedback: feedback })
          }
        />
      );
    case ExerciseType.FillInTheBlank:
      return (
        <FillInTheBlankExerciseInput
          question={value.question}
          answer={value.answer}
          feedback={value.feedback}
          onQuestionChange={(question: string) =>
            onChange({ ...value, question: question })
          }
          onAnswerChange={(answer: string) =>
            onChange({ ...value, answer: answer })
          }
          onFeedbackChange={(feedback: string) =>
            onChange({ ...value, feedback: feedback })
          }
        />
      );
    case ExerciseType.SentenceCorrection:
      return <SentenceCorrectionExerciseInput 
          sentence={value.sentence}
          correctedSentence={value.correctedSentence}
          feedback={value.feedback}
          onSentenceChange={(sentence: string) =>
            onChange({ ...value, sentence: sentence})
          }
          onCorrectedSentenceChange={(correctedSentence: string) =>
            onChange({ ...value, correctedSentence: correctedSentence })
          }
          onFeedbackChange={(feedback: string) =>
            onChange({ ...value, feedback: feedback })
          }
      />
    default:
      return <p>Unknown exercise type</p>
  }
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
    <div className="border border-black">
      Multiple Choice Exercise
      <div>
        <label>
          <span className="mr-3">Question</span>
          <input
            className="border border-black"
            type="text"
            value={question ?? ""}
            onChange={(e) => onQuestionChange(e.target.value)}
            required
          />
        </label>
      </div>
      <div>
        <label>
          <span className="mr-3">Choice 1</span>
          <input
            className="border border-black"
            type="text"
            value={choices?.[0] ?? ""}
            onChange={handleChoiceChange(0)}
            required
          />
        </label>
      </div>
      <div>
        <label>
          <span className="mr-3">Choice 2</span>
          <input
            className="border border-black"
            type="text"
            value={choices?.[1] ?? ""}
            onChange={handleChoiceChange(1)}
            required
          />
        </label>
      </div>
      <div>
        <label>
          <span className="mr-3">Choice 3</span>
          <input
            className="border border-black"
            type="text"
            value={choices?.[2] ?? ""}
            onChange={handleChoiceChange(2)}
          />
        </label>
      </div>
      <div>
        <label>
          <span className="mr-3">Choice 4</span>
          <input
            className="border border-black"
            type="text"
            value={choices?.[3] ?? ""}
            onChange={handleChoiceChange(3)}
          />
        </label>
      </div>
      <div>
        <label>
          <span className="mr-3">Answer</span>
          <input
            className="border border-black"
            type="text"
            value={answer ?? ""}
            onChange={(e) => onAnswerChange(e.target.value)}
            required
          />
        </label>
      </div>
      <div>
        <label>
          <span className="mr-3">Feedback</span>
          <input
            className="border border-black"
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
    <div className="border border-black">
      Fill In The Blank Exercise
      <div>
        <label>
          <span className="mr-3">Question</span>
          <input
            className="border border-black"
            type="text"
            value={question ?? ""}
            onChange={(e) => onQuestionChange(e.target.value)}
            required
          />
        </label>
      </div>
      <div>
        <label>
          <span className="mr-3">Answer</span>
          <input
            className="border border-black"
            type="text"
            value={answer ?? ""}
            onChange={(e) => onAnswerChange(e.target.value)}
            required
          />
        </label>
      </div>
      <div>
        <label>
          <span className="mr-3">Feedback</span>
          <input
            className="border border-black"
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
    <div className="border border-black">
      Sentence Correction Exercise
      <div>
        <label>
          <span className="mr-3">Sentence</span>
          <input
            className="border border-black"
            type="text"
            value={sentence ?? ""}
            onChange={(e) => onSentenceChange(e.target.value)}
          />
        </label>
      </div>
      <div>
        <label>
          <span className="mr-3">Corrected sentence</span>
          <input
            className="border border-black"
            type="text"
            value={correctedSentence ?? ""}
            onChange={(e) => onCorrectedSentenceChange(e.target.value)}
          />
        </label>
      </div>
      <div>
        <label>
          <span className="mr-3">Feedback</span>
          <input
            className="border border-black"
            type="text"
            value={feedback ?? ""}
            onChange={(e) => onFeedbackChange(e.target.value)}
          />
        </label>
      </div>
    </div>
  );
}

interface QuizFormValues {
    languageTag?: string
    name?: string
    sections?: QuizSectionFormValues[]
}

interface QuizSectionFormValues {
    _key: string
    name?: string
    exercises?: ExerciseFormValues[]
}

interface ExerciseFormValues {
  _key: string;
  type?: string;
  question?: string;
  choices?: string[];
  sentence?: string;
  correctedSentence?: string;
  answer?: string
  feedback?: string
}
