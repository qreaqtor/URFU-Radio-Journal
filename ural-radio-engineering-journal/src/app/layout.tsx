import type { Metadata } from "next";
import "./globals.css";
import Header from "../components/Header";
import Footer from "@/components/Footer";
import Head from "next/head";
import PageLayout from "../components/PageLayout"
import LanguageContextProvider from '../i18n/languageContext';
import QueryProvider from "@/lib/queryProvider";
import icon from '../../public/assets/_logo _urj 2.svg'


export const metadata: Metadata = {
  title: "Ural Radio Enjeenering journal",
  description: "Рецензируемый международный научно-технический журнал с открытым доступом, посвященный последним достижениям радиоэлектроники и связи.",
  openGraph: {
    type: "website",
    title: "Ural Radio Enjeenering journal",
    description: "Рецензируемый международный научно-технический журнал с открытым доступом, посвященный последним достижениям радиоэлектроники и связи."
  }
};


export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <div className="wrapper">
          <QueryProvider>
            <LanguageContextProvider>
              <Header/>
                <PageLayout>
                  {children}
                </PageLayout>
              <Footer/>  
            </LanguageContextProvider>
          </QueryProvider>
        </div>
      </body>
    </html>
  );
}
