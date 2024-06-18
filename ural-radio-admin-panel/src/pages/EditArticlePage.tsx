import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { updateArticle, getArticleById, deleteArticle } from '../service/api/ArticleService';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useEffect, useRef, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { IAuthor, IRuEng } from '../service/types/typesNew';
import { deleteFile, uploadFile } from '../service/api/FileSrvice';
import { createArticleSchema, TCreateArticleSchema } from '../service/types/schemes';
import { getComments } from '../service/api/CommentService';
// import { getComments } from '../service/api/CommentService';
// import { createComment } from '../service/api/CommentService';
import CommentComponent from '../components/CommentComponent';


const EditArticlePage = () => {
  const {id} = useParams();
  const [literature, setLitrature] = useState<string[]>([]);

  const [authorRu, setAuthorRu] = useState<string[]>([]);
  const [authorEng, setAuthorEng] = useState<string[]>([]);
  const [authorAffilation, setAuthorAffilation] = useState<string[]>([]);
  const [authorEmail, setAuthorEmail] = useState<string[]>([]);

  const [file, setFile] = useState<FormData>(); 
  const [video, setVideo] = useState<FormData>();
  const [fileId, setFileId] = useState('');
  const [videoId, setVideoId] = useState('');
  const fileRef = useRef<HTMLInputElement | null>(null)
  const videoRef = useRef<HTMLInputElement | null>(null)
  const [fileName, setFileName] = useState('');
  const [videoName, setVideoName] = useState('');

  const queryClient = useQueryClient();

  const navigate = useNavigate();

  const {register, handleSubmit, setValue, formState: {errors}} = useForm<TCreateArticleSchema>({resolver: zodResolver(createArticleSchema)});

  const { data: article, status: statusArticle} = useQuery({
    queryFn: async () => {
      const response = await getArticleById(id!)
      return response;
    },
    queryKey: ["article", id],
    staleTime: Infinity 
  }); 

  const { data: comments, status: statusComments} = useQuery({
    queryFn: async () => await getComments(Number(id!), false),
    queryKey: ["comments", id],
    staleTime: Infinity 
  }); 

  const { data: commentsApproved, status: statusCommentsApproved} = useQuery({
    queryFn: async () => await getComments(Number(id!), true),
    queryKey: ["commentsApproved", id],
    staleTime: Infinity 
  }); 

  const articleDeleteMutation = useMutation({
    mutationFn: async () => {
      await deleteArticle(id!);
      
    }, 
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey: ["articles", String(article?.editionId)]});
      navigate(`/issues/${String(article?.editionId)}`);
    }
  });

  useEffect(() => {
    if (statusArticle == 'success') {
      setValue('titleRu', String(article?.title.Ru));
      setValue('titleEng', String(article?.title.Eng));
      setValue('contentRu', String(article?.content.Ru));
      setValue('contentEng', String(article?.content.Eng));
      setValue('keywordsRu', String(article?.keywords.map(x => x.Ru).join(', ')));
      setValue('keywordsEng', String(article?.keywords.map(x => x.Eng).join(', ')));
      setValue('referenceRu', String(article?.reference.Ru));
      setValue('referenceEng', String(article?.reference.Eng));
      setValue('dateReceipt', String(article?.dateReceipt.split('T')[0]));
      setValue('dateAcceptance', String(article?.dateAcceptance.split('T')[0]));
      setValue('doi', String(article?.doi));

      setFileId(article?.documentID);
      
      setVideoId(article?.videoID);
      if (article?.literature == null) {
        setLitrature([]);
      } else {
        setLitrature(article?.literature);
      }
      
      setAuthorRu([]);
      setAuthorEng([]);
      setAuthorAffilation([]);
      setAuthorEmail([]);    

      if (article.authors == null) {
        setAuthorRu([]);
        setAuthorEng([]);
        setAuthorAffilation([]);
        setAuthorEmail([]);
      } else {
        for (let i = 0; i < article!.authors.length; i++) {
          setAuthorRu(prevAuthorRu => {
            return [...prevAuthorRu, article!.authors[i].fullname.Ru]
          });
          setAuthorEng(prevAuthorEng => {
            return [...prevAuthorEng, article!.authors[i].fullname.Eng]
          });
          setAuthorAffilation(prevAuthorAffilation => {
            return [...prevAuthorAffilation, article!.authors[i].affilation]
          });
          setAuthorEmail(prevAuthorEmail => {
            return [...prevAuthorEmail, article!.authors[i].email]
          });
        }
      }      
    }
  }, [article, setValue, statusArticle]); 

  const articleMutation = useMutation({
    mutationFn: async (data: TCreateArticleSchema) => {
      const authors: IAuthor[] = authorRu.map((_x, i) => ({
        fullname: {
          Ru: authorRu[i],
          Eng: authorEng[i]
        },
        affilation: authorAffilation[i],
        email: authorEmail[i]
      }));

      const keywordsRu = data.keywordsRu.split(', ');
      const keywordsEng = data.keywordsEng.split(', ');
      const keyWords: IRuEng[] = data.keywordsRu.split(', ').map((_x, i) => ({
        Ru: keywordsRu[i],
        Eng: keywordsEng[i],
      }))

      console.log(keyWords);
      
      let fileResponse = fileId;
      let videoResponse = videoId;

      if (file) {
        console.log('file');
        fileResponse = await uploadFile(file);
        await deleteFile(article!.documentID);
      }
      if (video) {
        videoResponse = await uploadFile(video);
        await deleteFile(article!.videoID);
      }
      const dateReceipt = `${data.dateReceipt}T00:00:00Z`;
      const dateAcceptance = `${data.dateAcceptance}T00:00:00Z`;


      await updateArticle(
        article!.id,
        data.titleRu,
        data.titleEng,
        authors,
        data.contentRu,
        data.contentEng,
        keyWords,
        fileResponse,
        literature,
        data.referenceRu,
        data.referenceEng,
        videoResponse,
        dateReceipt,
        dateAcceptance,
        data.doi
      );
    },
    onSuccess: () => {
      alert('Изменения сохранены');
      queryClient.invalidateQueries({queryKey: ["articles", String(article?.editionId)]});
      queryClient.invalidateQueries({queryKey: ["article", id]});
    }
  })

  const onSubmitHandle = (data: TCreateArticleSchema) => {
    articleMutation.mutate(data);
  }

  const onChangeLiterature = (i: number, e: React.ChangeEvent<HTMLInputElement>) => {
    setLitrature(prevLiterateure => {
      const newLiterature = [...prevLiterateure];
      newLiterature[i] = e.target.value;
      return newLiterature;
    })
  }

  const onAddLiterature = () => {
    setLitrature([...literature, '']);
  }

  const onDeleteLiterature = () => {
    setLitrature(literature.slice(0, literature.length - 1));
  }

  const onChangeAuthorRu = (i: number, e: React.ChangeEvent<HTMLInputElement>) => {
    setAuthorRu(prevAuthorRu => {
      const newAuthorRu = [...prevAuthorRu];
      newAuthorRu[i] = e.target.value;
      return newAuthorRu;
    })
  }

  const onChangeAuthorEng = (i: number, e: React.ChangeEvent<HTMLInputElement>) => {
    setAuthorEng(prevAuthorEng => {
      const newAuthorEng = [...prevAuthorEng];
      newAuthorEng[i] = e.target.value;
      return newAuthorEng;
    })
  }

  const onChangeAuthorAffilation = (i: number, e: React.ChangeEvent<HTMLInputElement>) => {
    setAuthorAffilation(prevAuthorAffilation => {
      const newAuthorAffilation = [...prevAuthorAffilation];
      newAuthorAffilation[i] = e.target.value;
      return newAuthorAffilation
    })
  }

  const onChangeAuthorEmail = (i: number, e: React.ChangeEvent<HTMLInputElement>) => {
    setAuthorEmail(prevAuthorEmail => {
      const newAuthorEmail = [...prevAuthorEmail];
      newAuthorEmail[i] = e.target.value;
      return newAuthorEmail
    })
  }

  const onAddAuthor = () => {
    setAuthorRu([...authorRu, '']);
    setAuthorEng([...authorEng, '']);
    setAuthorAffilation([...authorAffilation, '']);
    setAuthorEmail([...authorEmail, '']);
  }

  const onDeleteAuthor = () => {
    setAuthorRu(authorRu.slice(0, authorRu.length - 1));
    setAuthorEng(authorEng.slice(0, authorEng.length - 1));
    setAuthorAffilation(authorAffilation.slice(0, authorAffilation.length - 1));
    setAuthorEmail(authorEmail.slice(0, authorEmail.length - 1));
  }

  const onLoadFile = (e: React.ChangeEvent<HTMLInputElement>) => {
    const form = new FormData();
    form.append('file', e.target.files!["0"]);
    setFileName(e.target.files!["0"].name)
    setFile(form);
  }

  const onLoadVideo = (e: React.ChangeEvent<HTMLInputElement>) => {
    const form = new FormData();
    form.append('file', e.target.files!["0"]);
    setVideoName(e.target.files!["0"].name)
    setVideo(form);
  }

  const handleButtonLoadFile = () => {
    fileRef.current?.click();
  };

  const handleButtonLoadVideo = () => {
    videoRef.current?.click();
  };


  return (
    <div>
      <form onSubmit={handleSubmit(onSubmitHandle)}>
        <div className="edit-text">Заголовок</div>
        <input {...register("titleRu")} className="edit-input-short-margin-start" type="text" placeholder='Введите заголовок на русском...'/>
        {errors.titleRu && <p className='error-form-message'>{`${errors.titleRu.message}`}</p>}
        <input {...register("titleEng")} className="edit-input-short-margin" type="text" placeholder='Введите заголовок на английском...'/>
        {errors.titleEng && <p className='error-form-message'>{`${errors.titleEng.message}`}</p>}

        <div className="edit-text edit-margin-top">Авторы</div>
        {authorRu.map((_x, i) => 
          <div key={`${i}_author`}>
            <input  
              value={authorRu[i]}
              onChange={(e) => onChangeAuthorRu(i, e)}
              className={i == 0 ? "edit-input-short-margin-start" : "edit-input-short-margin edit-margin-top"} 
              type="text" 
              placeholder='Введите имя автора на русском...'
            />
            <input 
              value={authorEng[i]}
              onChange={(e) => onChangeAuthorEng(i, e)}
              className="edit-input-short-margin" 
              type="text" 
              placeholder='Введите имя автора на английском...'
            />
            <input 
              value={authorAffilation[i]}
              onChange={(e) => onChangeAuthorAffilation(i, e)}
              className="edit-input-short-margin" 
              type="text" 
              placeholder='Введите аффиляцию...'
            />
            <input 
              value={authorEmail[i]}
              onChange={(e) => onChangeAuthorEmail(i, e)}
              className="edit-input-short-margin" 
              type="text" 
              placeholder='Введите email...'
            />
          </div>
        )}
        <div className="article__literature-buttons edit-margin-short-top">
          <button className='small-button' type="button" onClick={onAddAuthor}>Добавить автора</button>
          <button className='small-button-delete' type="button" onClick={onDeleteAuthor}>Удалить автора</button>
        </div>
        
        <div className="edit-text">Аннотация</div>
        <textarea {...register("contentRu")} className="edit-input-short-margin-start edit-input-short-margin-start-textarea" placeholder='Введите аннотацию на русском...'/>
        {errors.contentRu && <p className='error-form-message'>{`${errors.contentRu.message}`}</p>}
        <textarea {...register("contentEng")} className="edit-input-short-margin edit-input-short-margin-textarea" placeholder='Введите аннотацию на английском...'/>
        {errors.contentEng && <p className='error-form-message'>{`${errors.contentEng.message}`}</p>}

        <div className="edit-text edit-margin-top">Ключевые слова</div>
        <textarea {...register("keywordsRu")} className="edit-input-short-margin-start edit-input-short-margin-start-textarea" placeholder='Ключевые слова на русском через запятую...'/>
        <textarea {...register("keywordsEng")} className="edit-input-short-margin edit-input-short-margin-textarea" placeholder='Ключевые слова на английском через запятую...'/>

        <div className="edit-text edit-margin-top">Литература</div>
        {literature.map((_x, i) => 
          <input 
            value={literature[i]}
            onChange={(e) => onChangeLiterature(i, e)}
            className={i == 0 ? "edit-input-short-margin-start" : "edit-input-short-margin"} 
            type="text" 
            placeholder='Литература...'
            key={`${i}_literature`}
          />
        )}
        <div className="article__literature-buttons edit-margin-short-top">
          <button className='small-button' type="button" onClick={onAddLiterature}>Добавить литературу</button>
          <button className='small-button-delete' type="button" onClick={onDeleteLiterature}>Удалить литературу</button>
        </div>
        

        <div className="edit-text">Ссылка</div>
        <input {...register("referenceRu")} className="edit-input-short-margin-start" type="text" placeholder='Ссылка на русском...'/>
        {errors.referenceRu && <p className='error-form-message'>{`${errors.referenceRu.message}`}</p>}
        <input {...register("referenceEng")} className="edit-input-short-margin" type="text" placeholder='Ссылка на английском...'/>
        {errors.referenceEng && <p className='error-form-message'>{`${errors.referenceEng.message}`}</p>}

        <div className="edit-text edit-margin-top">Даты</div>
        <input {...register("dateReceipt")} className="edit-input-short-margin-start" type="text" placeholder='Дата получения YYYY-MM-DD...'/>
        {errors.dateReceipt && <p className='error-form-message'>{`${errors.dateReceipt.message}`}</p>}
        <input {...register("dateAcceptance")} className="edit-input-short-margin" type="text" placeholder='Дата принятия YYYY-MM-DD...'/>
        {errors.dateAcceptance && <p className='error-form-message'>{`${errors.dateAcceptance.message}`}</p>}

        <div className="edit-text edit-margin-top">Doi</div>
        <input {...register("doi")} className="edit-input-short-margin-start" type="text" placeholder='Введите Doi...'/>

        {/* <input type="file" accept=".pdf" onChange={onLoadFile}/>
        <input type="file" onChange={onLoadVideo}/>

        
        <button className='catalog__create-issue-button' type='submit'>создать</button> */}
      
        <input className='hidden' ref={fileRef} type="file" name="file" accept=".pdf" onChange={onLoadFile}/>
        <input className='hidden' ref={videoRef} type="file" name="file" accept=".jpeg" onChange={onLoadVideo}/>

        <div className="create-issue__buttons edit-margin-short-top">
          <button className='small-button' onClick={handleButtonLoadFile} type="button">{fileName ? fileName : 'Загрузить файл'}</button>
          <button className='small-button' onClick={handleButtonLoadVideo} type="button">{videoName ? videoName : 'Загрузить Видео'}</button>
        </div>

        {/* <button className='catalog__create-issue-button' type='submit'>создать</button> */}
        <div className="create-issue__buttons-save">
          <button className='small-button' type="submit" disabled={articleMutation.status == 'pending'}>{articleMutation.status == 'pending' ? "Загрузка..." : "Сохранить"}</button>
          <button className='small-button-delete' type="button" onClick={() => articleDeleteMutation.mutate()}>Удалить статью</button>
        </div>
      </form>
      
      {/* <button className='small-button' type="button" onClick={() => createCommentMutation.mutate()}>Коммент</button> */}
      <div className="article-page__comments">
        <p className='article-page__comments-title'>Комментарии</p>
        {statusComments == 'success' && comments?.map((x, i) => <CommentComponent comment={x} key={`${x.author}_${x.id}_${i}`}/>)}
        <p className='article-page__comments-title'>Одобренные комментарии</p>
        {statusCommentsApproved == 'success' &&commentsApproved?.map((x, i) => <CommentComponent comment={x} key={`${x.author}_${x.id}_${i}`}/>)}
      </div>
    </div>
  );
};

export default EditArticlePage;