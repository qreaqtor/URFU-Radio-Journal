import { useMutation, useQueryClient } from '@tanstack/react-query';
import { createArticle } from '../service/api/ArticleService';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useRef, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { IAuthor, IRuEng } from '../service/types/typesNew';
import { deleteFile, uploadFile } from '../service/api/FileSrvice';
import { createArticleSchema, TCreateArticleSchema } from '../service/types/schemes';


const CreateArticlePage = () => {
  const {id} = useParams();
  const [literature, setLitrature] = useState(['']);

  const [authorRu, setAuthorRu] = useState(['']);
  const [authorEng, setAuthorEng] = useState(['']);
  const [authorAffilation, setAuthorAffilation] = useState(['']);
  const [authorEmail, setAuthorEmail] = useState(['']);

  const [fileId, setFileId] = useState('');
  const [videoId, setVideoId] = useState('');
  const [file, setFile] = useState<FormData>();
  const [video, setVideo] = useState<FormData>();
  const fileRef = useRef<HTMLInputElement | null>(null)
  const videoRef = useRef<HTMLInputElement | null>(null)
  const [fileName, setFileName] = useState('');
  const [videoName, setVideoName] = useState('');

  const queryClient = useQueryClient();

  const navigate = useNavigate();

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


  const {register, handleSubmit, formState: {errors}} = useForm<TCreateArticleSchema>({resolver: zodResolver(createArticleSchema)});

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
      

      let fileResponse = '';
      let videoResponse = '';

      if (fileId) {
        await deleteFile(fileId);
      }
      fileResponse = await uploadFile(file!);
      setFileId(fileResponse);

      if (videoId) {
        await deleteFile(videoId);
      }
      videoResponse = await uploadFile(video!);
      setVideoId(videoResponse);

      const dateReceipt = `${data.dateReceipt}T00:00:00Z`;
      const dateAcceptance = `${data.dateAcceptance}T00:00:00Z`;

      await createArticle(
        Number(id), 
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
      queryClient.invalidateQueries({queryKey: ["articles", id]});
      navigate(`/issues/${id}`);
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
        <button className='small-button' type="submit" disabled={articleMutation.status == 'pending'}>{articleMutation.status == 'pending' ? "Загрузка..." : "Создать"}</button>
      </form>
      
    </div>
  );
};

export default CreateArticlePage;