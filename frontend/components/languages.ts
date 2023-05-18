interface Language {
  languageTag: string;
  name: string;
  countryCode: string;
}

const languages: Language[] = [
  {
    languageTag: "ar-SA",
    name: "Arabic",
    countryCode: "sa",
  },
  {
    languageTag: "cs-CZ",
    name: "Czech",
    countryCode: "cz",
  },
  {
    languageTag: "da-DK",
    name: "Danish",
    countryCode: "dk",
  },
  {
    languageTag: "de-DE",
    name: "German",
    countryCode: "de",
  },
  {
    languageTag: "en-US",
    name: "English",
    countryCode: "us",
  },
  {
    languageTag: "es-ES",
    name: "Spanish",
    countryCode: "es",
  },
  {
    languageTag: "fi-FI",
    name: "Finnish",
    countryCode: "fi",
  },
  {
    languageTag: "fr-FR",
    name: "French",
    countryCode: "fr",
  },
  {
    languageTag: "he-IL",
    name: "Hebrew",
    countryCode: "il",
  },
  {
    languageTag: "hi-IN",
    name: "Hindi",
    countryCode: "in",
  },
  {
    languageTag: "hu-HU",
    name: "Hungarian",
    countryCode: "hu",
  },
  {
    languageTag: "id-ID",
    name: "Indonesian",
    countryCode: "id",
  },
  {
    languageTag: "it-IT",
    name: "Italian",
    countryCode: "it",
  },
  {
    languageTag: "ja-JP",
    name: "Japanese",
    countryCode: "jp",
  },
  {
    languageTag: "ko-KR",
    name: "Korean",
    countryCode: "kr",
  },
  {
    languageTag: "nl-NL",
    name: "Dutch",
    countryCode: "nl",
  },
  {
    languageTag: "no-NO",
    name: "Norwegian",
    countryCode: "no",
  },
  {
    languageTag: "pl-PL",
    name: "Polish",
    countryCode: "pl",
  },
  {
    languageTag: "pt-PT",
    name: "Portuguese",
    countryCode: "pt",
  },
  {
    languageTag: "ro-RO",
    name: "Romanian",
    countryCode: "ro",
  },
  {
    languageTag: "ru-RU",
    name: "Russian",
    countryCode: "ru",
  },
  {
    languageTag: "sk-SK",
    name: "Slovak",
    countryCode: "sk",
  },
  {
    languageTag: "sv-SE",
    name: "Swedish",
    countryCode: "se",
  },
  {
    languageTag: "th-TH",
    name: "Thai",
    countryCode: "th",
  },
  {
    languageTag: "tr-TR",
    name: "Turkish",
    countryCode: "tr",
  },
  {
    languageTag: "zh-CN",
    name: "Chinese",
    countryCode: "cn",
  },
];

export const getLanguageByTag = (languageTag: string) => {
  return languages.find((l) => l.languageTag == languageTag);
};

export default languages;
