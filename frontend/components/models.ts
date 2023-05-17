
export interface QuizDto {
  id: string
  languageTag: string
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
  languageTag: string;
  name: string;
  sections: CreateQuizSectionRequest[];
}

export interface CreateQuizSectionRequest {
  name: string;
  exercises: CreateExerciseRequest[];
}

export interface CreateExerciseRequest {
  type: string;
  question?: string;
  choices?: string[];
  sentence?: string;
  correctedSentence?: string;
  answer?: string;
  feedback?: string;
}

export enum ExerciseType {
  MultipleChoice = "multipleChoice",
  FillInTheBlank = "fillInTheBlank",
  SentenceCorrection = "sentenceCorrection",
}