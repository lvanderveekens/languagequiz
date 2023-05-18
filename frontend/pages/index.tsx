import { Inter } from 'next/font/google'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { QuizDto, getNumberOfExercises } from '../components/models'
import { getLanguageByTag } from '@/components/languages'
import "/node_modules/flag-icons/css/flag-icons.min.css";
import moment from 'moment';


const inter = Inter({ subsets: ['latin'] })

export default function HomePage() {
  const router = useRouter();

  const [quizzes, setQuizzes] = useState<QuizDto[]>([]);

  useEffect(() => {
    fetch("/api/quizzes")
      .then((res) => res.json())
      .then((quizzes) => {
        setQuizzes(quizzes);
      });
  }, []);

  return (
    <main>
      <nav className="py-4 mb-4">
        <div className="container text-xl text-white flex justify-between align-center">
          <div className="text-2xl font-bold self-center">LanguageQuiz</div>
        </div>
      </nav>

      <div className="container mx-auto text-white">
        <div className="text-center my-40">
          <div className="text-3xl font-bold mb-8">
            Language Quizzes Made Easy: Empowering Students, Supporting Teachers.
          </div>
          <div>
            <button
              className="text-xl bg-[#0066FF] border-2 border-[#0066FF] px-4 py-2 rounded-lg px-3 mb-4"
              onClick={() => router.push("#latest-quizzes")}
            >
              Take a quiz
            </button>
          </div>
          <div className="">
            <button
              className="text-xl px-4 py-2 border-2 border-white rounded-lg px-3"
              onClick={() => router.push("/create-quiz")}
            >
              Create a quiz
            </button>
          </div>
       </div> 
        <div id="latest-quizzes" className="my-24">
          <h2 className="text-white text-2xl font-bold mb-8">Latest Quizzes</h2>
          <div className="grid grid-cols-3 gap-4">
            {quizzes.length > 0 &&
              quizzes
                .sort((a, b) => {
                  const dateA = new Date(a.createdAt);
                  const dateB = new Date(b.createdAt);
                  return dateB.getTime() - dateA.getTime();
                })
                .map((quiz, i) => {
                  return (
                    <div key={quiz.id}>
                      <Link href={`/quizzes/${quiz.id}`}>
                        <div className="bg-white text-black relative rounded-lg aspect-[3/2]">
                          <div className="px-4 py-4">
                            <div>
                              <span className="font-bold">Name:</span>{" "}
                              {quiz.name}
                            </div>
                            <div>
                              <span className="font-bold">Created:</span>{" "}
                              {moment(quiz.createdAt).fromNow()}
                            </div>
                            <div>
                              <span className="font-bold">Language: </span>
                              <span
                                className={`mr-1 fi fi-${
                                  getLanguageByTag(quiz.languageTag)
                                    ?.countryCode
                                }`}
                              />
                              <span className="">
                                {getLanguageByTag(quiz.languageTag)?.name}
                              </span>
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
