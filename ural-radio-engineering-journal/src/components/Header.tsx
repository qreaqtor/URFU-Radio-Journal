"use client"
import React from 'react';
import image from "../../public/assets/image 4.svg"
import image2 from "../../public/assets/image 3.svg"
import Link from 'next/link';
import Image from 'next/image';
import { usePathname  } from 'next/navigation';
import ruIcon from "../../public/assets/ru_rus_russia_flag_icon_255814.png";
import engIcon from "../../public/assets/gb_flag_great_britain_england_union_jack_english_icon_228674.png";

import { useLanguageContext } from '../i18n/languageContext';
import { LOCALES } from '../i18n/locales'
import { FormattedMessage } from 'react-intl'


const Header = () => {
  
  const {localeChange} = useLanguageContext();
  const {lang} = useLanguageContext();
  const pathname = usePathname();
  const checkActiveLinc = (path: string) => pathname == path ? 'header__group-link header__group-link-active' : 'header__group-link';

  return (
    <header className='header'>
         <nav className='header__nav'>
          <div className="header__nav-group">
            <Link className={checkActiveLinc("/catalog")}  href="/catalog"><FormattedMessage id='header-navlink1' /></Link>
            <Link className={checkActiveLinc("/editors")} href="/editors"><FormattedMessage id='header-navlink2' /></Link>
            <Link className={checkActiveLinc("/")} href="/"><FormattedMessage id='header-navlink3' /></Link>
          </div>
          <div className="header__nav-group">
            <Link className="header__group-link header__group-link_logo" href="/">URAL RADIO ENGINEERING JOURNAL</Link>
          </div>
          <div className="header__nav-group">
            <Link className={checkActiveLinc("/contacts")} href="/contacts"><FormattedMessage id='header-navlink4' /></Link>
            <Link className={checkActiveLinc("/authors")} href="/authors"><FormattedMessage id='header-navlink5' /></Link>
          </div>
        </nav>

        <div className='header__languages'>
          <div className={`header__language-img ${lang == "Eng" ? "" : "header__language-img-none-active"}`} onClick={() => localeChange(LOCALES.ENGLISH)}>
            <Image src={engIcon} width="30" height="30" alt="" />
          </div>
          <div className={`header__language-img ${lang == "Ru" ? "" : "header__language-img-none-active"}`} onClick={() => localeChange(LOCALES.RUSSIAN)}>
            <Image src={ruIcon}  width="30" height="30" alt="" />
          </div>
        </div>
      </header>
  );
};

export default Header;