import React from 'react';
import { FormattedMessage } from 'react-intl'
import { IArticle, IRuEng } from '@/lib/typesNew';
import Link from 'next/link';
import { useQuery } from '@tanstack/react-query';
import { useLanguageContext } from '@/i18n/languageContext';
import { serverUrl } from '@/lib/utils';

type ArticleComponentProps = {
  article: IArticle
}

const Article = ({article} : ArticleComponentProps) => {
  const {lang} = useLanguageContext();

  return (
    <div className="article-paragraph">
    <a target='blank' href={`${serverUrl}/api/v1/files/download/${article?.documentID}`} download={true}>
      <button className='article-paragraph__button'><FormattedMessage id='article-article-paragraph__button' /></button>
    </a>
    <div className="article-paragraph__description">
      <Link href={`/catalog/article/${article.id}`} className='article-paragraph__name'>
        {article.title[lang as keyof IRuEng]}
      </Link>
      <p className='article-paragraph__authors'>{article.authors.map(x => x.fullname[lang as keyof IRuEng]).join(", ")}</p>
    </div>
  </div>
  );
};

export default Article;