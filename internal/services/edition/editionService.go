package editionsrv

import (
	"fmt"
	"urfu-radio-journal/internal/models"
)

type storage interface {
	InsertOne(*models.EditionCreate) (string, error)
	GetAll() ([]*models.EditionRead, error)
	FindOne(string) (*models.EditionRead, error)
	UpdateOne(*models.EditionUpdate) error
	Delete(string) error
}

type EditionService struct {
	repo storage
}

func NewEditionService(storage storage) *EditionService {
	return &EditionService{
		repo: storage,
	}
}

func (es *EditionService) Create(edition *models.EditionCreate) (id string, err error) {
	return es.repo.InsertOne(edition)
}

func (es *EditionService) GetAll() (editions []*models.EditionRead, err error) {
	return es.repo.GetAll()
}

func (es *EditionService) Get(id string) (*models.EditionRead, error) {
	if id == "" {
		return nil, fmt.Errorf("empty edition id")
	}
	return es.repo.FindOne(id)
}

func (es *EditionService) Update(newEdition *models.EditionUpdate) error {
	return es.repo.UpdateOne(newEdition)
}

func (es *EditionService) Delete(editionId string) error {
	return es.repo.Delete(editionId)
}
