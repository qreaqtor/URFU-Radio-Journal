import React from 'react';
import { IAuthor, IRuEng } from '@/lib/typesNew';
import { useLanguageContext } from '@/i18n/languageContext';

type AuthorProps = {
  author: IAuthor
}

const Author = ({author}: AuthorProps) => {
  const {lang} = useLanguageContext();
  return (
    <div className='article__author'>
      <p>{author.fullname[lang as keyof IRuEng]}</p>
      <p>{author.affilation}</p>
      <p>{author.email}</p>
    </div>
  );
};

export default Author;