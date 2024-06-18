import { FormattedMessage } from 'react-intl'
import { useLanguageContext } from '../i18n/languageContext';
import issueImage from "../../public/assets/issues/tom7n3.jpeg";
// import type { IIssue, IRuEng } from '@/lib/types';
import Image from 'next/image';
import Link from 'next/link';
import { IIssue, IRuEng} from '@/lib/typesNew';
import { serverUrl } from '@/lib/utils';
import { transformDate } from '@/lib/utils';


type IssueComponentProps = {
  issue: IIssue;
}

const Issue = ({issue}: IssueComponentProps) => {
  const {lang} = useLanguageContext();
  const date = transformDate(issue.date);
  return (
    <div className="catalog__item">
      <div className="catalog__item-image">
        <Link className="catalog__item-image-link" href={`/catalog/${issue.id}`}>
          <Image width={200} height={200} src={`${serverUrl}/api/v1/files/download/${issue.imageID}`} alt="" />
        </Link>
      </div>
      {/* исправить дату */}
      <p className='catalog__item-date'>{date}</p> 
      <p className='catalog__item-title'>
        {
          lang == 'Ru' 
            ? <>Том {issue.volume}, №{issue.number}({issue.year})</>
            : <>Volume {issue.volume}, №{issue.number}({issue.year})</>
        }
      </p>
      <Link  href={`/catalog/${issue.id}`}>
        <button className='catalog__item-button'><FormattedMessage id='catalog-catalog__item-button' /></button>
      </Link>
    </div>
  );
};


export default Issue;