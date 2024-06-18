import { Link } from 'react-router-dom';
import { IIssue } from '../service/types/typesNew';
import { serverUrl } from '../service/utils/serverPaths';
// import { serverUrl } from '../service/utils/serverPaths';

type IssueComponentProps = {
  item: IIssue
}

const IssuseComponent = ({item}: IssueComponentProps) => {
  return (
    <div className="catalog__item">
      <div className="catalog__item-image">
        <Link className="catalog__item-image-link" to={`/issues/${item["id"]}`}>
          <img src={`${serverUrl}/api/v1/files/download/${item.imageID}`} alt="" />
        </Link>
      </div>
      <p className='catalog__item-title'>Том {item.volume}, №{item.number}({item.year})</p>
      <Link to={`/issues/${item["id"]}`}>
        <button className='catalog__item-button'>Редактировать</button>
      </Link>
    </div>
  );
};

export default IssuseComponent;