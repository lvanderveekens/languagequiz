interface Language {
    languageTag: string
    name: string
}

const languages: Language[] = [
    {
      languageTag: "ar-SA",
      name: "Arabic"
    },
    {
      languageTag: "cs-CZ",
      name: "Czech"
    },
    {
      languageTag: "da-DK",
      name: "Danish"
    },
    {
      languageTag: "de-DE",
      name: "German"
    },
    {
      languageTag: "en-US",
      name: "English"
    },
    {
      languageTag: "es-ES",
      name: "Spanish"
    },
    {
      languageTag: "fi-FI",
      name: "Finnish"
    },
    {
      languageTag: "fr-FR",
      name: "French"
    },
    {
      languageTag: "he-IL",
      name: "Hebrew"
    },
    {
      languageTag: "hi-IN",
      name: "Hindi"
    },
    {
      languageTag: "hu-HU",
      name: "Hungarian"
    },
    {
      languageTag: "id-ID",
      name: "Indonesian"
    },
    {
      languageTag: "it-IT",
      name: "Italian"
    },
    {
      languageTag: "ja-JP",
      name: "Japanese"
    },
    {
      languageTag: "ko-KR",
      name: "Korean"
    },
    {
      languageTag: "nl-NL",
      name: "Dutch"
    },
    {
      languageTag: "no-NO",
      name: "Norwegian"
    },
    {
      languageTag: "pl-PL",
      name: "Polish"
    },
    {
      languageTag: "pt-PT",
      name: "Portuguese"
    },
    {
      languageTag: "ro-RO",
      name: "Romanian"
    },
    {
      languageTag: "ru-RU",
      name: "Russian"
    },
    {
      languageTag: "sk-SK",
      name: "Slovak"
    },
    {
      languageTag: "sv-SE",
      name: "Swedish"
    },
    {
      languageTag: "th-TH",
      name: "Thai"
    },
    {
      languageTag: "tr-TR",
      name: "Turkish"
    },
    {
      languageTag: "zh-CN",
      name: "Chinese"
    }
  ];

  export const getLanguageByTag = (languageTag: string) => {
    return languages.find((l) => l.languageTag == languageTag)
  }
  
  export default languages;
  