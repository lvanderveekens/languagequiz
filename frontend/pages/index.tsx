import Image from 'next/image'
import { Inter } from 'next/font/google'
import { useEffect, useState } from 'react'
import MultipleChoiceExercise from './multiple-choice-exercise'
import FillInTheBlankExercise from './fill-in-the-blank-exercise'

const inter = Inter({ subsets: ['latin'] })

type Exercise = {
  id: string;
  type: string;
  question?: string;
  options?: string[];
};

export default function Home() {
  const [exercises, setExercises] = useState<Exercise[]>([]);

  useEffect(() => {
    fetch("/api/exercises")
      .then((res) => res.json())
      .then((data) => {
        setExercises(data);
      });
  }, []);

  return (
    <main>
      <div className="container mx-auto">
        {exercises.map((e) => {
          switch (e.type) {
            case "multipleChoice":
              return <MultipleChoiceExercise question={e.question!} options={e.options!} />;
            case "fillInTheBlank":
              return <FillInTheBlankExercise question={e.question!}  />;
            default:
              return <p>Unexpected exercise type: {e.type}</p>;
          }
        })}
      </div>
    </main>
  );
}
