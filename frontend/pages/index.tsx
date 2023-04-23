import Image from 'next/image'
import { Inter } from 'next/font/google'
import { useEffect, useState } from 'react'

const inter = Inter({ subsets: ['latin'] })

type Exercise = {
  id: string
  createdAt: string
  updatedAt: string
}

export default function Home() {
  const [data, setData] = useState<Exercise[]>([])

  useEffect(() => {
    fetch('/api/exercises')
      .then((res) => res.json())
      .then((data) => {
        setData(data)
      })
  }, [])

  return (
    <main>
      {data.map((exercise) => (
        <div>
          ID: {exercise.id}, CreatedAt: {exercise.createdAt}, UpdatedAt: {exercise.updatedAt}
        </div>
      ))}
    </main>
  );
}
