import { useQuery } from '@tanstack/react-query';
import { getCouncils } from '../service/api/CouncilService';
import CouncilComponent from '../components/CouncilComponent';
import { Link, useNavigate } from 'react-router-dom';
import { getReductors } from '../service/api/ReductorService';
import ReductorComponent from '../components/ReductorComponent';
import { useEffect } from 'react';



const CouncilsPage = () => {
  const navigate = useNavigate();

  const { data: councils, isLoading, error } = useQuery({
    queryFn: async () => await getCouncils(),
    queryKey: ["councils"],
    staleTime: Infinity,
  });

  const { data: reductors, isLoading: isLoadingReductors, error: reductorsError } = useQuery({
    queryFn: async () => await getReductors(),
    queryKey: ["reductors"],
    staleTime: Infinity,
  });

  useEffect(() => {
    if (window.localStorage.getItem('token') == '') {
      navigate(`/files`);
    }
  }, [navigate]);

  if (error || reductorsError) {
    return <>{"произошла ошибка :("}</>
  }
  if (isLoading || isLoadingReductors) {
    return <>загрузка...</>
  }

  return (
    <>
      <p className='editors__title'>Редакционный совет</p>
      <div className="editors">
        <div className="editors__list">
         {councils?.map((council) => {
           return <CouncilComponent key={council.id} council={council}/>
         })}
       </div>
      </div>
      <Link to={`/councils/create`}>
        <button className="editors__button editors__button-margin">Добавить редактора</button>
      </Link>
      <p className='editors__title'>Редакция</p>
      <div className="editors">
        <div className="editors__list">
         {reductors?.map((reductor) => {
           return <ReductorComponent key={reductor.id} reductor={reductor}/>
         })}
       </div>
      </div>
      <Link to={`/reductors/create`}>
        <button className="editors__button">Добавить редактора</button>
      </Link>
    </>
  );
};

export default CouncilsPage;