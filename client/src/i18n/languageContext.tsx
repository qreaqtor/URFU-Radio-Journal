"use client"
import React, { useState, useContext } from "react";
import { IntlProvider } from "react-intl";
import { messages } from "./messages";
import { LOCALES } from "./locales";

interface ILangContext {
  localeChange: (language: string) => void;
  locale: string;
  lang: string;
}

const LanguageContext = React.createContext<ILangContext>({localeChange: () => {}, locale: LOCALES.RUSSIAN, lang: "Ru"});

export function useLanguageContext() {
  const context = useContext(LanguageContext);
  return context;
}

export function getByLang(lang: string, first: string, second: string) {
  return lang == "Ru" ? first : second;
}

export default function LanguageContextProvider({children} : {children: React.ReactNode}) {
  const [lang, setLang] = useState<"Eng" | "Ru">('Ru');
  const [locale, setLocale] = useState(LOCALES.RUSSIAN);

  const localeChange = (inputLocale: string): void => {
    setLocale(inputLocale);
    setLang(inputLocale == "en-US" ? "Eng" : "Ru");
  }

  return (
    <IntlProvider messages={messages[locale]} locale={locale} defaultLocale={LOCALES.ENGLISH}>
      <LanguageContext.Provider value={{localeChange, locale, lang}}>
        {children}
      </LanguageContext.Provider>
    </IntlProvider>
  );
}