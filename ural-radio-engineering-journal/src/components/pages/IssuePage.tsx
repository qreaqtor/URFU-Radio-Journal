"use client"
import React from 'react';
import { useQuery } from '@tanstack/react-query';
import { getIssueById } from '@/data/IssueApi';
// import { getArticlesByIssueId } from '@/data/AticleApi';
import { getArticles } from '@/data/AticleApi';
import { FormattedMessage } from 'react-intl';
import Article from '@/components/Article';
import { IRuEng } from '@/lib/types';
import { useLanguageContext } from '@/i18n/languageContext';
import Head from 'next/head';
import { serverUrl } from '@/lib/utils';

const IssuePageClient = ({ params }: { params: { id: string } }) => {
  const {lang} = useLanguageContext();
  const { data: issue, status: issueStatus, error } = useQuery({
    queryKey: ['issue', params.id],
    queryFn: async () => await getIssueById(params.id)
  });

  const {status: articlesStatus, data: articles} = useQuery({
    queryFn: async () => await getArticles(params.id),
    queryKey: ["articles", params.id],
    staleTime: Infinity
  });

  return (
    <div>
      <p className='issue__title'>Том {issue?.volume}, №{issue?.number}({issue?.year})</p>
      <a href={`${serverUrl}/api/v1/files/download/${issue?.documentID}`} download={true}>
        <button className='issue__button'>
          <FormattedMessage id='issue-issue__button' />
        </button>
      </a>
      
      <p className='issue__title'><FormattedMessage id='issue-issue__title' /></p>

      {articlesStatus == 'success' && articles?.map(article => <Article key={article["id"]} article={article}/>)}
    </div>
  );
};

export default IssuePageClient;