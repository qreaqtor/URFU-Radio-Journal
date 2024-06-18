import { useMutation } from '@tanstack/react-query';
import { deleteFile, uploadFile } from '../service/api/FileSrvice';
import { createPost } from '../service/api/IssueService';
import React, { useRef, useState } from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useQueryClient } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';
import { issueSchema, TIssueSchema } from '../service/types/schemes';


const CreateIssuePage = () => {
  const [file, setFile] = useState<FormData>();
  const [image, setImage] = useState<FormData>();
  const fileRef = useRef<HTMLInputElement | null>(null)
  const imageRef = useRef<HTMLInputElement | null>(null)
  const [fileName, setFileName] = useState('');
  const [imageName, setImageName] = useState('');
  const [fileId, setFileId] = useState('');
  const [imageId, setImageId] = useState('');

  const queryClient = useQueryClient();
  const navigate = useNavigate();

  const {register, handleSubmit, formState: {errors}} = useForm<TIssueSchema>({resolver: zodResolver(issueSchema)});

  const issueMutation = useMutation({
    mutationFn: async (data: TIssueSchema) => {
      let fileResponse = '';
      let imageResponse = '';

      if (fileId) {
        await deleteFile(fileId);
      }
      fileResponse = await uploadFile(file!);
      setFileId(fileResponse);

      if (imageId) {
        await deleteFile(imageId);
      }
      imageResponse = await uploadFile(image!);
      setImageId(imageResponse);
      
      const dateResult = `${data.date}T12:00:00Z`
      await createPost( Number(data.year), Number(data.number), Number(data.volume), dateResult, fileResponse, imageResponse);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey: ["issues"]})
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
      
      <button className='small-button' type="submit" disabled={issueMutation.status == 'pending'}>{issueMutation.status == 'pending' ? "Загрузка..." : "Создать"}</button>
    </form>
  );
};

export default CreateIssuePage;