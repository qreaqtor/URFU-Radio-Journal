package redactionsrv

import (
	"urfu-radio-journal/internal/models"
)

type storage interface {
	InsertOne(*models.RedactionMemberCreate) (string, error)
	GetAll() ([]*models.RedactionMemberRead, error)
	FindOne(string) (*models.RedactionMemberRead, error)
	UpdateOne(string, *models.RedactionMemberUpdate) error
	Delete(string) error
}

type RedactionService struct {
	repo storage
}

func NewRedactionService(storage storage) *RedactionService {
	return &RedactionService{
		repo: storage,
	}
}

func (rs *RedactionService) Create(member *models.RedactionMemberCreate) error {
	_, err := rs.repo.InsertOne(member)
	return err
}

func (rs *RedactionService) Update(memberIdStr string, memberUpdate *models.RedactionMemberUpdate) error {
	return rs.repo.UpdateOne(memberIdStr, memberUpdate)
}

func (rs *RedactionService) Delete(id string) error {
	return rs.repo.Delete(id)
}

func (rs *RedactionService) GetAll() ([]*models.RedactionMemberRead, error) {
	return rs.repo.GetAll()
}

func (rs *RedactionService) Get(memberIdStr string) (*models.RedactionMemberRead, error) {
	return rs.repo.FindOne(memberIdStr)
}
