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

      const res = await fetch(`http://localhost:8888/v1/quizzes/${id}/answers`, {
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

  const getScore = (results: SubmitAnswerResult[]) => {
    let fraction = results.filter((result) => result.correct).length / results.length;
    const percentage = Math.ceil(fraction * 100);
    return `${percentage}%`;
  };

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
                    return <div key={exerciseIndex} className="mb-4">{exerciseComponent}</div>;
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
          {results && results.length > 0 && (
            <div className="mb-8">
              <span className="items-center px-4 py-2 border-2 border-[#cdcdcd] rounded-lg bg-[#eeeeee] inline-flex">
                <span className="text-4xl mr-2">🎓</span>
                Your score: {getScore(results)}
              </span>
            </div>
          )}
        </form>
      </div>
    </div>
  );
};

export default Quiz;