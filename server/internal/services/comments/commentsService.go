package commentsrv

import (
	"unicode"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/internal/utils"
)

type storage interface {
	InsertOne(*models.CommentCreate) (string, error)
	GetAll(*models.CommentQuery) ([]*models.CommentRead, error)
	UpdateOne(*models.CommentUpdate) error
	Delete(string) error
	Approve(*models.CommentApprove, string) error
	GetCount(bool) (int, error)
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
	unicodeRange := utils.DetermineLanguage(comment.ContentPart)
	if unicodeRange == unicode.Latin {
		comment.Content.Eng = comment.ContentPart
	} else {
		comment.Content.Ru = comment.ContentPart
	}
	_, err := cs.repo.InsertOne(comment)
	return err
}

func (cs *CommentsService) GetAll(args *models.CommentQuery) ([]*models.CommentRead, int, error) {
	result, err := cs.repo.GetAll(args)
	if err != nil {
		return nil, 0, err
	}

	count, err := cs.repo.GetCount(args.OnlyApproved)
	if err != nil {
		return nil, 0, err
	}

	return result, count, nil
}

func (cs *CommentsService) Update(comment *models.CommentUpdate) error {
	return cs.repo.UpdateOne(comment)
}

func (cs *CommentsService) Delete(id string) error {
	return cs.repo.Delete(id)
}

func (cs *CommentsService) Approve(commentApprove *models.CommentApprove) error {
	unicodeRange := utils.DetermineLanguage(commentApprove.ContentPart)
	contentField := "content_ru"
	if unicodeRange == unicode.Latin {
		contentField = "content_en"
	}
	return cs.repo.Approve(commentApprove, contentField)
}
