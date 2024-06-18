import { Link } from 'react-router-dom';
import { ICouncil } from '../service/types/typesNew';
import { serverUrl } from '../service/utils/serverPaths';

type CouncilComponentProps = {
  council: ICouncil,
}

const CouncilComponent = ({council}: CouncilComponentProps) => {
  return (
    <div className="editors__item">
      <Link to={`/councils/${council.id}`}>
        <div className="editors__item-image">
          <img src={`${serverUrl}/api/v1/files/download/${council.imageID}`} alt="" />
        </div>
      </Link>
      <div className="editors__item-description">
        <div className="editors__item-name">{council.name.Ru}</div>
        <div className="editors__item-ediucation">{council.location.Ru}</div>
        <div className="editors__item-rank">{council.rank}</div>
      </div>
      <Link to={`/councils/${council.id}`}>
        <button className="editors__item-button">Редактировать</button>
      </Link>
    </div>
  );
};

export default CouncilComponent;