import { useMutation } from '@tanstack/react-query';
import { login } from '../service/api/AuthService';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';

const loginSchema = z.object({
  username: z.string(),
  password: z.string()
});

type TLoginSchema = z.infer<typeof loginSchema>;

const FilesPage = () => {
  const navigate = useNavigate();
  
  const {register, handleSubmit} = useForm<TLoginSchema>({resolver: zodResolver(loginSchema)});
  
  const loginMutation = useMutation({
    mutationFn: async (data: TLoginSchema) => {
      await login(data.username, data.password);
    },
    onSuccess: () => {
      navigate(`/issues`);
    }
  })

  const onSubmitHandle = (data: TLoginSchema) => {
    loginMutation.mutate(data);
  }
  const onExitHandle = () => {
    window.localStorage.setItem('token', '');
    alert('Вы успешно вышли');
  }

  return (
    <div className='login'>
      <form onSubmit={handleSubmit(onSubmitHandle)}>
        <input {...register("username")} className='login__login-input' type="text" placeholder='login...'/>
        <input {...register("password")} className='login__password-input' type="password" placeholder='password...'/>

        <button className='login__button' type='submit'>войти</button>
        <button className='login__exit-button' type='button' onClick={onExitHandle}>выйти</button>
      </form>
      
    </div>
  );
};

export default FilesPage;