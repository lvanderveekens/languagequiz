
import { QuizDto } from '@/components/models';
import Quiz from '@/components/quiz';
import { Inter } from 'next/font/google'
import Link from 'next/link';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
// import Quiz from '../quiz';
// import { QuizDto } from '../components/models';

const inter = Inter({ subsets: ['latin'] })

export default function QuizPage() {
  const router = useRouter();
  const quizId = router.query.id;

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
    <div>
      <nav className="py-4 b-4 bg-[#003259] mb-8">
        <div className="container text-xl text-white flex justify-between align-center">
          <div className="text-2xl font-bold self-center">
            <Link href="/">LanguageQuiz</Link>
          </div>
        </div>
      </nav>
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
    </div>
  );
}
