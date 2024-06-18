import { zodResolver } from '@hookform/resolvers/zod';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { useEffect, useRef, useState } from 'react';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
// import { getCouncilById } from '../service/api/CouncilService';
import { deleteFile, uploadFile } from '../service/api/FileSrvice';
import { useNavigate, useParams } from 'react-router-dom';
import { deleteReductor, getReductorById, updateReductor } from '../service/api/ReductorService';


const reductorSchema = z.object({
  nameRu: z.string(),
  nameEng: z.string(),
  email: z.string(),
  descriptionRu: z.string(),
  descriptionEng: z.string(),
  contentRu: z.string(),
  contentEng: z.string(),
  rank: z.string(),
  locationRu: z.string(),
  locationEng: z.string(),
  dateJoin: z.string()
});

type TReductorSchema = z.infer<typeof reductorSchema>;

const EditCouncilPage = () => {
  const {id} = useParams();
  const [image, setImage] = useState<FormData>();
  const imageRef = useRef<HTMLInputElement | null>(null)
  const [imageName, setImageName] = useState('');
  const [imageId, setImageId] = useState('');

  const queryClient = useQueryClient();
  const navigate = useNavigate();

  const {register, handleSubmit, setValue} = useForm<TReductorSchema>({resolver: zodResolver(reductorSchema)});
  
  const { data: reductor, status: statusReductor} = useQuery({
    queryFn: async () => await getReductorById(id!),
    queryKey: ["reductor", id],
    staleTime: Infinity 
  }); 

  useEffect(() => {
    if (statusReductor == 'success') {
      setValue('nameRu', String(reductor?.name.Ru));
      setValue('nameEng', String(reductor?.name.Eng));
      setValue('email', String(reductor?.email));
      setValue('descriptionRu', String(reductor?.description.Ru));
      setValue('descriptionEng', String(reductor?.description.Eng));
      setValue('contentRu', String(reductor?.content.Ru));
      setValue('contentEng', String(reductor?.content.Eng));
      setValue('rank', String(reductor?.rank));
      setValue('locationRu', String(reductor?.location.Ru));
      setValue('locationEng', String(reductor?.location.Eng));
      setValue('dateJoin', String(reductor?.dateJoin.split('T')[0]));

      setImageId(reductor?.imageID);
    }
  }, [reductor, setValue, statusReductor]); 

  const reductorMutation = useMutation({
    mutationFn: async (data: TReductorSchema) => {
      const dateResult = `${data.dateJoin}T12:00:00Z`;
      let imageResponse = imageId;

      if (image) {
        console.log('file');
        imageResponse = await uploadFile(image!);
        try {
          await deleteFile(reductor!.imageID);
        } catch (error) {
          console.log(error);
        }
      }

      updateReductor(
        Number(id),
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
      alert('Изменения сохранены');
      queryClient.invalidateQueries({queryKey: ["reductors"]});
      queryClient.invalidateQueries({queryKey: ["reductor", id]});
    }
  });

  const deleteCouncilMutation = useMutation({
    mutationFn: async () => await deleteReductor(id!),
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey: ["reductors"]});
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
      <input {...register("nameRu")} className="edit-input-short-margin" type="text" placeholder='Введите ФИО редактора на русском...'/>
      <input {...register("nameEng")} className="edit-input" type="text" placeholder='Введите ФИО редактора на английском...'/>

      <div className="reductor__edit-text">Email</div>
      <input {...register("email")} className="edit-input" type="text" placeholder='Введите email...'/>

      <div className="reductor__edit-text edit-margin-top">Описание</div>
      <textarea {...register("descriptionRu")} className="edit-input-short-margin-start edit-input-short-margin-start-textarea" placeholder='Введите описание на русском...'/>
      {/* {errors.descriptionRu && <p className='error-form-message'>{`${errors.descriptionRu.message}`}</p>} */}
      <textarea {...register("descriptionEng")} className="edit-input-short-margin edit-input-short-margin-textarea" placeholder='Введите описание на английском...'/>
      {/* {errors.descriptionEng && <p className='error-form-message'>{`${errors.descriptionEng.message}`}</p>} */}

      <div className="reductor__edit-text">Контент</div>
      <input {...register("contentRu")} className="edit-input-short-margin" type="text" placeholder='Введите контент на русском...'/>
      <input {...register("contentEng")} className="edit-input" type="text" placeholder='Введите контент на английском...'/>

      <div className="reductor__edit-text">Должность</div>
      <input {...register("rank")} className="edit-input" type="text" placeholder='Введите должность редактора...'/>

      <div className="reductor__edit-text">Местоположение</div>
      <input {...register("locationRu")} className="edit-input-short-margin" type="text" placeholder='Введите место работы редактора на русском...'/>
      <input {...register("locationEng")} className="edit-input" type="text" placeholder='Введите место работы редактора на английском...'/>

      <div className="reductor__edit-text">Дата вступления</div>
      <input {...register("dateJoin")} className="edit-input" type="text" placeholder='Введите дату вступления редактора YYYY-MM-DD...'/>
      
      <input className='hidden' ref={imageRef} type="file" name="file" accept=".jpeg" onChange={onLoadImage}/>
      <button className='small-button' onClick={handleButtonLoadImage} type="button">{imageName ? imageName : 'Загрузить изображение'}</button>
      
      <div className="create-issue__buttons-save">
          <button className='small-button' type="submit">Сохранить</button>
          <button className='small-button-delete' type="button" onClick={() => deleteCouncilMutation.mutate()}>Удалить редактора</button>
        </div>
    </form>
  );
};

export default EditCouncilPage;