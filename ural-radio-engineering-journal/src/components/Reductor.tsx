import React from 'react';
import Image from 'next/image';
import Link from 'next/link';
import { IReductor, IRuEng } from '@/lib/typesNew';
import { useLanguageContext } from '@/i18n/languageContext';
import { serverUrl } from '@/lib/utils';

type BoardEditorProps = {
  editor: IReductor;
}

const BoardEditor = ({editor}: BoardEditorProps) => {
  const {lang} = useLanguageContext();

  return (
    <div>
      <div className="editors__board-element">
        <div className="editors__board-image">
          <Link  href={`/reductors/${editor.id}`}>
            <Image className="editor__image-img" width={200} height={200} src={`${serverUrl}/api/v1/files/download/${editor.imageID}`} alt="" />
          </Link>
        </div>
        <p className='editors__board-name'>{editor.name[lang as keyof IRuEng]}</p>
      </div>
    </div>
  );
};

export default BoardEditor;