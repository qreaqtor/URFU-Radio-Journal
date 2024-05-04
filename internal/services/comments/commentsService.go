package commentsrv

import (
	"fmt"
	"unicode"
	"urfu-radio-journal/internal/models"
)

type storage interface {
	InsertOne(*models.CommentCreate) (string, error)
	GetAll(bool, string) ([]*models.CommentRead, error)
	UpdateOne(*models.CommentUpdate) error
	Delete(string) error
	Approve(*models.CommentApprove, string) error
}

type CommentsService struct {
	repo storage
}

func NewCommentsService(storage storage) *CommentsService {
	return &CommentsService{
		repo: storage,
	}
}

func (cs *CommentsService) Create(comment *models.CommentCreate) error {
	unicodeRange, err := cs.determineLanguage(comment.ContentPart)
	if err != nil {
		return err
	}
	if unicodeRange == unicode.Latin {
		comment.Content.Eng = comment.ContentPart
	} else {
		comment.Content.Ru = comment.ContentPart
	}
	_, err = cs.repo.InsertOne(comment)
	return err
}

func (cs *CommentsService) determineLanguage(str string) (unicodeRange *unicode.RangeTable, err error) {
	ruCount, engCount := 0, 0
	for _, r := range str {
		if unicode.Is(unicode.Cyrillic, r) {
			ruCount++
		} else if unicode.Is(unicode.Latin, r) {
			engCount++
		} else if !unicode.In(r, unicode.Cyrillic, unicode.Latin, unicode.Number, unicode.Space, unicode.Punct) {
			return nil, fmt.Errorf("the string contains an unsupported character: \"%s\"", string(r))
		}
	}
	if ruCount > engCount {
		unicodeRange = unicode.Cyrillic
	} else {
		unicodeRange = unicode.Latin
	}
	return unicodeRange, err
}

func (cs *CommentsService) GetAll(onlyApproved bool, articleIdStr string) ([]*models.CommentRead, error) {
	if articleIdStr == "" {
		return nil, fmt.Errorf("articleId is empty")
	}
	result, err := cs.repo.GetAll(onlyApproved, articleIdStr)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (cs *CommentsService) Update(comment *models.CommentUpdate) error {
	return cs.repo.UpdateOne(comment)
}

func (cs *CommentsService) Delete(id string) error {
	return cs.repo.Delete(id)
}

// func (cs *CommentsService) DeleteManyHandler(filter primitive.M) error {
// 	_, err := cs.repo.DeleteMany(cs.ctx, filter)
// 	return err
// }

func (cs *CommentsService) Approve(commentApprove *models.CommentApprove) error {
	unicodeRange, err := cs.determineLanguage(commentApprove.ContentPart)
	if err != nil {
		return err
	}
	contentField := "content.Ru"
	if unicodeRange == unicode.Latin {
		contentField = "content.Eng"
	}
	return cs.repo.Approve(commentApprove, contentField)
}
