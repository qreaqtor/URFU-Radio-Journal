package councilst

import (
	"database/sql"
	"fmt"
	"urfu-radio-journal/internal/models"
)

type CouncilStorage struct {
	db    *sql.DB
	table string
}

func NewCouncilStorage(db *sql.DB, table string) *CouncilStorage {
	return &CouncilStorage{
		db:    db,
		table: table,
	}
}

func getColumns() string {
	return "fullname_ru, fullname_en, description_ru, description_en, location_ru, location_en, email, scopus, photo_path, date_join, rank, content_ru, content_en"
}

func generateValuesID(count int) string {
	res := ""
	for i := 1; i < count; i++ {
		res += fmt.Sprintf("$%v, ", i)
	}
	res += fmt.Sprintf("$%v", count)
	return res
}

func (cs *CouncilStorage) InsertOne(member *models.CouncilMemberCreate) (string, error) {
	var id string

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING member_id",
		cs.table,
		getColumns(),
		generateValuesID(13),
	)
	row := cs.db.QueryRow(
		query,
		member.Name.Ru,
		member.Name.Eng,
		member.Description.Ru,
		member.Description.Eng,
		member.Location.Ru,
		member.Location.Eng,
		member.Email,
		member.ScopusURL,
		member.ImageID,
		member.DateJoin,
		member.Rank,
		member.Content.Ru,
		member.Content.Eng,
	)

	err := row.Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (cs *CouncilStorage) UpdateOne(memberIdStr string, memberUpdate *models.CouncilMemberUpdate) error {
	query := fmt.Sprintf(
		`UPDATE %s 
		SET 
		fullname_ru = $1, 
		fullname_en = $2, 
		description_ru = $3, 
		description_en = $4, 
		location_ru = $5, 
		location_en = $6, 
		email = $7, 
		scopus = $8, 
		photo_path = $9, 
		date_join = $10, 
		rank = $11, 
		content_ru = $12, 
		content_en = $13 
		WHERE member_id = $14`,
		cs.table,
	)

	_, err := cs.db.Exec(
		query,
		memberUpdate.Name.Ru,
		memberUpdate.Name.Eng,
		memberUpdate.Description.Ru,
		memberUpdate.Description.Eng,
		memberUpdate.Location.Ru,
		memberUpdate.Location.Eng,
		memberUpdate.Email,
		memberUpdate.ScopusURL,
		memberUpdate.ImageID,
		memberUpdate.DateJoin,
		memberUpdate.Rank,
		memberUpdate.Content.Ru,
		memberUpdate.Content.Eng,
		memberIdStr,
	)

	if err != nil {
		return err
	}

	return nil
}

func (cs *CouncilStorage) Delete(idStr string) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE member_id = $1",
		cs.table,
	)
	_, err := cs.db.Exec(query, idStr)

	if err != nil {
		return err
	}
	return nil
}

func (cs *CouncilStorage) GetImageID(idStr string) (string, error) {
	var imageID string

	query := fmt.Sprintf(
		"SELECT photo_path FROM %s WHERE member_id = $1",
		cs.table,
	)

	row := cs.db.QueryRow(query, idStr)
	err := row.Scan(&imageID)

	if err != nil {
		return "", err
	}

	return imageID, nil
}

func (cs *CouncilStorage) GetAll() ([]*models.CouncilMemberRead, error) {
	var members []*models.CouncilMemberRead
	query := fmt.Sprintf(
		"SELECT member_id, %s FROM %s",
		getColumns(),
		cs.table,
	)

	rows, err := cs.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var member models.CouncilMemberRead

		err = rows.Scan(
			&member.Id,
			&member.Name.Ru,
			&member.Name.Ru,
			&member.Description.Ru,
			&member.Description.Eng,
			&member.Location.Ru,
			&member.Location.Eng,
			&member.Email,
			&member.ScopusURL,
			&member.ImageID,
			&member.DateJoin,
			&member.Rank,
			&member.Content.Ru,
			&member.Content.Eng,
		)

		if err != nil {
			return nil, err
		}

		members = append(members, &member)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}

func (cs *CouncilStorage) FindOne(memberIdStr string) (*models.CouncilMemberRead, error) {
	var member models.CouncilMemberRead
	query := fmt.Sprintf(
		"SELECT member_id, %s FROM %s WHERE member_id = $1",
		getColumns(),
		cs.table,
	)

	row := cs.db.QueryRow(query, memberIdStr)

	err := row.Scan(
		&member.Id,
		&member.Name.Ru,
		&member.Name.Ru,
		&member.Description.Ru,
		&member.Description.Eng,
		&member.Location.Ru,
		&member.Location.Eng,
		&member.Email,
		&member.ScopusURL,
		&member.ImageID,
		&member.DateJoin,
		&member.Rank,
		&member.Content.Ru,
		&member.Content.Eng,
	)

	if err != nil {
		return nil, err
	}

	return &member, nil
}
