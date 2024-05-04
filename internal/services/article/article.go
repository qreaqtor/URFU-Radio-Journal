package articlesrv

import (
	"errors"
	"urfu-radio-journal/internal/models"
)

var (
	errBadId = errors.New("wrong id's length")
)

type storage interface {
	InsertOne(*models.ArticleCreate) (string, error)
	Find(string) ([]*models.ArticleRead, error)
	FindOne(string) (*models.ArticleRead, error)
	UpdateOne(*models.ArticleUpdate) error
	Delete(string) error
}

type ArticleService struct {
	repo storage
}

func NewArticleService(repo storage) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
}

func (as *ArticleService) Create(article *models.ArticleCreate) (string, error) {
	id, err := as.repo.InsertOne(article)
	return id, err
}

func (as *ArticleService) GetAll(editionIdStr string) ([]*models.ArticleRead, error) {
	if editionIdStr == "" {
		return nil, errBadId
	}
	return as.repo.Find(editionIdStr)
}

func (as *ArticleService) Get(articleIdStr string) (*models.ArticleRead, error) {
	if articleIdStr == "" {
		return nil, errBadId
	}
	return as.repo.FindOne(articleIdStr)
}

func (as *ArticleService) Update(newArticle *models.ArticleUpdate) error {
	return as.repo.UpdateOne(newArticle)
}

func (as *ArticleService) Delete(id string) error {
	if id == "" {
		return errBadId
	}
	return as.repo.Delete(id)
}

// func (as *ArticleService) GetIdsByEditionId(editionId primitive.ObjectID) (articlesId, filePathsId []primitive.ObjectID, err error) {
// 	articlesId = make([]primitive.ObjectID, 0)
// 	filePathsId = make([]primitive.ObjectID, 0)
// 	filter := bson.M{"editionId": editionId}
// 	cur, err := as.repo.Find(as.ctx, filter)
// 	if err != nil {
// 		return
// 	}
// 	var res []models.ArticleRead
// 	if err = cur.All(as.ctx, &res); err != nil {
// 		return
// 	}
// 	for _, v := range res {
// 		articlesId = append(articlesId, v.Id)
// 		filePathsId = append(filePathsId, v.FilePathId)
// 	}
// 	return
// }

// func (as *ArticleService) GetFilePathId(id primitive.ObjectID) (filePathId primitive.ObjectID, err error) {
// 	filter := bson.M{"_id": id}
// 	var article models.ArticleRead
// 	err = as.repo.FindOne(as.ctx, filter).Decode(&article)
// 	if err != nil {
// 		return
// 	}
// 	filePathId = article.FilePathId
// 	return
// }
