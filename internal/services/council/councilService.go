package councilsrv

import (
	"urfu-radio-journal/internal/models"
)

type storage interface {
	InsertOne(*models.CouncilMemberCreate) (string, error)
	GetAll() ([]*models.CouncilMemberRead, error)
	FindOne(string) (*models.CouncilMemberRead, error)
	UpdateOne(string, *models.CouncilMemberUpdate) error
	Delete(string) error
}

type CouncilService struct {
	repo storage
}

func NewCouncilService(storage storage) *CouncilService {
	return &CouncilService{
		repo: storage,
	}
}

func (cs *CouncilService) Create(member *models.CouncilMemberCreate) error {
	_, err := cs.repo.InsertOne(member)
	return err
}

func (cs *CouncilService) Update(memberIdStr string, memberUpdate *models.CouncilMemberUpdate) error {
	return cs.repo.UpdateOne(memberIdStr, memberUpdate)
}

func (cs *CouncilService) Delete(id string) error {
	return cs.repo.Delete(id)
}

func (cs *CouncilService) GetAll() ([]*models.CouncilMemberRead, error) {
	return cs.repo.GetAll()
}

func (cs *CouncilService) Get(memberIdStr string) (*models.CouncilMemberRead, error) {
	return cs.repo.FindOne(memberIdStr)
}
