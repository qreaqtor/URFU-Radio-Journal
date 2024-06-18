"use client"
import React, { useEffect, useMemo, useState } from 'react';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { getArticlesBySearch, getIssues } from '@/data/IssueApi';
import Issue from '@/components/Issue';
import SearchInput from '@/components/SearchInput';
import ArticleSearch from '@/components/ArticleSearch';
import { IArticle, IIssue } from '@/lib/typesNew';
import { getArticlesSearch } from '@/data/AticleApi';
import { FormattedMessage } from 'react-intl';

const CatalogPage = () => {
  const [update, setUpdate] = useState('')
  const [isAppdating, setIsAppdating] = useState(false);
  const [isSearching, setIsSearching] = useState(false);

  const queryClient = useQueryClient();

  const [offset, setOffset] = useState(0);
  const [limit, setlimit] = useState(9);
  const [issuesList, setIssuesList] = useState<IIssue[]>([])

  const [offsetArticle, setOffsetArticle] = useState(0);
  const [limitArticle, setlimitArticle] = useState(10);
  const [articlesList, setArticlesList] = useState<IArticle[]>([])

  const [title, setTitle] = useState(true);
  const [annotation, setAnnotation] = useState(false);
  const [authors, setAuthors] = useState(false);
  const [keywords, setKeywords] = useState(false);



  const {data: issues, isLoading, error } = useQuery({
    queryFn: async () => {
      const response = await getIssues(offset, limit)
      setIssuesList([...issuesList, ...response.data]);
      setOffset(offset + limit);
      return {...response, data: [...issuesList, ...response.data]}
    }, 
    queryKey: ["issues"],
  });

  const { data: articlesSearch, status: articlesSearchStatus, error: articlesSearchError } = useQuery({
    queryKey: ['articlesSearch'],
    queryFn: async () => {
      if (update == '') {
        return {allCount: 0, data: []};
      }
      const titleSearch = title || annotation || authors || keywords;
      const response = await getArticlesSearch(update, offsetArticle, limitArticle, title || !titleSearch, annotation, authors, keywords)
      setArticlesList([...articlesList, ...response.data]);
      setOffsetArticle(offsetArticle + limitArticle);
      
      return {...response, data: [...articlesList, ...response.data]};
    },
    staleTime: Infinity
  });

  const onButtonBackToIssues = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault();
    setUpdate('');
    setIsAppdating(true);
    setIsSearching(false);
  }

  const handleButtonAddIssues = () => {
    queryClient.invalidateQueries({queryKey: ['issues']});
  }
  const handleButtonAddArticles = () => {
    queryClient.invalidateQueries({queryKey: ['articlesSearch']});
  }



  useEffect(() => {
    if (isAppdating) {
      queryClient.invalidateQueries({queryKey: ["articlesSearch"]});
      setIsAppdating(false);
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isAppdating]);

  useEffect(() => {
    if (issues) {
      setOffset(issues.data.length);
      setIssuesList([...issuesList, ...issues.data]);
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // useEffect(() => {
  //   if (articlesSearch) {
  //     setOffsetArticle(articlesSearch.data.length);
  //     setArticlesList([...articlesList, ...articlesSearch.data]);
  //   }
  // // eslint-disable-next-line react-hooks/exhaustive-deps
  // }, []);

  if (error) {
    return <p>Произошла ошибка</p>
  }
  if (isLoading) {
    return <p>Загрузка...</p>
  }


  if (isSearching) {
    return (
      <>
        <SearchInput setUpdate={setUpdate} setIsAppdating={setIsAppdating} setIsSearching={setIsSearching} onButtonBackToIssues={onButtonBackToIssues}/>
        <div className="catalog__search-buttons">
          <p><FormattedMessage id='catalog-catalog__item-button-search-text' />:</p>
          <button className={!title? 'catalog__search-button' : 'catalog__search-button-pressed'} onClick={() => setTitle(!title)} type='button'><FormattedMessage id='catalog-catalog__item-button-search1' /></button>
          <button className={!annotation? 'catalog__search-button' : 'catalog__search-button-pressed'} onClick={() => setAnnotation(!annotation)} type='button'><FormattedMessage id='catalog-catalog__item-button-search2' /></button>
          <button className={!authors? 'catalog__search-button' : 'catalog__search-button-pressed'} onClick={() => setAuthors(!authors)} type='button'><FormattedMessage id='catalog-catalog__item-button-search3' /></button>
          <button className={!keywords? 'catalog__search-button' : 'catalog__search-button-pressed'} onClick={() => setKeywords(!keywords)} type='button'><FormattedMessage id='catalog-catalog__item-button-search4' /></button>
        </div>
        {/* <button className='search__back-button' onClick={onButtonBackToIssues} type='button'>Вернуться к выпускам</button> */}
        {articlesSearchStatus == 'success' && articlesSearch.data?.map(article => <ArticleSearch key={article["id"]} article={article}/>)}
        {articlesSearchStatus == 'success' && articlesSearch.data.length == 0 && <p>Ничего не найдено</p>}
        {articlesSearch && articlesSearch!.data.length < articlesSearch!.allCount && <button className='catalog__add-issues-button' onClick={() => handleButtonAddArticles()}>Загрузить ещё</button>}
      </>
    )
  }


  return (
    <>
      <SearchInput setUpdate={setUpdate} setIsAppdating={setIsAppdating} setIsSearching={setIsSearching}/>
      
      <div className="catalog__search-buttons">
          <p><FormattedMessage id='catalog-catalog__item-button-search-text' />:</p>
          <button className={!title? 'catalog__search-button' : 'catalog__search-button-pressed'} onClick={() => setTitle(!title)} type='button'><FormattedMessage id='catalog-catalog__item-button-search1' /></button>
          <button className={!annotation? 'catalog__search-button' : 'catalog__search-button-pressed'} onClick={() => setAnnotation(!annotation)} type='button'><FormattedMessage id='catalog-catalog__item-button-search2' /></button>
          <button className={!authors? 'catalog__search-button' : 'catalog__search-button-pressed'} onClick={() => setAuthors(!authors)} type='button'><FormattedMessage id='catalog-catalog__item-button-search3' /></button>
          <button className={!keywords? 'catalog__search-button' : 'catalog__search-button-pressed'} onClick={() => setKeywords(!keywords)} type='button'><FormattedMessage id='catalog-catalog__item-button-search4' /></button>
        </div>
      <div className="catalog">
        {issues!.data.map((issue, id) => {
          return <Issue key={id} issue={issue}/>
        })}
      </div>
      {issues!.data.length < issues!.allCount && <button className='catalog__add-issues-button' onClick={() => handleButtonAddIssues()}>Загрузить ещё</button>}
    </>
  );
};

export default CatalogPage;