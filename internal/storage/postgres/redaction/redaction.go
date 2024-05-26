package redactionst

import (
	"database/sql"
	"urfu-radio-journal/internal/models"
)

type RedactionStorage struct {
	db *sql.DB
}

func NewRedactionStorage(db *sql.DB) *RedactionStorage {
	return &RedactionStorage{
		db: db,
	}
}

func (rs *RedactionStorage) InsertOne(member *models.RedactionMemberCreate) (string, error) {
	stmt, err := rs.db.Prepare("INSERT INTO redaction (fullname_ru, fullname_en, description_ru, description_en, location_ru, location_en, email, photo_path, date_join, rank, content_ru, content_en) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING member_id")

	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var id string
	err = stmt.QueryRow(member.Name.Ru, member.Name.Eng, member.Description.Ru, member.Description.Eng,
		member.Location.Ru, member.Location.Eng, member.Email, member.ImagePathId,
		member.DateJoin, member.Rank, member.Content.Ru, member.Content.Eng).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (rs *RedactionStorage) UpdateOne(memberIdStr string, memberUpdate *models.RedactionMemberUpdate) error {
	_, err := rs.db.Exec("UPDATE redaction SET fullname_ru = $1, fullname_en = $2, description_ru = $3, description_en = $4, location_ru = $5, location_en = $6, email = $7, photo_path = $8, date_join = $9, rank = $10, content_ru = $11, content_en = $12 WHERE member_id = $13",
		memberUpdate.Name.Ru, memberUpdate.Name.Eng, memberUpdate.Description.Ru, memberUpdate.Description.Eng,
		memberUpdate.Location.Ru, memberUpdate.Location.Eng, memberUpdate.Email, memberUpdate.ImagePathId,
		memberUpdate.DateJoin, memberUpdate.Rank, memberUpdate.Content.Ru, memberUpdate.Content.Eng, memberIdStr)

	if err != nil {
		return err
	}

	return nil
}

func (rs *RedactionStorage) Delete(idStr string) error {
	_, err := rs.db.Exec("DELETE FROM redaction WHERE member_id = $1", idStr)

	if err != nil {
		return err
	}
	return nil
}

func (rs *RedactionStorage) GetImagePathId(idStr string) (string, error) {
	var imagePathId string
	err := rs.db.QueryRow("SELECT photo_path FROM redaction WHERE member_id = $1", idStr).Scan(&imagePathId)
	if err != nil {
		return "", err
	}
	return imagePathId, nil
}

func (rs *RedactionStorage) GetAll() ([]*models.RedactionMemberRead, error) {
	rows, err := rs.db.Query("SELECT member_id, fullname_ru, fullname_en, description_ru, description_en, location_ru, location_en, email, photo_path, date_join, rank, content_ru, content_en FROM redaction")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []*models.RedactionMemberRead
	for rows.Next() {
		var member models.RedactionMemberRead
		err = rows.Scan(&member.Id, &member.Name.Ru, &member.Name.Ru, &member.Description.Ru, &member.Description.Eng,
			&member.Location.Ru, &member.Location.Eng, &member.Email, &member.ImagePathId, &member.DateJoin, &member.Rank,
			&member.Content.Ru, &member.Content.Eng)

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
	err := rs.db.QueryRow("SELECT member_id, fullname_ru, fullname_en, description_ru, description_en, location_ru, location_en, email, photo_path, date_join, rank, content_ru, content_en FROM redaction WHERE member_id = $1", memberIdStr).
		Scan(&member.Id, &member.Name.Ru, &member.Name.Ru, &member.Description.Ru, &member.Description.Eng,
			&member.Location.Ru, &member.Location.Eng, &member.Email, &member.ImagePathId, &member.DateJoin, &member.Rank,
			&member.Content.Ru, &member.Content.Eng)

	if err != nil {
		return nil, err
	}

	return &member, nil
}
