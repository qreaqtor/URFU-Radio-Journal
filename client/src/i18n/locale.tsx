import React from "react";
import { LOCALES } from "./locales";

// костыль с вопросами
interface ILangContext {
  languageChange: (language: string) => void;
  currentLocale: string;
}

export const Context = React.createContext<ILangContext>({languageChange: () => {}, currentLocale: LOCALES.RUSSIAN});