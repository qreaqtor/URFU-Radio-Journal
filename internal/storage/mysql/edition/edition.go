package editionst

import (
	"database/sql"
	"urfu-radio-journal/internal/models"
)

type EditionStorage struct {
	db *sql.DB
}

func NewEditionStorage(db *sql.DB) *EditionStorage {
	return &EditionStorage{
		db: db,
	}
}

func (es *EditionStorage) InsertOne(edition *models.EditionCreate) (string, error) {
	var id string

	err := es.db.QueryRow("INSERT INTO issues (year, number, volume, cover_path, file_path, date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING issue_id",
		edition.Year, edition.Number, edition.Volume, edition.CoverPathId, edition.FilePathId, edition.Date).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (es *EditionStorage) GetAll() ([]*models.EditionRead, error) {
	rows, err := es.db.Query("SELECT issue_id, year, number, volume, cover_path, file_path, date FROM issues")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var editions []*models.EditionRead
	for rows.Next() {
		var edition models.EditionRead

		err := rows.Scan(&edition.Id, &edition.Year, &edition.Number, &edition.Volume, &edition.CoverPathId, &edition.FilePathId, &edition.Date)

		if err != nil {
			return nil, err
		}
		editions = append(editions, &edition)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return editions, nil
}

func (es *EditionStorage) FindOne(id string) (*models.EditionRead, error) {
	var edition models.EditionRead
	err := es.db.QueryRow("SELECT issue_id, year, number, volume, cover_path, file_path, date FROM issues WHERE issue_id = $1", id).
		Scan(&edition.Id, &edition.Year, &edition.Number, &edition.Volume, &edition.CoverPathId, &edition.FilePathId, &edition.Date)

	if err != nil {
		return nil, err
	}

	return &edition, nil
}

func (es *EditionStorage) UpdateOne(newEdition *models.EditionUpdate) error {
	_, err := es.db.Exec("UPDATE issues SET year = $1, number = $2, volume = $3, cover_path = $4, file_path = $5 WHERE issue_id = $6",
		newEdition.Year, newEdition.Number, newEdition.Volume, newEdition.CoverPathId, newEdition.FilePathId, newEdition.Id)
	if err != nil {
		return err
	}
	return nil
}

func (es *EditionStorage) Delete(editionIdStr string) error {
	_, err := es.db.Exec("DELETE FROM issues WHERE issue_id = $1", editionIdStr)
	if err != nil {
		return err
	}
	return nil
}
