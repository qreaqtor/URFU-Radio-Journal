package editionst

import (
	"database/sql"
	"fmt"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/internal/storage/postgres/utils"
)

type EditionStorage struct {
	table string
	db    *sql.DB
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
		edition.ImageID,
		edition.DocumentID,
		edition.Date,
	)

	err := row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (es *EditionStorage) GetAll(args *models.BatchArgs) ([]*models.EditionRead, error) {
	query := fmt.Sprintf(
		"SELECT edition_id, year, number, volume, cover_path, file_path, date FROM %s",
		es.table,
	)

	queryBatch := utils.AddBatchToQuery(query, args)

	rows, err := es.db.Query(queryBatch)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var editions []*models.EditionRead
	for rows.Next() {
		var edition models.EditionRead

		err := rows.Scan(&edition.Id, &edition.Year, &edition.Number, &edition.Volume, &edition.ImageID, &edition.DocumentID, &edition.Date)

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

	row := es.db.QueryRow(query, id)

	err := row.Scan(&edition.Id, &edition.Year, &edition.Number, &edition.Volume, &edition.ImageID, &edition.DocumentID, &edition.Date)
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
		newEdition.ImageID,
		newEdition.DocumentID,
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

func (es *EditionStorage) GetCount() (int, error) {
	query := fmt.Sprintf(
		"SELECT COUNT(*) FROM %s",
		es.table,
	)

	row := es.db.QueryRow(query)

	count := 0
	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
