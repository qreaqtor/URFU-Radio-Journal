// import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { IArticle } from '../service/types/typesNew';

type ArticleComponentProps = {
  article: IArticle
}

const ArticleComponent = ({article}: ArticleComponentProps) => {

  return (
    <div className="article-paragraph">      
      <div className="article-paragraph__description">
        <Link to={`/issues/article/${article["id"]}`} className='article-paragraph__name'>
          {article["title"]["Ru"]}
        </Link>
        <p className='article-paragraph__authors'>{article.authors && article.authors.map(x => x.fullname.Ru).join(", ")}</p>
      </div>
    </div>
  );
};

export default ArticleComponent;