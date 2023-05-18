import { Inter } from 'next/font/google'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { QuizDto, getNumberOfExercises } from '../components/models'
import { getLanguageByTag } from '@/components/languages'
import "/node_modules/flag-icons/css/flag-icons.min.css";

const inter = Inter({ subsets: ['latin'] })

export default function HomePage() {
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
      <nav className="bg-black py-4 mb-4">
        <div className="container text-xl text-white flex justify-between align-center">
          <div className='text-2xl'>
            LanguageQuiz
          </div>
          <button
            className="px-4 py-2 border-2 border-white px-3"
            onClick={() => router.push("/create-quiz")}
          >
            Create quiz
          </button>
        </div>
      </nav>

      <div className="container mx-auto">
        <div>
          <h2 className="text-2xl font-bold mb-4">Latest Quizzes</h2>
          <div className="grid grid-cols-3 gap-4">
            {quizzes.length > 0 &&
              quizzes.map((quiz, i) => {
                return (
                  <div key={quiz.id}>
                    <Link href={`/quizzes/${quiz.id}`}>
                      <div className="relative border border-black aspect-[3/2]">
                        <div
                          className={`opacity-10 absolute bg-cover bg-center w-full h-full fi-${
                            getLanguageByTag(quiz.languageTag)?.countryCode
                          }`}
                        ></div>
                        <div className="px-4 py-4">
                          <div>
                            <span className="font-bold">Name:</span> {quiz.name}
                          </div>
                          <div>
                            <span className="font-bold">Language:</span>{" "}
                            {getLanguageByTag(quiz.languageTag)?.name}
                          </div>
                          <div>
                            <span className="font-bold">Sections:</span>{" "}
                            {quiz.sections.length}
                          </div>
                          <div>
                            <span className="font-bold">Exercises:</span>{" "}
                            {getNumberOfExercises(quiz)}
                          </div>
                        </div>
                      </div>
                    </Link>
                  </div>
                );
              })}
          </div>
        </div>
      </div>
    </main>
  );
}
