package redactionst

import (
	"database/sql"
	"fmt"
	"urfu-radio-journal/internal/models"
)

type RedactionStorage struct {
	db    *sql.DB
	table string
}

func NewRedactionStorage(db *sql.DB, table string) *RedactionStorage {
	return &RedactionStorage{
		db:    db,
		table: table,
	}
}

func getColumns() string {
	return "fullname_ru, fullname_en, description_ru, description_en, location_ru, location_en, email, photo_path, date_join, rank, content_ru, content_en"
}

func generateValuesID(count int) string {
	res := ""
	for i := 1; i < count; i++ {
		res += fmt.Sprintf("$%v, ", i)
	}
	res += fmt.Sprintf("$%v", count)
	return res
}

func (rs *RedactionStorage) InsertOne(member *models.RedactionMemberCreate) (string, error) {
	var id string

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING member_id;",
		rs.table,
		getColumns(),
		generateValuesID(12),
	)

	row := rs.db.QueryRow(
		query,
		member.Name.Ru,
		member.Name.Eng,
		member.Description.Ru,
		member.Description.Eng,
		member.Location.Ru,
		member.Location.Eng,
		member.Email,
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

func (rs *RedactionStorage) UpdateOne(memberIdStr string, memberUpdate *models.RedactionMemberUpdate) error {
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
		photo_path = $8, 
		date_join = $9, 
		rank = $10, 
		content_ru = $11, 
		content_en = $12 
		WHERE 
		member_id = $13`,
		rs.table,
	)
	_, err := rs.db.Exec(
		query,
		memberUpdate.Name.Ru,
		memberUpdate.Name.Eng,
		memberUpdate.Description.Ru,
		memberUpdate.Description.Eng,
		memberUpdate.Location.Ru,
		memberUpdate.Location.Eng,
		memberUpdate.Email,
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

func (rs *RedactionStorage) Delete(idStr string) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE member_id = $1",
		rs.table,
	)
	_, err := rs.db.Exec(query, idStr)

	if err != nil {
		return err
	}
	return nil
}

func (rs *RedactionStorage) GetImageID(idStr string) (string, error) {
	var imageID string

	query := fmt.Sprintf(
		"SELECT photo_path FROM %s WHERE member_id = $1",
		rs.table,
	)

	row := rs.db.QueryRow(query, idStr)
	err := row.Scan(&imageID)

	if err != nil {
		return "", err
	}
	return imageID, nil
}

func (rs *RedactionStorage) GetAll() ([]*models.RedactionMemberRead, error) {
	var members []*models.RedactionMemberRead

	query := fmt.Sprintf(
		"SELECT member_id, %s FROM %s",
		getColumns(),
		rs.table,
	)

	rows, err := rs.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var member models.RedactionMemberRead

		err = rows.Scan(
			&member.Id,
			&member.Name.Ru,
			&member.Name.Eng,
			&member.Description.Ru,
			&member.Description.Eng,
			&member.Location.Ru,
			&member.Location.Eng,
			&member.Email,
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

func (rs *RedactionStorage) FindOne(memberIdStr string) (*models.RedactionMemberRead, error) {
	var member models.RedactionMemberRead

	query := fmt.Sprintf(
		"SELECT member_id, %s FROM %s WHERE member_id = $1",
		getColumns(),
		rs.table,
	)

	row := rs.db.QueryRow(query, memberIdStr)

	err := row.Scan(
		&member.Id,
		&member.Name.Ru,
		&member.Name.Eng,
		&member.Description.Ru,
		&member.Description.Eng,
		&member.Location.Ru,
		&member.Location.Eng,
		&member.Email,
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
