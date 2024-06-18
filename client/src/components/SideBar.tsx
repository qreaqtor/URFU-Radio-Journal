"use client"
// import { useState } from 'react';
// import { Link } from 'react-router-dom';
// import IssuesService from '../api/IssuesService';
import Link from 'next/link';
// import FileService from '../api/FileService';
import { FormattedMessage } from 'react-intl'
import imageStock from "../assets/issues/tom7n3.jpeg";
import Image from 'next/image';
import { getIssues } from '@/data/IssueApi';

// import { useQuery } from 'react-query';
import { useQuery } from '@tanstack/react-query';
import { serverUrl } from '@/lib/utils';
import { IIssue } from '@/lib/typesNew';


const SideBar = () => {
  const {status: issueStatus, data: issue } = useQuery({
    queryFn: async () => {
      const issues = await getIssues(0, 200)
      const lastIssue = issues.data.reduce((accumulator: IIssue, currentValue: IIssue) => {
        const currentDate = new Date(currentValue.date);
        const latestDate = new Date(accumulator.date);
        return currentDate > latestDate ? currentValue : accumulator;
      });
      return lastIssue;
    },
    queryKey: ["issue"],
    staleTime: Infinity,
  });


  if (issueStatus !== "success") {
    return <></>;
  }

  return (
    <div className="sidebar">
      <div className="sidebar__image">
        <Link href={`/catalog/${issue?.id}`}>
          <Image className='sidebar__image' src={`${serverUrl}/api/v1/files/download/${issue.imageID}`} alt="" width={240} height={380}/>
        </Link>
      </div>
      <a href="https://www.akc.ru/itm/ural-radio-engineering-journal/" target='blank'>
        <button className="sidebar__button">
          <FormattedMessage id='sidebar__button' />
        </button>
      </a>
      
      <p className="sidebar__text"><FormattedMessage id='sidebar-sidebar__issn-online' /></p>
      <p className="sidebar__text"><FormattedMessage id='sidebar-sidebar__issn-print' /></p>
    </div>
  );
};

export default SideBar;