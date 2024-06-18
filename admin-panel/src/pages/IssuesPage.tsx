import IssuseComponent from '../components/IssueComponent';
import { Link, useNavigate } from 'react-router-dom';
import { getIssues } from '../service/api/IssueService';
import { useQuery, useQueryClient } from '@tanstack/react-query';
import { useEffect, useState } from 'react';
import { IIssue } from '../service/types/typesNew';

const IssuesPage = () => {
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  const [offset, setOffset] = useState(0);
  const [limit] = useState(3);
  const [issuesList, setIssuesList] = useState<IIssue[]>([]);

  const {data: issues, isLoading, error } = useQuery({
    queryFn: async () => {
      const response = await getIssues(offset, limit)
      setIssuesList([...issuesList, ...response.data]);
      setOffset(offset + limit);
      return {...response, data: [...issuesList, ...response.data]};
    }, 
    queryKey: ["issues"],
    staleTime: Infinity,
  });

  

  useEffect(() => {
    if (!window.localStorage.getItem('token')) {
      navigate(`/files`);
    }
  }, [navigate]);

  useEffect(() => {
    if (issues) {
      setOffset(issues.data.length);
      setIssuesList([...issuesList, ...issues.data]);
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const handleButtonAddIssues = () => {
    queryClient.invalidateQueries({queryKey: ['issues']});
  }


  if (error) {
    return <>произошла ошибка</>
  }
  if (isLoading) {
    return <>загрузка...</>
  }

  return (
    <div className="catalog">
      <Link to={`/issues/create`}>
        <button className='catalog__create-issue-button'>Создать выпуск</button>
      </Link>
    
      <div className="catalog__container">
        {issues?.data.map((item, id) => {
          return <IssuseComponent key={id} item={item}/>
        })}
      </div>
      {issues && issues.data.length < issues.allCount && <button className='catalog__add-issues-button' onClick={() => handleButtonAddIssues()}>Загрузить ещё</button>}
    </div>
  )
};

export default IssuesPage;