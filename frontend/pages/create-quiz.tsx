
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { QuizDto } from '../components/models'

export default function CreateQuizPage() {
  const router = useRouter();

  const [quizzes, setQuizzes] = useState<QuizDto[]>([]);

  useEffect(() => {
    fetch("/api/quizzes")
      .then((res) => res.json())
      .then((quizzes) => {
        setQuizzes(quizzes);
      })
  }, []);

  return (
      <div className="container mx-auto">
        Create quiz page
      </div>
  );
}
