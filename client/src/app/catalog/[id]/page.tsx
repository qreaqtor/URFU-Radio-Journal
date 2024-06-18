// "use client"
import React from 'react';
import { useQuery } from '@tanstack/react-query';
import { getIssueById } from '@/data/IssueApi';
// import { getArticlesByIssueId } from '@/data/AticleApi';
import { FormattedMessage } from 'react-intl';
import Article from '@/components/Article';
import { IRuEng } from '@/lib/types';
import { useLanguageContext } from '@/i18n/languageContext';
import Head from 'next/head';
import { Metadata } from 'next';
import { frontUrl } from '@/lib/utils';

import IssuePageClient from '@/components/pages/IssuePage';

export async function generateMetadata({ params }: { params: { id: string } }): Promise<Metadata> {

  const issue = await getIssueById(params.id);
  const title = `Том ${issue.volume}, №${issue.number}(${issue.year})`;

  return {
    title: title,
    openGraph: {
      type: "website",
      url: `${frontUrl}/catalog/${params.id}`,
      title: title,
      siteName: "Ural Radio Enjeenering journal"
    }
  }
}

const IssuePage = ({ params }: { params: { id: string } }) => {

  return (
    <>
      <IssuePageClient params={params}/>
    </>
  );
};

export default IssuePage;