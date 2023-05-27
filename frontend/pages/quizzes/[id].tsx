
import { QuizDto } from '@/components/models';
import Navbar from '@/components/navbar';
import Quiz from '@/components/quiz';
import { Inter } from 'next/font/google';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';

export default function QuizPage() {
  const router = useRouter();
  const quizId = router.query.id;

  const [quiz, setQuiz] = useState<QuizDto | null>(null);

  useEffect(() => {
    if (quizId) {
      fetch(`http://localhost:8888/v1/quizzes/${quizId}`)
        .then((res) => res.json())
        .then((quiz) => {
          setQuiz(quiz);
        });
    }
  }, [quizId]);

  return (
    <div>
      <Navbar className="mb-8" />
      <div className="container mx-auto">
        <div className="max-w-screen-sm">
          {quiz && (
            <Quiz key={quiz.id} id={quiz.id} languageTag={quiz.languageTag} name={quiz.name} sections={quiz.sections} />
          )}
        </div>
      </div>
    </div>
  );
}
