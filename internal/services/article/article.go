package articlesrv

import (
	"errors"
	"strconv"
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
	GetCount() (int, error)
}

type authorStorage interface {
	InsertMany([]models.Author, string) error
	Find(string) ([]models.Author, error)
	Delete(string) error
}

type ArticleService struct {
	articleRepo articleStorage
	authorRepo  authorStorage
}

func NewArticleService(articleRepo articleStorage, authorRepo authorStorage) *ArticleService {
	return &ArticleService{
		articleRepo: articleRepo,
		authorRepo:  authorRepo,
	}
}

func (as *ArticleService) Create(article *models.ArticleCreate) (string, error) {
	id, err := as.articleRepo.InsertOne(article)
	if err != nil {
		return "", err
	}

	err = as.authorRepo.InsertMany(article.Authors, id)
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

	for _, article := range articles {
		id := strconv.Itoa(article.Id)
		authors, err := as.authorRepo.Find(id)
		if err != nil {
			return nil, 0, err
		}

		article.Authors = authors
	}

	count, err := as.articleRepo.GetCount()
	if err != nil {
		return nil, 0, err
	}

	return articles, count, err
}

func (as *ArticleService) Get(articleIdStr string) (*models.ArticleRead, error) {
	var authors []models.Author

	if articleIdStr == "" {
		return nil, errBadId
	}

	article, err := as.articleRepo.FindOne(articleIdStr)
	if err != nil {
		return nil, err
	}

	authors, err = as.authorRepo.Find(articleIdStr)
	if err != nil {
		return nil, err
	}
	article.Authors = authors

	return article, err
}

func (as *ArticleService) Update(newArticle *models.ArticleUpdate) error {
	err := as.articleRepo.UpdateOne(newArticle)
	if err != nil {
		return err
	}

	err = as.authorRepo.Delete(strconv.Itoa(newArticle.Id))
	if err != nil {
		return err
	}

	id := strconv.Itoa(newArticle.Id)
	return as.authorRepo.InsertMany(newArticle.Authors, id)
}

func (as *ArticleService) Delete(id string) error {
	if id == "" {
		return errBadId
	}
	return as.articleRepo.Delete(id)
}
