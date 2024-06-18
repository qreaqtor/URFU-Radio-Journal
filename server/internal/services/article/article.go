package articlesrv

import (
	"errors"
	"urfu-radio-journal/internal/models"
)

var (
	errBadId = errors.New("wrong id's length")
)

type articleStorage interface {
	InsertOne(*models.ArticleCreate) (string, error)
	Find(*models.ArticleQuery) ([]*models.ArticleRead, error)
	FindOne(string) (*models.ArticleRead, error)
	UpdateOne(*models.ArticleUpdate) error
	Delete(string) error
	GetCount(*models.ArticleSearch) (int, error)
}

type ArticleService struct {
	articleRepo articleStorage
}

func NewArticleService(articleRepo articleStorage) *ArticleService {
	return &ArticleService{
		articleRepo: articleRepo,
	}
}

func (as *ArticleService) Create(article *models.ArticleCreate) (string, error) {
	id, err := as.articleRepo.InsertOne(article)
	if err != nil {
		return "", err
	}

	return id, err
}

func (as *ArticleService) GetAll(args *models.ArticleQuery) ([]*models.ArticleRead, int, error) {
	articles, err := as.articleRepo.Find(args)
	if err != nil {
		return nil, 0, err
	}

	count, err := as.articleRepo.GetCount(&args.ArticleSearch)
	if err != nil {
		return nil, 0, err
	}

	return articles, count, err
}

func (as *ArticleService) Get(articleIdStr string) (*models.ArticleRead, error) {
	if articleIdStr == "" {
		return nil, errBadId
	}

	article, err := as.articleRepo.FindOne(articleIdStr)
	if err != nil {
		return nil, err
	}

	return article, err
}

func (as *ArticleService) Update(newArticle *models.ArticleUpdate) error {
	err := as.articleRepo.UpdateOne(newArticle)
	if err != nil {
		return err
	}

	return err
}

func (as *ArticleService) Delete(id string) error {
	if id == "" {
		return errBadId
	}

	return as.articleRepo.Delete(id)
}
