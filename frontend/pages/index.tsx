import { Inter } from 'next/font/google'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { QuizDto, getNumberOfExercises } from '../components/models'
import { getLanguageByTag } from '@/components/languages'
import "/node_modules/flag-icons/css/flag-icons.min.css";
import moment from 'moment';
import Button from '@/components/button'
import Navbar from '@/components/navbar'


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
      <div id="hero">
        <Navbar className="bg-transparent" />

        <div className="text-center">
          <div className="container mx-auto text-white">
            <div className="py-40">
              <div className="text-3xl font-bold mb-8">
                Language Quizzes Made Easy: Empowering Students, Supporting Teachers.
              </div>
              <div>
                <Button className="text-xl mb-4" variant="primary-light" onClick={() => router.push("#latest-quizzes")}>
                  Take a quiz
                </Button>
              </div>
              <div className="">
                <Button variant="secondary-light" className="text-xl" onClick={() => router.push("/create-quiz")}>
                  Create a quiz
                </Button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="container mx-auto text-black">
        <div id="latest-quizzes" className="my-8">
          <h2 className="text-2xl font-bold mb-8">Latest Quizzes</h2>
          <div className="grid grid-cols-4 gap-4">
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
                        <div className="border border-2 border-black hover:text-white hover:bg-[#003259] hover:border-[#003259] text-black relative rounded-lg aspect-[3/2]">
                          <div className="px-4 py-4">
                            <div>
                              <span className="font-bold">Name:</span> {quiz.name}
                            </div>
                            <div>
                              <span className="font-bold">Created:</span> {moment(quiz.createdAt).fromNow()}
                            </div>
                            <div>
                              <span className="font-bold">Language: </span>
                              <span className={`mr-1 fi fi-${getLanguageByTag(quiz.languageTag)?.countryCode}`} />
                              <span className="">{getLanguageByTag(quiz.languageTag)?.name}</span>
                            </div>
                            <div>
                              <span className="font-bold">Sections:</span> {quiz.sections.length}
                            </div>
                            <div>
                              <span className="font-bold">Exercises:</span> {getNumberOfExercises(quiz)}
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
