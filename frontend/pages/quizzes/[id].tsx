
import { QuizDto } from '@/components/models';
import Quiz from '@/components/quiz';
import { Inter } from 'next/font/google'
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
// import Quiz from '../quiz';
// import { QuizDto } from '../components/models';

const inter = Inter({ subsets: ['latin'] })

export default function QuizPage() {
  const router = useRouter();
  const quizId = router.query.id

  const [quiz, setQuiz] = useState<QuizDto | null>(null);

  useEffect(() => {
    if (quizId) {
      fetch(`/api/quizzes/${quizId}`)
        .then((res) => res.json())
        .then((quiz) => {
          setQuiz(quiz);
        });
    }
  }, [quizId]);

  return (
    <div className="container mx-auto">
      {quiz && (
        <Quiz
          key={quiz.id}
          id={quiz.id}
          languageTag={quiz.languageTag}
          name={quiz.name}
          sections={quiz.sections}
        />
      )}
    </div>
  );
}
