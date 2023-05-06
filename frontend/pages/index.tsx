import Image from 'next/image'
import { Inter } from 'next/font/google'
import { useEffect, useState } from 'react'
import MultipleChoiceExercise from './multiple-choice-exercise'
import FillInTheBlankExercise from './fill-in-the-blank-exercise'
import SentenceCorrectionExercise from './sentence-correction-exercise'

const inter = Inter({ subsets: ['latin'] })

type Exercise = {
  id: string;
  type: string;
  question?: string;
  choices?: string[];
  sentence?: string;
};

export default function Home() {
  const [exercises, setExercises] = useState<Exercise[]>([]);
  const [answers, setAnswers] = useState<any[]>([]);

  const [results, setResults] = useState<boolean[]>();

  useEffect(() => {
    fetch("/api/exercises")
      .then((res) => res.json())
      .then((exercises) => {
        setExercises(exercises);
        setAnswers(Array.from({ length: exercises.length }, () => null));
      });
  }, []);

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    console.log(answers)

    try {
      const req: SubmitAnswersRequest = {
        submissions: exercises.map((exercise, i) => {
          return { exerciseId: exercise.id, answer: answers[i] };
        }),
      };

      const res = await fetch("/api/submit-answers", {
        method: "POST",
        body: JSON.stringify(req),
        headers: { "Content-Type": "application/json" },
      });

      const resBody = await res.json() as SubmitAnswersResponse
      setResults(resBody.results.map((result) => result.correct))
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

  return (
    <main>
      <div className="container mx-auto">
        {exercises.length > 0 && (
          <form onSubmit={handleSubmit}>
            {exercises.map((exercise, i) => {
              switch (exercise.type) {
                case "multipleChoice":
                  return (
                    <MultipleChoiceExercise
                      key={exercise.id}
                      question={exercise.question!}
                      choices={exercise.choices!}
                      answer={answers[i]}
                      setAnswer={setAnswer(i)}
                    />
                  );
                case "fillInTheBlank":
                  return (
                    <FillInTheBlankExercise
                      key={exercise.id}
                      question={exercise.question!}
                      answer={answers[i]}
                      setAnswer={setAnswer(i)}
                    />
                  );
                case "sentenceCorrection":
                  return (
                    <SentenceCorrectionExercise
                      key={exercise.id}
                      sentence={exercise.sentence!}
                      answer={answers[i]}
                      setAnswer={setAnswer(i)}
                    />
                  );
                default:
                  return <p>Unexpected exercise type: {exercise.type}</p>;
              }
            })}

            <button type="submit">Submit</button>
          </form>
        )}
        {results && (
          <p>{JSON.stringify(results)}</p>
        )}
      </div>
    </main>
  );
}

interface SubmitAnswersRequest {
  submissions: ExerciseSubmission[]
}

interface ExerciseSubmission {
  exerciseId: string
  answer: any
}

interface SubmitAnswersResponse{
  results: ExerciseResult[]
}

interface ExerciseResult {
  exerciseId: string
  exerciseType: string
  correct: boolean
  answer: any
}