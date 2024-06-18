import React from 'react';
import { getArticleById } from '@/data/AticleApi';
import { transformDate } from '@/lib/utils';
import ArticlePageClient from '@/components/pages/ArticlePage';
import { Metadata } from 'next';
import { frontUrl } from '@/lib/utils';

export async function generateMetadata({ params }: { params: { id: string } }): Promise<Metadata> {
  const article = await getArticleById(params.id);
  const keywords = article?.keywords.map(x => x.Ru).join(', ');

  const title = article?.title.Ru;
  const description = `${article?.content.Ru}. Авторы: ${article?.authors.map(x => x.fullname.Ru).join(', ')}. Дата поступления: ${transformDate(article?.dateReceipt)}, Дата принятия: ${transformDate(article?.dateAcceptance)}`;
  return {
    title: title,
    description: description,
    keywords, 
    openGraph: {
      type: "website",
      url: `${frontUrl}/catalog/article/${params.id}`,
      title: title,
      description: description,
      siteName: "Ural Radio Enjeenering Journal"
    }
  }
}

const ArticlePage = ({ params }: { params: { id: string } }) => {

  return (
    <>
      <ArticlePageClient params={params}/>
    </>
  );
};

export default ArticlePage;