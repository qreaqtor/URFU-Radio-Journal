"use client"
import { getAboutJournalStatic } from "@/data/WordPressApi";
import { useLanguageContext } from "@/i18n/languageContext";
import { useQuery } from "@tanstack/react-query";
import { FormattedMessage } from "react-intl";


export default function MainPage() {
  const {lang} = useLanguageContext();

  const {status: articleStatus, data: staticAbout, error: articleError} = useQuery({
    queryFn: async () => await getAboutJournalStatic(),
    queryKey: ["aboutStatic"],
    staleTime: Infinity
  });

  if (articleStatus =='pending') {
    return <>загрузка</>
  }

  return (
    <>
      <div className="main__paragraph paragraph">
        <div className="paragraph__title"><p>{staticAbout[lang.toLowerCase()][0][0].replaceAll('_', ' ')}</p></div>
        <div className="paragraph__text">
          <p>{staticAbout[lang.toLowerCase()][0][1].text1}</p>
          <p>{staticAbout[lang.toLowerCase()][0][1].text2}</p>
          <p>{staticAbout[lang.toLowerCase()][0][1].text3}</p>
        </div>
      </div>

      <div className="main__paragraph paragraph">
        <div className="paragraph__title"><p>{staticAbout[lang.toLowerCase()][1][0].replaceAll('_', ' ')}</p></div>
        <div className="paragraph__text">
          {staticAbout[lang.toLowerCase()][1][1].list_custom.map((item: string[], i: number) => 
            <p key={i}>
              <span className='paragraph__number'>{item[0]}</span> 
              {item[1]}
            </p>
          )}
          {/* <p><span className='paragraph__number'>2.2.2.</span> <FormattedMessage id='mainpage-paragraph__text3' /></p>
          <p><span className='paragraph__number'>2.2.8.</span> <FormattedMessage id='mainpage-paragraph__text4' /></p>
          <p><span className='paragraph__number'>2.2.13.</span> <FormattedMessage id='mainpage-paragraph__text5' /></p>
          <p><span className='paragraph__number'>2.2.14.</span> <FormattedMessage id='mainpage-paragraph__text6' /></p>
          <p><span className='paragraph__number'>2.2.15.</span> <FormattedMessage id='mainpage-paragraph__text7' /></p>
          <p><span className='paragraph__number'>2.2.16.</span> <FormattedMessage id='mainpage-paragraph__text8' /></p> */}
        </div>
        <div className="paragraph__text">
          <p>{staticAbout[lang.toLowerCase()][2][1].text1}</p>
          <p>{staticAbout[lang.toLowerCase()][2][1].text2}</p>
          <p>{staticAbout[lang.toLowerCase()][2][1].text3}</p>
        </div>
      </div>
    </>
  )
}
