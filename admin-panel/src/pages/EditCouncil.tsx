import { zodResolver } from '@hookform/resolvers/zod';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { useEffect, useRef, useState } from 'react';
import { useForm } from 'react-hook-form';
import { getCouncilById, updateCouncil } from '../service/api/CouncilService';
import { deleteFile, uploadFile } from '../service/api/FileSrvice';
import { useNavigate, useParams } from 'react-router-dom';
import { deleteCouncil } from '../service/api/CouncilService';
import { councilSchema, TCouncilSchema } from '../service/types/schemes';


const EditCouncilPage = () => {
  const {id} = useParams();
  const [image, setImage] = useState<FormData>();
  const imageRef = useRef<HTMLInputElement | null>(null)
  const [imageName, setImageName] = useState('');
  const [imageId, setImageId] = useState('');

  const queryClient = useQueryClient();
  const navigate = useNavigate();

  const {register, handleSubmit, formState: {errors}, setValue} = useForm<TCouncilSchema>({resolver: zodResolver(councilSchema)});
  
  const { data: council, status: statusCouncil} = useQuery({
    queryFn: async () => await getCouncilById(id!),
    queryKey: ["council", id],
    staleTime: Infinity 
  }); 

  useEffect(() => {
    if (statusCouncil == 'success') {
      setValue('nameRu', String(council?.name.Ru));
      setValue('nameEng', String(council?.name.Eng));
      setValue('email', String(council?.email));
      setValue('scopus', String(council?.scopus));
      setValue('descriptionRu', String(council?.description.Ru));
      setValue('descriptionEng', String(council?.description.Eng));
      setValue('contentRu', String(council?.content.Ru));
      setValue('contentEng', String(council?.content.Eng));
      setValue('rank', String(council?.rank));
      setValue('locationRu', String(council?.location.Ru));
      setValue('locationEng', String(council?.location.Eng));
      setValue('dateJoin', String(council?.dateJoin.split('T')[0]));

      setImageId(council?.imageID);
    }
  }, [council, setValue, statusCouncil]); 

  const councilMutation = useMutation({
    mutationFn: async (data: TCouncilSchema) => {
      const dateResult = `${data.dateJoin}T12:00:00Z`;
      let imageResponse = imageId;

      if (image) {
        imageResponse = await uploadFile(image!);
        setImageId(imageResponse);
        try {
          await deleteFile(council!.imageID);
        } catch { /* empty */ }
        try {
          await deleteFile(imageId);
        } catch { /* empty */ }
      }

      updateCouncil(
        Number(id),
        data.nameRu,
        data.nameEng,
        data.email,
        imageResponse,
        data.scopus,
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
      queryClient.invalidateQueries({queryKey: ["councils"]});
      queryClient.invalidateQueries({queryKey: ["council", id]});
      alert('Изменения сохранены');
    }
  });

  const deleteCouncilMutation = useMutation({
    mutationFn: async () => await deleteCouncil(id!),
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey: ["councils"]})
      navigate(`/reductors`);
    }
  });
  
  const onSubmitHandle = (data: TCouncilSchema) => {
    councilMutation.mutate(data);
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

      <div className="reductor__edit-text edit-margin-short-top">Scopus</div>
      <input {...register("scopus")} className="edit-input" type="text" placeholder='Введите scopus...'/>
      {errors.scopus && <p className='error-form-message'>{`${errors.scopus.message}`}</p>}

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
      
      <div className="create-issue__buttons-save">
        <button className='small-button' type="submit">Сохранить</button>
        <button className='small-button-delete' type="button" onClick={() => deleteCouncilMutation.mutate()}>Удалить редактора</button>
      </div>
    </form>
  );
};

export default EditCouncilPage;