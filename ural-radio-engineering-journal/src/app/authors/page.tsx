"use client"
import { FormattedMessage } from 'react-intl';
import image from '../../../public/assets/photo_5269545055805365753_y.jpg'
import Image from 'next/image';
import { title } from 'process';
import { getAuthorsStatic } from '@/data/WordPressApi';
import { useQuery } from '@tanstack/react-query';
import { getByLang, useLanguageContext } from '@/i18n/languageContext';
// import { AwaitedReactNode, JSXElementConstructor, AwaitedReactNode, JSXElementConstructor, AwaitedReactNode, JSXElementConstructor, Key, ReactElement, ReactNode, ReactPortal, ReactElement, ReactNode, ReactPortal, ReactElement, ReactNode, ReactPortal } from 'react';



const ForAuthorspage = () => {
  const {lang} = useLanguageContext();
  
  const {status: articleStatus, data: staticAuthors, error: articleError} = useQuery({
    queryFn: async () => await getAuthorsStatic(),
    queryKey: ["authorsStatic"],
    staleTime: Infinity
  });

  if (articleStatus =='error') {
    return <>загрузка</>
  }

  if (articleStatus =='pending') {
    return <>загрузка</>
  }

  console.log(staticAuthors);
  return (
    <>
      <div className='image'></div>

      <div className="main__paragraph paragraph">
        <div className="paragraph__title"><p>{staticAuthors[lang.toLowerCase()][0][0].replaceAll('_', ' ')}</p></div>
        <div className="paragraph__text">
          <p>{staticAuthors[lang.toLowerCase()][0][1].text}</p>
        </div>
        <ol className="paragraph__list-ol">
          {staticAuthors[lang.toLowerCase()][0][1].files.map((item: string[], i: number) => (
            <li className='paragraph__item' key={i}>
              <a href={item[1]} download>
                {item[0].replaceAll('_', ' ')}
              </a>
            </li>
          ))}
        </ol>
      </div>

      <div className="main__paragraph paragraph">
        <div className="paragraph__title"><p>{staticAuthors[lang.toLowerCase()][1][0].replaceAll('_', ' ')}</p></div>
        <ul className="paragraph__list-ul paragraph__text_list">
          {staticAuthors[lang.toLowerCase()][1][1].list.map((item: string, i: number) => 
            <li className='paragraph__item' key={i}>
              {item}
            </li>
          )}
        </ul>
      </div>

      <div className="forauthors__image">
        <Image src={staticAuthors[lang.toLowerCase()][1][1].images[0][1]} width={400} height={1000} alt=""/>
      </div>
      

      <div className="main__paragraph paragraph">
        <div className="paragraph__title"><p>{staticAuthors[lang.toLowerCase()][2][0].replaceAll('_', ' ')}</p></div>
        <div className="paragraph__text">
          <p>{staticAuthors[lang.toLowerCase()][2][1].text}</p>
        </div>
        <div className="paragraph__text paragraph__text_list">
          {staticAuthors[lang.toLowerCase()][2][1].list_custom.map((item: string[], i: number) => 
            <p key={i}>
              <span className='paragraph__number'>{item[0]}</span>
              {item[1]}
            </p>
          )}
        </div>
      </div>

      <div className="main__paragraph paragraph">
        <div className="paragraph__title"><p>{staticAuthors[lang.toLowerCase()][3][0].replaceAll('_', ' ')}</p></div>
        <ol className="paragraph__list-ol">
        {staticAuthors[lang.toLowerCase()][3][1].list.map((item: string, i: number) => 
            <li className='paragraph__item' key={i}>
              {item}
            </li>
          )}
        </ol>
      </div>
    </>
  );
};

export default ForAuthorspage;