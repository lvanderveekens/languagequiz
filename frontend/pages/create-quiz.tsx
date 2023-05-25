import Button from "@/components/button";
import languages from "@/components/languages";
import { CreateQuizRequest, ExerciseFormValues, QuizFormValues, QuizSectionFormValues } from "@/components/models";
import Navbar from "@/components/navbar";
import { useRouter } from 'next/router';
import { useEffect, useRef, useState } from "react";
import { v4 as uuidv4 } from 'uuid';
import "/node_modules/flag-icons/css/flag-icons.min.css";
import QuizSectionInput from "@/components/quiz-section-input";
import { FaExclamationCircle } from 'react-icons/fa';

const getInitialQuizSectionFormValues: () => QuizSectionFormValues = () => ({
  _key: uuidv4(),
  exercises: [],
});

const getInitialQuizFormValues: () => QuizFormValues = () => ({
  sections: [getInitialQuizSectionFormValues()],
});

export default function CreateQuizPage() {
  const [formValues, setFormValues] = useState<QuizFormValues>(getInitialQuizFormValues())
  const [errorMessage, setErrorMessage] = useState<string | null>(null)
  const router = useRouter();

  const errorContainerRef = useRef<HTMLDivElement>(null);

  const handleNameChange = (event: any) => {
    setFormValues({ ...formValues, name: event.target.value });
  };

  const handleLanguageChange = (event: any) => {
    setFormValues({ ...formValues, languageTag: event.target.value });
  };

  const resetForm = () => {
    setFormValues(getInitialQuizFormValues())
  }

  useEffect(() => {
    if (!errorMessage) {
      return;
    }
    if (errorContainerRef.current) {
      errorContainerRef.current.scrollIntoView({ behavior: "smooth" });
    }
  }, [errorMessage]);

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    console.log(`Submitting form with values: ${JSON.stringify(formValues)}`);
    setErrorMessage(null);

    try {
      const req = mapToRequest(formValues);
      console.log(`Request: ${JSON.stringify(req)}`);

      const res = await fetch(`/api/quizzes`, {
        method: "POST",
        body: JSON.stringify(req),
        headers: { "Content-Type": "application/json" },
      });

      const responseBody = await res.json();
      if (res.status == 201) {
        router.push(`/quizzes/${responseBody.id}`);
        resetForm();
      } else {
        setErrorMessage(responseBody.error);
      }
    } catch (error) {
      console.error(error);
    }
  };

  const mapToRequest = (formValues: QuizFormValues) => {
    const req: CreateQuizRequest = {
      languageTag: formValues.languageTag!,
      name: formValues.name!,
      sections: formValues.sections!.map((section) => ({
        name: section.name!,
        exercises: section.exercises!.map((exercise) => ({
          type: exercise.type!,
          question: exercise.question,
          choices: exercise.choices,
          sentence: exercise.sentence,
          correctedSentence: exercise.correctedSentence,
          answer: exercise.answer,
          feedback: exercise.feedback,
        })),
      })),
    };
    return req;
  }

  const handleSectionNameChange = (sectionIndex: any) => (sectionName: string) => {
    const updatedSections = [...(formValues.sections ?? [])]
    updatedSections[sectionIndex].name = sectionName

    const updatedFormValues = {...formValues, sections: updatedSections}
    setFormValues(updatedFormValues)
  }

  const handleExercisesChange = (sectionIndex: any) => (exercises: ExerciseFormValues[]) => {
    const updatedSections = [...(formValues.sections ?? [])]
    updatedSections[sectionIndex].exercises = exercises

    const updatedFormValues = {...formValues, sections: updatedSections}
    setFormValues(updatedFormValues)
  }

  const handleAddSection = () => {
    const newSection = getInitialQuizSectionFormValues(); 
    setFormValues({ ...formValues, sections: [...(formValues.sections ?? []), newSection] });
  };
  
  const handleRemoveSection = (index: number) => () => {
    setFormValues((prevState) => ({
      ...prevState,
      sections: [...(prevState.sections ?? [])].filter((_, i) => i !== index),
    }));

  };
  const getExerciseNumberStart = (sectionIndex: number) => {
    let exerciseNumberStart = 1
    let sections = formValues.sections ?? []
    for (let i = 0; i < sections.length; i++) {
      if (sectionIndex === i) {
        break;
      } else {
        exerciseNumberStart += sections[i].exercises?.length ?? 0
      }
    }
    return exerciseNumberStart
  };

  return (
    <div>
      <Navbar className="mb-8" />
      <div className="container mx-auto">
        <div className="max-w-screen-sm">
          <form onSubmit={handleSubmit}>
            <div className="text-2xl font-bold mb-8">
              <span className="mr-2">Create a quiz</span>
            </div>
            <div className="mb-4">
              <label>
                <div className="">Name</div>
                <input
                  className="w-full"
                  type="text"
                  placeholder="Enter a name"
                  value={formValues.name ?? ""}
                  onChange={handleNameChange}
                  required
                />
              </label>
            </div>
            <div className="mb-4">
              <label className="">
                <div className="">Language</div>
                <div className="flex">
                  <select className="w-full" value={formValues.languageTag} onChange={handleLanguageChange} required>
                    <option selected disabled value="">
                      Select a language
                    </option>
                    {languages
                      .sort((a, b) => a.name.localeCompare(b.name))
                      .map((language) => (
                        <option key={language.languageTag} value={language.languageTag}>
                          {language.name}
                        </option>
                      ))}
                  </select>
                </div>
              </label>
            </div>
            {formValues.sections &&
              formValues.sections.map((formValues: QuizSectionFormValues, i) => {
                return (
                  <QuizSectionInput
                    className="border mb-4 p-4"
                    index={i}
                    key={formValues._key}
                    name={formValues.name}
                    onNameChange={handleSectionNameChange(i)}
                    exercises={formValues.exercises}
                    onExercisesChange={handleExercisesChange(i)}
                    onRemove={i != 0 ? handleRemoveSection(i) : undefined}
                    exerciseNumberStart={getExerciseNumberStart(i)}
                  />
                );
              })}

            <div className="mb-8">
              <button
                type="button"
                className="border p-4 w-full px-3 flex items-center justify-center hover:border-black"
                onClick={handleAddSection}
              >
                <span className="text-2xl mr-1">➕</span>
                Add section
              </button>
            </div>

            <Button className="mb-8" variant="primary-dark" type="submit">
              Create
            </Button>
          </form>
          {errorMessage && (
            <div ref={errorContainerRef} className="mb-8">
              <span className="items-center px-4 py-2 border-2 border-red-400 rounded-lg bg-red-100 inline-flex">
                <span className="mr-2 text-xl text-red-500"><FaExclamationCircle /></span>
                Error: {errorMessage}
              </span>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
