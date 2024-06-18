"use client"
import React, { useMemo } from 'react';
import MainEditor from '@/components/MainEditor';
import BoardEditor from '@/components/BoardEditor';
import { FormattedMessage } from 'react-intl';
// import { getEditors } from '@/data/ReductorApi';
import { getCouncils } from '@/data/CouncilApi';
import { useQuery } from '@tanstack/react-query';
import { getReductors } from '@/data/ReductorService';
import Reductor from '../Reductor';


 


const EditorsPageClient = () => {

  const {status: councilsStatus, data: councils, error} = useQuery({
    queryFn: async () => await getCouncils(),
    queryKey: ["councils"],
    staleTime: Infinity
  });
  const {status: reductorsStatus, data: reductors} = useQuery({
    queryFn: async () => await getReductors(),
    queryKey: ["reductors"],
    staleTime: Infinity
  });

  const mainEditor = useMemo(() => councils?.find(item => item["rank"] == "Главный редактор"), [councils]);
  const subEditor = useMemo(() => councils?.find(item => item["rank"] == "Заместитель главного редактора"), [councils]);
  const menagerEditor = useMemo(() => councils?.find(item => item["rank"] == "Зав. редакцией"), [councils]);



  if (reductorsStatus == 'pending' && reductorsStatus == 'pending') {
    return <p>Загрузка...</p>
  }

  return (
    <>
      <div className="editors">
        <div className="editor__title"><p className='title'><FormattedMessage id='editors-title1' /></p></div>
        {mainEditor && <MainEditor editor={mainEditor!}/>}
        <div className="editor__title"><p className='title'><FormattedMessage id='editors-title2' /></p></div>
        {subEditor && <MainEditor editor={subEditor!}/>}
        <div className="editor__title"><p className='title'><FormattedMessage id='editors-title3' /></p></div>
        {menagerEditor && <MainEditor editor={menagerEditor!}/>}
        
        <div className="editors__board">
          <div className="editors__board-title title"><p><FormattedMessage id='editors-title4' /></p></div>
          <div className="editors__board-container">
            {councils?.filter((item) => item["rank"] == "Член редакционного совета").map(editor => {
              return <BoardEditor key={editor["id"]} editor={editor}/>
            })}
          </div>
        </div>

        {/* <div className="editors__board">
          <div className="editors__board-title title"><p><FormattedMessage id='editors-title5' /></p></div>
          <div className="editors__board-container">
            {reductors?.map(editor => {
              return <Reductor key={editor["id"]} editor={editor}/>
            })}
          </div>
        </div> */}
      </div>
    </>
  )
};

export default EditorsPageClient;