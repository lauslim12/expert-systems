import i18n from 'i18next';
import LanguageDetector from 'i18next-browser-languagedetector';
import Backend from 'i18next-http-backend';
import { initReactI18next } from 'react-i18next';

// Settings from: https://react.i18next.com/latest/using-with-hooks.
i18n
  // Load translations with HTTP.
  .use(Backend)

  // Checks the language which the user prefers.
  .use(LanguageDetector)

  // Pass i18n instance.
  .use(initReactI18next)

  // Initialize all. Translations are fetched automatically from 'public/locales/<lang_code>/translation.json'.
  .init({
    fallbackLng: 'en',
    debug: process.env.NODE_ENV === 'development',
    interpolation: {
      escapeValue: false,
    },
  });

export default i18n;
