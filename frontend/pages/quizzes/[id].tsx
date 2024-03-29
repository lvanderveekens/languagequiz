
import FeedbackButton from '@/components/feedback-button';
import { QuizDto } from '@/components/models';
import Navbar from '@/components/navbar';
import Quiz from '@/components/quiz';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';

export default function QuizPage() {
  const router = useRouter();
  const quizId = router.query.id;

  const [quiz, setQuiz] = useState<QuizDto | null>(null);

  useEffect(() => {
    if (quizId) {
      fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/v1/quizzes/${quizId}`)
        .then((res) => res.json())
        .then((quiz) => {
          setQuiz(quiz);
        });
    }
  }, [quizId]);

  return (
    <div>
      <Navbar className="mb-8" />
      <FeedbackButton />

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
