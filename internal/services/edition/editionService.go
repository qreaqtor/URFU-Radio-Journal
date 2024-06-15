package editionsrv

import (
	"fmt"
	"urfu-radio-journal/internal/models"
)

type storage interface {
	InsertOne(*models.EditionCreate) (string, error)
	GetAll(*models.BatchArgs) ([]*models.EditionRead, error)
	FindOne(string) (*models.EditionRead, error)
	UpdateOne(*models.EditionUpdate) error
	Delete(string) error
	GetCount() (int, error)
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

func (es *EditionService) GetAll(args *models.BatchArgs) ([]*models.EditionRead, int, error) {
	result, err := es.repo.GetAll(args)
	if err != nil {
		return nil, 0, err
	}

	count, err := es.repo.GetCount()
	if err != nil {
		return nil, 0, err
	}

	return result, count, nil
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
