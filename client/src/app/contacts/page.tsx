"use client"
import React from 'react';
import Image from 'next/image';
import map from "../../../public/assets/photo_5269545055805365753_y.jpg";
import { FormattedMessage } from 'react-intl'
import { useQuery } from '@tanstack/react-query';
import { getReductors } from '@/data/ReductorService';
import Reductor from '@/components/Reductor';

const ContactsPage = () => {
  const {status: reductorsStatus, data: reductors} = useQuery({
    queryFn: async () => await getReductors(),
    queryKey: ["reductors"],
    staleTime: Infinity
  });
  
  return (
    <div className="contacts">
      <div className="contacts__mail-address contacts__title"><FormattedMessage id='contacts-mail-address' /></div>
      <div className="contacts__address contacts__text"><FormattedMessage id='contacts-address1' /></div>
      <div className="contacts__address contacts__text"><FormattedMessage id='contacts-address2' /></div>
      <div className="contacts__address contacts__address_last contacts__text"><FormattedMessage id='contacts-address3' /></div>
      <div className="contacts__map">
        <Image src={map} alt="" />
      </div>
      {/* <div className="contacts__editor contacts__title"><FormattedMessage id='contacts-editor' /></div>
      <div className="contacts__editor-information">
        <div className="contacts__editor-information-name contacts__text"><FormattedMessage id='contacts-editor-information-name' /></div>
        <div className="contacts__editor-information-phone contacts__text"><FormattedMessage id='contacts-editor-information-phone' /></div>
        <div className="contacts__editor-information-email contacts__text"><FormattedMessage id='contacts-editor-information-email' /></div>
      </div>

      <div className="contacts__support contacts__title"><FormattedMessage id='contacts-support' /></div>
      <div className="contacts__support-information">
        <div className="contacts__support-information-name contacts__text"><FormattedMessage id='contacts-support-information-name' /></div>
        <div className="contacts__support-information-email contacts__text"><FormattedMessage id='contacts-support-information-email' /></div>
      </div> */}

      <div className="editors__board">
          <div className="editors__board-title contacts__title"><p><FormattedMessage id='editors-title5' /></p></div>
          <div className="editors__board-container">
            {reductors?.map(editor => {
              return <Reductor key={editor["id"]} editor={editor}/>
            })}
          </div>
        </div>
    </div>
  )
};

export default ContactsPage;