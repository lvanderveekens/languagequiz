
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

export interface CreateQuizRequest {
    name: string
    sections: CreateQuizSectionDto[]
}

export interface CreateQuizSectionDto {
    name: string
    exercises: CreateExerciseDto[]
}

export interface CreateExerciseDto {
  type: string;
  question?: string;
  choices?: string[];
  sentence?: string;
  correctedSentence?: string;
  answer?: string
}

export enum ExerciseType {
  MultipleChoice = "multipleChoice",
  FillInTheBlank = "fillInTheBlank",
  SentenceCorrection = "sentenceCorrection",
}