"use client"
import { useRouter, useSearchParams } from 'next/navigation';
import React, {useState} from 'react';
import { useQueryClient } from '@tanstack/react-query';
import { FormattedMessage } from 'react-intl';
import { useLanguageContext } from '@/i18n/languageContext';

const SearchInput = ({setUpdate, setIsAppdating, setIsSearching, onButtonBackToIssues}: 
  {setUpdate: React.Dispatch<React.SetStateAction<string>>, 
    setIsAppdating: React.Dispatch<React.SetStateAction<boolean>>, 
    setIsSearching: React.Dispatch<React.SetStateAction<boolean>>,
    onButtonBackToIssues?: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void
  }) => {
  const {lang} = useLanguageContext()
  const [searchQuery, setSearchQuery] = useState('');
  const queryClient = useQueryClient();

  const onSearch = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setUpdate(searchQuery);
    setIsAppdating(true);
    setIsSearching(true);
    
  }
  const onClickButton = async (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    event.preventDefault();
    setUpdate(searchQuery);
    setIsAppdating(true);
    setIsSearching(true);
  }

  return (
    <form className="search-form" onSubmit={onSearch}>
      <input 
        className='search-input' 
        type="text" 
        onChange={(event) => setSearchQuery(event.target.value)}
        placeholder={lang == "Ru"? 'Поиск...': 'Search...'}
      />
      <button onClick={onClickButton} className='search-button' type='button'><FormattedMessage id='catalog-catalog__item-button-search' /></button>
      {onButtonBackToIssues && <button className='search-button search-button-back' onClick={onButtonBackToIssues} type='button'><FormattedMessage id='catalog-catalog__item-button-search-back' /></button>}
    </form>
  );
};

export default SearchInput;