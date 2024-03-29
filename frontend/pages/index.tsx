import Button from '@/components/button'
import FeedbackButton from '@/components/feedback-button'
import { getLanguageByTag } from '@/components/languages'
import Navbar from '@/components/navbar'
import moment from 'moment'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { QuizDto, getNumberOfExercises } from '../components/models'
import "/node_modules/flag-icons/css/flag-icons.min.css"

export default function HomePage() {
  const router = useRouter();

  const [quizzes, setQuizzes] = useState<QuizDto[]>([]);

  useEffect(() => {
    fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/v1/quizzes`)
      .then((res) => res.json())
      .then((quizzes) => {
        setQuizzes(quizzes);
      });
  }, []);

  return (
    <main>
      <div id="hero">
        <Navbar className="bg-transparent" />
        <FeedbackButton />

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
          <div className="grid sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
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
                      <div className="border border-2 text-black relative rounded-lg aspect-[3/2]">
                        <div className="px-4 py-4">
                          <div>
                            <span className="font-bold">Name:</span> {quiz.name}
                          </div>
                          <div>
                            <span className="font-bold">Created:</span> {moment(quiz.createdAt).fromNow()}
                          </div>
                          <div>
                            <span className="font-bold">Language: </span>
                            <span
                              className={`border box-content mr-1 fi fi-${
                                getLanguageByTag(quiz.languageTag)?.countryCode
                              }`}
                            />
                            <span className="">{getLanguageByTag(quiz.languageTag)?.name}</span>
                          </div>
                          <div>
                            <span className="font-bold">Exercises:</span> {getNumberOfExercises(quiz)}
                          </div>
                          <div className="mt-2">
                            <Link href={`/quizzes/${quiz.id}`}>
                              <Button type="button" variant="primary-dark">
                                Take quiz
                              </Button>
                            </Link>
                          </div>
                        </div>
                      </div>
                    </div>
                  );
                })}
          </div>
        </div>
      </div>
    </main>
  );
}
