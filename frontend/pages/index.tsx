import Image from 'next/image'
import { Inter } from 'next/font/google'
import { useEffect, useState } from 'react'
import MultipleChoiceExercise from './multiple-choice-exercise'
import FillInTheBlankExercise from './fill-in-the-blank-exercise'
import SentenceCorrectionExercise from './sentence-correction-exercise'
import Quiz from './quiz'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  const [exercises, setExercises] = useState<Exercise[]>([]);
  const [quizzes, setQuizzes] = useState<Quiz[]>([]);

  useEffect(() => {
    fetch("/api/quizzes")
      .then((res) => res.json())
      .then((quizzes) => {
        setQuizzes(quizzes);
      })
  }, []);

  return (
    <main>
      <div className="container mx-auto">
        {quizzes.length > 0 && (
          quizzes.map((quiz, i) => {
            return <Quiz key={quiz.id} id={quiz.id} name={quiz.name} sections={quiz.sections} />;
          })
        )}
      </div>
    </main>
  );
}

export interface Quiz {
  id: string
  name: string
  sections: QuizSection[]
}

export interface QuizSection {
  name: string
  exercises: Exercise[]
}

export interface Exercise {
  id: string
  type: string
  question?: string
  choices?: string[]
  sentence?: string
}

export interface SubmitAnswersRequest {
  userAnswers: any[]
}

export interface SubmitAnswersResponse{
  results: SubmitAnswerResult[]
}

export interface SubmitAnswerResult {
  correct: boolean
  answer: any
}