package councilst

import (
	"database/sql"
	"urfu-radio-journal/internal/models"
)

type CouncilStorage struct {
	db *sql.DB
}

func NewCouncilStorage(db *sql.DB) *CouncilStorage {
	return &CouncilStorage{
		db: db,
	}
}

func (cs *CouncilStorage) InsertOne(member *models.CouncilMemberCreate) (string, error) {
	stmt, err := cs.db.Prepare("INSERT INTO editorial_board (fullname_ru, fullname_en, description_ru, description_en, location_ru, location_en, email, scopus, photo_path, date_join, rank, content_ru, content_en) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING member_id")

	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var id string
	err = stmt.QueryRow(member.Name.Ru, member.Name.Eng, member.Description.Ru, member.Description.Eng,
		member.Location.Ru, member.Location.Eng, member.Email, member.ScopusURL, member.ImagePathId,
		member.DateJoin, member.Rank, member.Content.Ru, member.Content.Eng).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (cs *CouncilStorage) UpdateOne(memberIdStr string, memberUpdate *models.CouncilMemberUpdate) error {
	_, err := cs.db.Exec("UPDATE editorial_board SET fullname_ru = $1, fullname_en = $2, description_ru = $3, description_en = $4, location_ru = $5, location_en = $6, email = $7, scopus = $8, photo_path = $9, date_join = $10, rank = $11, content_ru = $12, content_en = $13 WHERE member_id = $14",
		memberUpdate.Name.Ru, memberUpdate.Name.Eng, memberUpdate.Description.Ru, memberUpdate.Description.Eng,
		memberUpdate.Location.Ru, memberUpdate.Location.Eng, memberUpdate.Email, memberUpdate.ScopusURL, memberUpdate.ImagePathId,
		memberUpdate.DateJoin, memberUpdate.Rank, memberUpdate.Content.Ru, memberUpdate.Content.Eng, memberIdStr)

	if err != nil {
		return err
	}

	return nil
}

func (cs *CouncilStorage) Delete(idStr string) error {
	_, err := cs.db.Exec("DELETE FROM editorial_board WHERE member_id = $1", idStr)

	if err != nil {
		return err
	}
	return nil
}

func (cs *CouncilStorage) GetImagePathId(idStr string) (string, error) {
	var imagePathId string
	err := cs.db.QueryRow("SELECT photo_path FROM editorial_board WHERE member_id = $1", idStr).Scan(&imagePathId)
	if err != nil {
		return "", err
	}
	return imagePathId, nil
}

func (cs *CouncilStorage) GetAll() ([]*models.CouncilMemberRead, error) {
	rows, err := cs.db.Query("SELECT member_id, fullname_ru, fullname_en, description_ru, description_en, location_ru, location_en, email, scopus, photo_path, date_join, rank, content_ru, content_en FROM editorial_board")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []*models.CouncilMemberRead
	for rows.Next() {
		var member models.CouncilMemberRead
		err = rows.Scan(&member.Id, &member.Name.Ru, &member.Name.Ru, &member.Description.Ru, &member.Description.Eng,
			&member.Location.Ru, &member.Location.Eng, &member.Email, &member.ScopusURL, &member.ImagePathId, &member.DateJoin, &member.Rank,
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

func (cs *CouncilStorage) FindOne(memberIdStr string) (*models.CouncilMemberRead, error) {
	var member models.CouncilMemberRead
	err := cs.db.QueryRow("SELECT member_id, fullname_ru, fullname_en, description_ru, description_en, location_ru, location_en, email, scopus, photo_path, date_join, rank, content_ru, content_en FROM editorial_board WHERE member_id = $1", memberIdStr).
		Scan(&member.Id, &member.Name.Ru, &member.Name.Ru, &member.Description.Ru, &member.Description.Eng,
			&member.Location.Ru, &member.Location.Eng, &member.Email, &member.ScopusURL, &member.ImagePathId, &member.DateJoin, &member.Rank,
			&member.Content.Ru, &member.Content.Eng)

	if err != nil {
		return nil, err
	}

	return &member, nil
}
