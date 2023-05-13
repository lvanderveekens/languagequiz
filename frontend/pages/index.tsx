import { Inter } from 'next/font/google'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { QuizDto } from '../components/models'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
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
    <main>
      <div className="container mx-auto">
        Quizzes:
        {quizzes.length > 0 &&
          quizzes.map((quiz, i) => {
            return (
              <div key={quiz.id}>
                <Link href={`/quizzes/${quiz.id}`}>{quiz.name}</Link>
              </div>
            );
          })}
        <button className="border border-black px-3" onClick={() => router.push("/create-quiz")}>Create quiz</button>
      </div>
    </main>
  );
}
