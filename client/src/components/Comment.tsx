import React from 'react';
import userImage from '../../public/assets/user.svg';
import { IComment, IRuEng } from '@/lib/typesNew';
import { useLanguageContext } from '@/i18n/languageContext';
import Image from 'next/image';
import { transformDate } from '@/lib/utils';

type CommentProps = {
  comment: IComment;
}

const Comment = ({comment}: CommentProps) => {
  const {lang} = useLanguageContext();
  const date = new Date(comment.date);

  return (
    <div className="article__comment">
      {/* <div className="comment__image">
        <Image src={userImage} alt="" />
      </div> */}
      <div className="comment_description">
        <p className="comment__author">{comment.author}</p>
        <p className="comment__text">{comment.content[lang as keyof IRuEng]}</p>
        <p className="comment__date">{transformDate(comment.date)}</p>
      </div>
    </div>
  );
};

export default Comment;