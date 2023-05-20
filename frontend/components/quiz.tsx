import React, { useState } from 'react';
import FillInTheBlankExercise from './fill-in-the-blank-exercise';
import { getLanguageByTag } from './languages';
import { ExerciseDto, QuizSectionDto, SubmitAnswerResult, SubmitAnswersRequest, SubmitAnswersResponse } from './models';
import MultipleChoiceExercise from './multiple-choice-exercise';
import SentenceCorrectionExercise from './sentence-correction-exercise';
import "/node_modules/flag-icons/css/flag-icons.min.css";
import Button from './button';

type Props = {
  id: string
  languageTag: string
  name: string
  sections: QuizSectionDto[]
};

const Quiz: React.FC<Props> = ({
  id,
  languageTag,
  name,
  sections,
}) => {
  const [exercises] = useState<ExerciseDto[]>(sections.flatMap(section => section.exercises));
  const [answers, setAnswers] = useState<any[]>(Array.from({ length: exercises.length }, () => null));
  const [results, setResults] = useState<SubmitAnswerResult[]>();
  const [submitted, setSubmitted] = useState<boolean>(false);

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    console.log(answers)

    try {
      const req: SubmitAnswersRequest = {
        userAnswers: exercises.map((exercise, i) => {
          return answers[i];
        }),
      };

      const res = await fetch(`/api/quizzes/${id}/answers`, {
        method: "POST",
        body: JSON.stringify(req),
        headers: { "Content-Type": "application/json" },
      });

      const resBody = await res.json() as SubmitAnswersResponse
      setResults(resBody.results)
    } catch (error) {
      console.error(error);
    }

    setSubmitted(true);
  };

  const resetForm = () => {
    setAnswers(Array.from({ length: exercises.length }, () => null));
    setResults([]);
    setSubmitted(false);
  };

  const setAnswer = (index: number) => {
    return (updatedAnswer: any) => {
      setAnswers((prevAnswers) =>
        prevAnswers.map((prevAnswer, i) => {
          if (i !== index) {
            return prevAnswer;
          }
          return updatedAnswer;
        })
      );
    };
  };

  let exerciseIndex = -1;

  return (
    <div className="">
      <div className="text-2xl font-bold mb-8">
        <span className="mr-2">{name}</span>
      </div>

      <div>
        <form onSubmit={handleSubmit}>
          {sections.map((section: QuizSectionDto, sectionIndex) => (
            <div key={section.name} className="mb-8">
              <div className="text-xl font-bold mb-4">
                {String.fromCharCode(65 + sectionIndex)}. {section.name}
              </div>
              <div>
                {section.exercises.length > 0 &&
                  section.exercises.map((exercise, _) => {
                    exerciseIndex++;
                    let exerciseComponent;

                    switch (exercise.type) {
                      case "multipleChoice":
                        exerciseComponent = (
                          <MultipleChoiceExercise
                            key={exercise.question!}
                            index={exerciseIndex}
                            question={exercise.question!}
                            choices={exercise.choices!}
                            answer={answers[exerciseIndex]}
                            setAnswer={setAnswer(exerciseIndex)}
                            correctAnswer={results?.[exerciseIndex]?.answer}
                            feedback={results?.[exerciseIndex]?.feedback}
                            disabled={submitted}
                          />
                        );
                        break;
                      case "fillInTheBlank":
                        exerciseComponent = (
                          <FillInTheBlankExercise
                            key={exercise.question!}
                            index={exerciseIndex}
                            question={exercise.question!}
                            answer={answers[exerciseIndex]}
                            setAnswer={setAnswer(exerciseIndex)}
                            correctAnswer={results?.[exerciseIndex]?.answer}
                            feedback={results?.[exerciseIndex]?.feedback}
                            disabled={submitted}
                          />
                        );
                        break;
                      case "sentenceCorrection":
                        exerciseComponent = (
                          <SentenceCorrectionExercise
                            key={exercise.sentence!}
                            index={exerciseIndex}
                            sentence={exercise.sentence!}
                            answer={answers[exerciseIndex]}
                            setAnswer={setAnswer(exerciseIndex)}
                            correctAnswer={results?.[exerciseIndex]?.answer}
                            feedback={results?.[exerciseIndex]?.feedback}
                            disabled={submitted}
                          />
                        );
                        break;
                      default:
                        exerciseComponent = <p>Unexpected exercise type: {exercise.type}</p>;
                    }
                    return <div className="mb-4">{exerciseComponent}</div>;
                  })}
              </div>
            </div>
          ))}
          <Button type="submit" variant="primary-dark" className="mb-8" disabled={submitted}>
            Check
          </Button>
          {submitted && (
            <Button type="button" className="mb-8 ml-3" variant="secondary-dark" onClick={resetForm}>
              Reset
            </Button>
          )}
        </form>
      </div>
    </div>
  );
};

export default Quiz;