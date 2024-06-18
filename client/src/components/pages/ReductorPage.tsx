"use client"
import React from 'react';
import Image from 'next/image';
import { useQuery } from '@tanstack/react-query';
import { useLanguageContext } from '@/i18n/languageContext';
import { IRuEng } from '@/lib/types';
import { getCouncilById } from '@/data/CouncilApi';
import { serverUrl, transformDate } from '@/lib/utils';
import { getReductorById } from '@/data/ReductorService';


const ReductorPageClient = ({ params }: { params: { id: string } }) => {
  const {lang} = useLanguageContext();

  const {status: editorStatus, data: editor, error} = useQuery({
    queryFn: async () => await getReductorById(params.id),
    queryKey: ["editor", params.id],
    staleTime: Infinity
  });

  if (editorStatus == 'pending') {
    return <>Загрузка...</>
  }

  return (
    <div>
      <div className="editor-page">
        <div className="editor-page__block">
          <div className="editor-page__image">
            <Image className="editor__image-img" width={200} height={200} src={`${serverUrl}/api/v1/files/download/${editor!.imageID}`} alt="" />
          </div>
          <div className="editor-page__description">
            <p className='editor-page__name'>{editor?.name[lang as keyof IRuEng]}</p>
            <p className='editor-page__workplace'>{editor?.content[lang as keyof IRuEng]}</p>
            <p className='editor-page__workplace'>{editor?.location[lang as keyof IRuEng]}</p>
            <p className='editor-page__workplace'>Дата вступления в совет: {transformDate(editor!.dateJoin)}</p>
          </div>
        </div>

        <div className="editor-page__email-block">
          <div className="editor-page__email-block-container">
            <div className="editor-page__email">{editor?.email}</div>
          </div>
        </div>

        <div className="editor-page__text">
          <p>{editor?.description[lang as keyof IRuEng]}</p>
        </div>
      </div>
    </div>
  );
};

export default ReductorPageClient;