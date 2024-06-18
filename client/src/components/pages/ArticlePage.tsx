'use client'
import React, { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { getArticleById } from '@/data/AticleApi';
import { getComments, createComment } from '@/data/CommentApi';
import { IComment, IRuEng } from '@/lib/typesNew';
import { FormattedMessage } from 'react-intl';
import { useLanguageContext } from '@/i18n/languageContext';
import Comment from '@/components/Comment';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import Author from '../Author';
import { serverUrl, transformDate } from '@/lib/utils';
import { downloadFile } from '@/data/FileApi';


const commentSchema = z.object({
  author: z.string().min(1, "Комментарий не должен быть пустым"),
  comment: z.string().min(1, "Комментарий не должен быть пустым")
});

type ICommentSchema = z.infer<typeof commentSchema>;

const ArticlePageClient = ({ params }: { params: { id: string } }) => {
  const [offset, setOffset] = useState(0);
  const [limit, setlimit] = useState(10);
  const [commentsList, setCommentsList] = useState<IComment[]>([])
  const {lang} = useLanguageContext();
  const queryClient = useQueryClient();

  const {status: articleStatus, data: article, error: articleError} = useQuery({
    queryFn: async () => await getArticleById(params.id),
    queryKey: ["article", params.id],
    staleTime: Infinity
  });

  const {status: commentsStatus, data: comments, error: commentsError} = useQuery({
    queryFn: async () => {
      const response = await getComments(Number(params.id), true, offset, limit)
      setCommentsList([...commentsList, ...response.data]);
      setOffset(offset + limit);
      return {...response, data: [...commentsList, ...response.data]}
    },
    queryKey: ["comments", params.id],
    staleTime: Infinity
  });

  const mutation = useMutation({
    mutationFn: async (data: ICommentSchema) => {
      const nowDate = new Date();
      const resultDate = nowDate.toISOString();
      await createComment(Number(params.id), data.comment, data.author, resultDate);

    },
    onSuccess: () => {
      alert("Ваш отзыв отпрвлен администратору на одобрение и добавление перевода");
    }
  });

  const { register, handleSubmit, formState: { errors, isSubmitting }, reset: resetComment} = useForm<ICommentSchema>({resolver: zodResolver(commentSchema)});

  const onSubmit = (data: ICommentSchema) => {
    mutation.mutate(data);
    resetComment();
  }

  const handleButtonAddComments = () => {
    queryClient.invalidateQueries({queryKey: ["comments", params.id]});
  }

  if (articleError) {
    return <p>Произошла ошибка</p>
  }
  if (articleStatus == 'pending' || commentsStatus == 'pending') {
    return <p>Загрузка...</p>;
  }

  return (
    <>
      <div className='article'>
        <p className='article__title'>{article?.title[lang as keyof IRuEng]}</p>
        <p className='article__authors'>{article?.authors.map((x, i) => 
          <span className='article__author' key={i} title={`${lang == "Ru"? "Аффилиация": "Affiliation"}: ${x.affilation}, Email: ${x.email}`}>{x.fullname[lang as keyof IRuEng]}, </span>
        )}</p>
        <div className="article__buttons">
          <a className="article__button-link" href={`${serverUrl}/api/v1/files/download/${article.documentID}`} download={true}>
            <button className='article__pdf-download'><FormattedMessage id='article-article__pdf-download'/></button>
          </a>
          <a className="article__button-link" href='#'>
            <button className='article__pdf-download article__pdf-download_second' onClick={() => {
              navigator.clipboard.writeText(article?.reference[lang as keyof IRuEng] ?? '');
            }}><FormattedMessage id='article-article__pdf-download_second'/></button>
          </a>
        </div>
              
            
        <p className='article__subtitle'><FormattedMessage id='article-article__annotation'/></p>
        <p className='article__text2'>{article?.content[lang as keyof IRuEng]}</p>

        <p className='article__text2'>Дата поступления: {transformDate(article?.dateReceipt)}</p>
        <p className='article__text2'>Дата принятия: {transformDate(article?.dateAcceptance)}</p>
        <p className='article__text'>DOI: {article?.doi}</p>
        
        <video className='article__video' height="396" width="728" controls>
          <source src={`${serverUrl}/api/v1/files/download/${article.videoID}?tr=w-728,h-396`} />
          Your browser does not support the video tag...
        </video>
        


        <p className='article__subtitle'><FormattedMessage id='article-article__words'/></p>
        <p className='article__text'>{article?.keywords.map(item => item[lang as keyof IRuEng]).join(", ")}</p>
        
        {/* <p className='article__subtitle'><FormattedMessage id='article-article__authors'/></p>
        <div className="article__authors-list">
        {article?.authors.map((author, i) => <Author author={author} key={`author_${author.fullname.Ru}_${i}`}/>)}
        </div> */}
        

        <p className='article__subtitle'><FormattedMessage id='article-article__literature'/></p>
        <ol className='article__literature-list article__text'>
          {
            article?.literature.map((item, itemId) => <li className='article__literature-book' key={itemId}>{item}</li>)
          }
        </ol>

              
              
        <p className='article__title'><FormattedMessage id='article-article__comments-title'/></p>
        <div className="article__comments">
          {comments?.data?.map((comment, commentId) => <Comment key={commentId} comment={comment}/>)}
        </div>
        {comments && comments!.data.length < comments!.allCount && <button className='catalog__add-issues-button' onClick={() => handleButtonAddComments()}>Загрузить ещё</button>}

  
        <p className='article__subtitle'><FormattedMessage id='article-article__comments-subtitle1'/></p>
        <p className='article__comment-title'><FormattedMessage id='article-article__comments-subtitle2'/></p>
        
        <form onSubmit={handleSubmit(onSubmit)}>
          <input className='article__input' {...register("author")} type="text" placeholder='Ваше имя...'/>
          <textarea 
            {...register("comment")}
            className='article__comment-field' 
            placeholder='Введите текст...' 
          >
          </textarea>
          {errors.comment && <p>{`${errors.comment.message}`}</p>}
          <button className='article__comment-button' disabled={isSubmitting} type="submit"><FormattedMessage id='article-article__comment-button'/></button>
        </form>
        <div className="buttom"></div>
      </div>
    </>
  );
};

export default ArticlePageClient;