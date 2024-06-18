import React, { useEffect, useRef, useState } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import ArticleComponent from '../components/ArticleComponent';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useMutation, useQuery } from '@tanstack/react-query';
import { deleteFile, uploadFile } from '../service/api/FileSrvice';
import { deleteIssue, updatePost } from '../service/api/IssueService';
import { getIssueById } from '../service/api/IssueService';
import { getArticles } from '../service/api/ArticleService';
import { useQueryClient } from '@tanstack/react-query';
import { issueSchema, TIssueSchema } from '../service/types/schemes';


const EditIssuePage = () => {
  const {id} = useParams();
  const [fileId, setFileId] = useState('');
  const [imageId, setImageId] = useState('');

  const [file, setFile] = useState<FormData>();
  const [image, setImage] = useState<FormData>();
  const fileRef = useRef<HTMLInputElement | null>(null)
  const imageRef = useRef<HTMLInputElement | null>(null)
  const [fileName, setFileName] = useState('');
  const [imageName, setImageName] = useState('');

  const queryClient = useQueryClient();

  const navigate = useNavigate();

  const {register, handleSubmit, formState: {errors}, setValue } = useForm<TIssueSchema>({resolver: zodResolver(issueSchema)});

  const { data: issue, status: issueStatus } = useQuery({
    queryFn: async () => await getIssueById(id!),
    queryKey: ["issue", id],
    staleTime: Infinity
  });  

  useEffect(() => {
    if (issueStatus == 'success') {
      setValue('year', String(issue?.year));
      setValue('number', String(issue?.number));
      setValue('volume', String(issue?.volume));
      setValue('date', String(issue?.date.split('T')[0]));
      setFileId(issue!.documentID);
      setImageId(issue!.imageID);
    }
    
  }, [issue, setValue, issueStatus])

  const { data: articles, isLoading: isLoadingArtcles } = useQuery({
    queryFn: async () => await getArticles(id!),
    queryKey: ["articles", id],
    staleTime: Infinity,
    enabled: issueStatus == 'success'
  });

  const issueMutation = useMutation({
    mutationFn: async (data: TIssueSchema) => {

      let fileResponse = fileId;
      let imageResponse = imageId;

      // будут ошибки при обновлении файла если его изначально не было, поправлю позже
      if (file) {
        fileResponse = await uploadFile(file);
        await deleteFile(issue!.documentID);
      }
      if (image) {
        imageResponse = await uploadFile(image);
        await deleteFile(issue!.imageID);
      }
      await updatePost(issue!.id, Number(data.year), Number(data.number), Number(data.volume), data.date, fileResponse, imageResponse);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey: ["issues"]});
      alert('Изменения сохранены');
    }
  });

  const issueDeleteMutation = useMutation({
    mutationFn: async () => {
      await deleteIssue(id!);
      queryClient.invalidateQueries({queryKey: ["issues"]});
    },
    onSuccess: () => {
      navigate(`/issues`);
    }
  });

  const onSubmitHandle = (data: TIssueSchema) => {
    issueMutation.mutate(data);
  }

  const onLoadFile = (e: React.ChangeEvent<HTMLInputElement>) => {
    const form = new FormData();
    form.append('file', e.target.files!["0"]);
    setFileName(e.target.files!["0"].name)
    setFile(form);
  }

  const onLoadImage = (e: React.ChangeEvent<HTMLInputElement>) => {
    const form = new FormData();
    form.append('file', e.target.files!["0"]);
    setImageName(e.target.files!["0"].name)
    setImage(form);
  }

  const handleButtonLoadFile = () => {
    fileRef.current?.click();
  };

  const handleButtonLoadImage = () => {
    imageRef.current?.click();
  };

  return (
    <>
      <form onSubmit={handleSubmit(onSubmitHandle)}>
        <div className="edit-text">Год выпуска</div>
        <input {...register("year")} className="edit-input" type="text" placeholder='Введите год...'/>
        {errors.year && <p className='error-form-message'>{`${errors.year.message}`}</p>}

        <div className="edit-text">Номер выпуска</div>
        <input {...register("number")} className="edit-input" type="text" placeholder='Введите номер...'/>
        {errors.number && <p className='error-form-message'>{`${errors.number.message}`}</p>}

        <div className="edit-text">Том выпуска</div>
        <input {...register("volume")} className="edit-input" type="text" placeholder='Введите том...'/>
        {errors.volume && <p className='error-form-message'>{`${errors.volume.message}`}</p>}

        <div className="edit-text">Дата</div>
        <input {...register("date")} className="edit-input" type="text" placeholder='Введите дату YYYY-MM-DD...'/>
        {errors.date && <p className='error-form-message'>{`${errors.date.message}`}</p>}

        <input className='hidden' ref={fileRef} type="file" name="file" accept=".pdf" onChange={onLoadFile}/>
        <input className='hidden' ref={imageRef} type="file" name="file" accept=".jpeg" onChange={onLoadImage}/>

        <div className="create-issue__buttons edit-margin-top">
          <button className='small-button' onClick={handleButtonLoadFile} type="button">{fileName ? fileName : 'Загрузить файл'}</button>
          <button className='small-button' onClick={handleButtonLoadImage} type="button">{imageName ? imageName : 'Загрузить обложку'}</button>
        </div>
        
        <div className="create-issue__buttons-save">
          <button className='small-button' type="submit">Сохранить</button>
          <button className='small-button-delete' type="button" onClick={() => issueDeleteMutation.mutate()}>Удалить выпуск</button>
        </div>
      </form>

      <p className='issue__title'>Содержание</p>

      {!isLoadingArtcles && articles?.map(article => <ArticleComponent key={article["id"]} article={article}/>)}
      <Link to={`/issues/article/create/${id}`}>
        <button className='catalog__create-issue-button'>Добавить статью</button>
      </Link> 
    </>
  );
};

export default EditIssuePage;