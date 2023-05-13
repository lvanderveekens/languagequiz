
export interface QuizDto {
  id: string
  name: string
  sections: QuizSectionDto[]
}

export interface QuizSectionDto {
  name: string
  exercises: ExerciseDto[]
}

export interface ExerciseDto {
  id: string
  type: string
  question?: string
  choices?: string[]
  sentence?: string
}

export interface SubmitAnswersRequest {
  userAnswers: any[]
}

export interface SubmitAnswersResponse{
  results: SubmitAnswerResult[]
}

export interface SubmitAnswerResult {
  correct: boolean
  answer: any
}