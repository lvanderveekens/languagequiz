import React, { useState } from 'react';
import { ChangeEvent } from 'react';
import { ExerciseDto, QuizSectionDto, SubmitAnswerResult, SubmitAnswersRequest, SubmitAnswersResponse } from './models';
import MultipleChoiceExercise from './multiple-choice-exercise';
import FillInTheBlankExercise from './fill-in-the-blank-exercise';
import SentenceCorrectionExercise from './sentence-correction-exercise';
import { getLanguageByTag } from './languages';

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
  const [exercises, setExercises] = useState<ExerciseDto[]>(sections.flatMap(section => section.exercises));
  const [answers, setAnswers] = useState<any[]>(Array.from({ length: exercises.length }, () => null));
  const [results, setResults] = useState<SubmitAnswerResult[]>();

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
    <div className="border border-black">
      <div className="font-bold">Quiz: {name}</div>
      <div>Language: {getLanguageByTag(languageTag)?.name}</div>
      <div>
        <form onSubmit={handleSubmit}>
          {sections.map((section: QuizSectionDto) => (
            <div key={section.name}>
              <div className="font-bold">Section: {section.name}</div>
              <div>
                {section.exercises.length > 0 &&
                  section.exercises.map((exercise, _) => {
                  exerciseIndex++;
                    switch (exercise.type) {
                      case "multipleChoice":
                        return (
                          <MultipleChoiceExercise
                            key={exercise.question!}
                            question={exercise.question!}
                            choices={exercise.choices!}
                            answer={answers[exerciseIndex]}
                            setAnswer={setAnswer(exerciseIndex)}
                          />
                        );
                      case "fillInTheBlank":
                        return (
                          <FillInTheBlankExercise
                            key={exercise.question!}
                            question={exercise.question!}
                            answer={answers[exerciseIndex]}
                            setAnswer={setAnswer(exerciseIndex)}
                          />
                        );
                      case "sentenceCorrection":
                        return (
                          <SentenceCorrectionExercise
                            key={exercise.sentence!}
                            sentence={exercise.sentence!}
                            answer={answers[exerciseIndex]}
                            setAnswer={setAnswer(exerciseIndex)}
                          />
                        );
                      default:
                        return <p>Unexpected exercise type: {exercise.type}</p>;
                    }
                  })
                  }
              </div>
            </div>
          ))}
          <button type="submit">Submit</button>
        </form>
        {results && <p>{JSON.stringify(results)}</p>}
      </div>
    </div>
  );
};

export default Quiz;