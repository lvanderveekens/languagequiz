import Button from "@/components/button";
import languages from "@/components/languages";
import { CreateQuizRequest, ExerciseFormValues, QuizFormValues, QuizSectionFormValues } from "@/components/models";
import Navbar from "@/components/navbar";
import { useRouter } from 'next/router';
import { useState } from "react";
import { v4 as uuidv4 } from 'uuid';
import "/node_modules/flag-icons/css/flag-icons.min.css";
import QuizSectionInput from "@/components/quiz-section-input";

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

  const handleNameChange = (event: any) => {
    setFormValues({ ...formValues, name: event.target.value });
  };

  const handleLanguageChange = (event: any) => {
    setFormValues({ ...formValues, languageTag: event.target.value });
  };

  const resetForm = () => {
    setFormValues(getInitialQuizFormValues())
  }

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
        resetForm();
        router.push(`/quizzes/${responseBody.id}`);
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

  let exerciseCounterStart = 0;

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
                    exerciseCounterStart={exerciseCounterStart}
                    key={formValues._key}
                    name={formValues.name}
                    onNameChange={handleSectionNameChange(i)}
                    exercises={formValues.exercises}
                    onExercisesChange={handleExercisesChange(i)}
                    onRemove={i != 0 ? handleRemoveSection(i) : undefined}
                  />
                );
              })}

            <div className="mb-4">
              <button
                type="button"
                className="border p-4 w-full px-3 flex items-center justify-center hover:border-black"
                onClick={handleAddSection}
              >
                <span className="text-2xl mr-1">➕</span>
                Add section
              </button>
            </div>

            <Button className="mb-4" variant="primary-dark" type="submit">
              Create
            </Button>
          </form>
          {errorMessage && <div>Error: {errorMessage}</div>}
        </div>
      </div>
    </div>
  );
}
