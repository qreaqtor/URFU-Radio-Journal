import { zodResolver } from '@hookform/resolvers/zod';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import { useRef, useState } from 'react';
import { useForm } from 'react-hook-form';
import { deleteFile, uploadFile } from '../service/api/FileSrvice';
import { createReductor } from '../service/api/ReductorService';
import { reductorSchema, TReductorSchema } from '../service/types/schemes';
import { useNavigate } from 'react-router-dom';



const CreateReductorPage = () => {
  const [image, setImage] = useState<FormData>();
  const imageRef = useRef<HTMLInputElement | null>(null)
  const [imageName, setImageName] = useState('');
  const [imageId, setImageId] = useState('');

  const queryClient = useQueryClient();
  const navigate = useNavigate();

  const {register, handleSubmit, formState: {errors}} = useForm<TReductorSchema>({resolver: zodResolver(reductorSchema)});
  
  const reductorMutation = useMutation({
    mutationFn: async (data: TReductorSchema) => {
      let imageResponse = '';
      if (imageId) {
        try {
          await deleteFile(imageId);
        } catch (error) { /* empty */ }
      }
      imageResponse = await uploadFile(image!);
      setImageId(imageResponse);
      const dateResult = `${data.dateJoin}T12:00:00Z`

      createReductor(
        data.nameRu,
        data.nameEng,
        data.email,
        imageResponse,
        data.descriptionRu,
        data.descriptionEng,
        data.contentRu,
        data.contentEng,
        data.rank,
        data.locationRu,
        data.locationEng,
        dateResult
      );
    },
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey: ["reductors"]})
      navigate(`/reductors`);
    }
  });
  
  const onSubmitHandle = (data: TReductorSchema) => {
    reductorMutation.mutate(data);
  }

  const onLoadImage = (e: React.ChangeEvent<HTMLInputElement>) => {
    const form = new FormData();
    form.append('file', e.target.files!["0"]);
    setImageName(e.target.files!["0"].name)
    setImage(form);
  }

  const handleButtonLoadImage = () => {
    imageRef.current?.click();
  };


  return (
    <form onSubmit={handleSubmit(onSubmitHandle)}>
      <div className="reductor__edit-text">ФИО редактора</div>
      <input {...register("nameRu")} className="edit-input-short-margin-start" type="text" placeholder='Введите ФИО редактора на русском...'/>
      {errors.nameRu && <p className='error-form-message'>{`${errors.nameRu.message}`}</p>}
      <input {...register("nameEng")} className="edit-input-short-margin" type="text" placeholder='Введите ФИО редактора на английском...'/>
      {errors.nameEng && <p className='error-form-message'>{`${errors.nameEng.message}`}</p>}

      <div className="reductor__edit-text edit-margin-top">Email</div>
      <input {...register("email")} className="edit-input" type="text" placeholder='Введите email...'/>
      {errors.email && <p className='error-form-message'>{`${errors.email.message}`}</p>}

      <div className="reductor__edit-text edit-margin-top">Описание</div>
      <textarea {...register("descriptionRu")} className="edit-input-short-margin-start edit-input-short-margin-start-textarea" placeholder='Введите описание на русском...'/>
      {errors.descriptionRu && <p className='error-form-message'>{`${errors.descriptionRu.message}`}</p>}
      <textarea {...register("descriptionEng")} className="edit-input-short-margin edit-input-short-margin-textarea" placeholder='Введите описание на английском...'/>
      {errors.descriptionEng && <p className='error-form-message'>{`${errors.descriptionEng.message}`}</p>}

      <div className="reductor__edit-text edit-margin-top">Контент</div>
      <input {...register("contentRu")} className="edit-input-short-margin-start" type="text" placeholder='Введите контент на русском...'/>
      {errors.contentRu && <p className='error-form-message'>{`${errors.contentRu.message}`}</p>}
      <input {...register("contentEng")} className="edit-input-short-margin" type="text" placeholder='Введите контент на английском...'/>
      {errors.contentEng && <p className='error-form-message'>{`${errors.contentEng.message}`}</p>}

      <div className="reductor__edit-text edit-margin-top">Должность</div>
      <input {...register("rank")} className="edit-input-short-margin" type="text" placeholder='Введите должность редактора...'/>
      {errors.rank && <p className='error-form-message'>{`${errors.rank.message}`}</p>}

      <div className="reductor__edit-text edit-margin-top">Местоположение</div>
      <input {...register("locationRu")} className="edit-input-short-margin-start" type="text" placeholder='Введите место работы редактора на русском...'/>
      {errors.locationRu && <p className='error-form-message'>{`${errors.locationRu.message}`}</p>}
      <input {...register("locationEng")} className="edit-input-short-margin" type="text" placeholder='Введите место работы редактора на английском...'/>
      {errors.locationEng && <p className='error-form-message'>{`${errors.locationEng.message}`}</p>}

      <div className="reductor__edit-text edit-margin-top">Дата вступления</div>
      <input {...register("dateJoin")} className="edit-input" type="text" placeholder='Введите дату вступления редактора YYYY-MM-DD...'/>
      {errors.dateJoin && <p className='error-form-message'>{`${errors.dateJoin.message}`}</p>}

      <input className='hidden' ref={imageRef} type="file" name="file" accept=".jpeg" onChange={onLoadImage}/>
      <button className='small-button edit-margin-top' onClick={handleButtonLoadImage} type="button">{imageName ? imageName : 'Загрузить изображение'}</button>
      
      <button className='small-button' type="submit">Создать</button>
    </form>
  );
};

export default CreateReductorPage;