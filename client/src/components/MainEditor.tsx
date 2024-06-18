import React from 'react';
import Link from 'next/link';
import Image from 'next/image';
import { IEditor, IRuEng } from '@/lib/types';
import { useLanguageContext } from '@/i18n/languageContext';
import { useQuery } from '@tanstack/react-query';
import { ICouncil } from '@/lib/typesNew';
import { serverUrl } from '@/lib/utils';

type MainEditorProps = {
  editor: ICouncil;
}

const MainEditor = ({editor} : MainEditorProps) => {
  const {lang} = useLanguageContext();

  // const {status: editorStatus, data: image} = useQuery({
  //   queryFn: async () => await getFileLink(editor["imagePathId"]),
  //   queryKey: ["image"],
  // });

  return (
    <div className="editors__editor editor">
      <div className="editor__block">
        <div className="editor__image">
          <Link href={`/editors/${editor.id}`}>
            <Image className="editor__image-img" width={200} height={200} src={`${serverUrl}/api/v1/files/download/${editor.imageID}`} alt="" />
          </Link>
        </div>
        <div className="editor__description">
          <p className='editor__name'>{editor.name[lang as keyof IRuEng]}</p>
          <p className='editor__workplace'>{editor.location[lang as keyof IRuEng]}</p>
        </div>
      </div>
    </div>
  );
};

export default MainEditor;