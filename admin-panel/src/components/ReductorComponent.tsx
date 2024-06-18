import { Link } from 'react-router-dom';
import { IReductor } from '../service/types/typesNew';
import { serverUrl } from '../service/utils/serverPaths';

type ReductorComponentProps = {
  reductor: IReductor,
}

const ReductorComponent = ({reductor}: ReductorComponentProps) => {
  return (
    <div className="editors__item">
      <Link to={`/reductors/${reductor.id}`}>
        <div className="editors__item-image">
          <img src={`${serverUrl}/api/v1/files/download/${reductor.imageID}`} alt="" />
        </div>
      </Link>
      <div className="editors__item-description">
        <div className="editors__item-name">{reductor.name.Ru}</div>
        <div className="editors__item-ediucation">{reductor.location.Ru}</div>
        <div className="editors__item-rank">{reductor.rank}</div>
      </div>
      <Link to={`/reductors/${reductor.id}`}>
        <button className="editors__item-button">Редактировать</button>
      </Link>
    </div>
  );
};

export default ReductorComponent;