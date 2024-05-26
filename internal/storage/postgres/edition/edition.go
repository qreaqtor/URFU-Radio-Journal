package editionst

import (
	"database/sql"
	"fmt"
	"urfu-radio-journal/internal/models"
)

type EditionStorage struct {
	db    *sql.DB
	table string
}

func NewEditionStorage(db *sql.DB, table string) *EditionStorage {
	return &EditionStorage{
		db:    db,
		table: table,
	}
}

func (es *EditionStorage) InsertOne(edition *models.EditionCreate) (string, error) {
	var id string

	query := fmt.Sprintf(
		"INSERT INTO %s (year, number, volume, cover_path, file_path, date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING edition_id",
		es.table,
	)
	row := es.db.QueryRow(
		query,
		edition.Year,
		edition.Number,
		edition.Volume,
		edition.CoverPathId,
		edition.FilePathId,
		edition.Date,
	)

	err := row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (es *EditionStorage) GetAll() ([]*models.EditionRead, error) {
	query := fmt.Sprintf(
		"SELECT edition_id, year, number, volume, cover_path, file_path, date FROM %s",
		es.table,
	)
	rows, err := es.db.Query(query)
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
	query := fmt.Sprintf(
		"SELECT edition_id, year, number, volume, cover_path, file_path, date FROM %s WHERE edition_id = $1",
		es.table,
	)
	err := es.db.QueryRow(query, id).
		Scan(&edition.Id, &edition.Year, &edition.Number, &edition.Volume, &edition.CoverPathId, &edition.FilePathId, &edition.Date)

	if err != nil {
		return nil, err
	}

	return &edition, nil
}

func (es *EditionStorage) UpdateOne(newEdition *models.EditionUpdate) error {
	query := fmt.Sprintf(
		"UPDATE %s SET year = $1, number = $2, volume = $3, cover_path = $4, file_path = $5 WHERE edition_id = $6",
		es.table,
	)
	_, err := es.db.Exec(
		query,
		newEdition.Year,
		newEdition.Number,
		newEdition.Volume,
		newEdition.CoverPathId,
		newEdition.FilePathId,
		newEdition.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (es *EditionStorage) Delete(editionIdStr string) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE edition_id = $1",
		es.table,
	)
	_, err := es.db.Exec(query, editionIdStr)
	if err != nil {
		return err
	}
	return nil
}
