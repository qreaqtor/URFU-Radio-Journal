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

const ArticleSearch = ({article} : ArticleComponentProps) => {
  const {lang} = useLanguageContext();

  return (
    <div className="article-search-paragraph">
      <div className="article-search-paragraph__description">
        <Link href={`/catalog/article/${article.id}`} className='article-paragraph__name'>
          {article.title[lang as keyof IRuEng]}
        </Link>
        <p className='article-paragraph__authors'>{article.authors.map(x => x.fullname[lang as keyof IRuEng]).join(", ")}</p>
      </div>
    </div>
  );
};

export default ArticleSearch;