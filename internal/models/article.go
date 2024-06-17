package models

import (
	"time"
)

type ArticleCreate struct {
	EditionId      int       `json:"editionId" bson:"editionId" binding:"required"`
	Title          Text      `json:"title" bson:"title" binding:"required"`
	Authors        []Author  `json:"authors" bson:"authors" binding:"required"`
	Content        Text      `json:"content" bson:"content" binding:"required"`
	Keywords       []Text    `json:"keywords" bson:"keywords" binding:"required"`
	DocumentID     string    `json:"documentID" bson:"documentID" binding:"required"`
	VideoID        string    `json:"videoID" bson:"videoID" binding:"required"`
	Literature     []string  `json:"literature" bson:"literature" binding:"required"`
	Reference      Text      `json:"reference" bson:"reference" binding:"required"`
	DateReceipt    time.Time `json:"dateReceipt" bson:"dateReceipt" binding:"required"`
	DateAcceptance time.Time `json:"dateAcceptance" bson:"dateAcceptance" binding:"required"`
	DOI            string    `json:"doi" bson:"doi" binding:"required"`
}

type ArticleRead struct {
	Id             int       `json:"id" bson:"_id" binding:"required"`
	EditionId      int       `json:"editionId" bson:"editionId" binding:"required"`
	Title          Text      `json:"title" bson:"title" binding:"required"`
	Authors        []Author  `json:"authors" bson:"authors" binding:"required"`
	Content        Text      `json:"content" bson:"content" binding:"required"`
	Keywords       []Text    `json:"keywords" bson:"keywords" binding:"required"`
	DocumentID     string    `json:"documentID" bson:"documentID" binding:"required"`
	VideoID        string    `json:"videoID" bson:"videoID" binding:"required"`
	Literature     []string  `json:"literature" bson:"literature" binding:"required"`
	Reference      Text      `json:"reference" bson:"reference" binding:"required"`
	DateReceipt    time.Time `json:"dateReceipt" bson:"dateReceipt" binding:"required"`
	DateAcceptance time.Time `json:"dateAcceptance" bson:"dateAcceptance" binding:"required"`
	DOI            string    `json:"doi" bson:"doi" binding:"required"`
}

type ArticleUpdate struct {
	Id             int       `json:"id" binding:"required"`
	Title          Text      `json:"title" bson:"title,omitempty" binding:"-"`
	Authors        []Author  `json:"authors" bson:"authors,omitempty" binding:"-"`
	Content        Text      `json:"content" bson:"content,omitempty" binding:"-"`
	Keywords       []Text    `json:"keywords" bson:"keywords,omitempty" binding:"-"`
	DocumentID     string    `json:"documentID" bson:"documentID,omitempty" binding:"-"`
	Literature     []string  `json:"literature" bson:"literature,omitempty" binding:"-"`
	VideoID        string    `json:"videoID" bson:"videoID,omitempty" binding:"-"`
	Reference      Text      `json:"reference" bson:"reference,omitempty" binding:"-"`
	DateReceipt    time.Time `json:"dateReceipt" bson:"dateReceipt,omitempty" binding:"-"`
	DateAcceptance time.Time `json:"dateAcceptance" bson:"dateAcceptance,omitempty" binding:"-"`
	DOI            string    `json:"doi" bson:"doi,omitempty" binding:"-"`
}

type ArticleQuery struct {
	BatchArgs
	ArticleSearch
}

type ArticleSearch struct {
	EditionID int    `form:"editionId"`
	Search    string `form:"search" binding:"required_if=EditionID 0"`
	Title     *bool  `form:"title"`
	//Keywords  *bool  `form:"keywords"`
	//Authors    *bool  `form:"authors"`
	Affilation *bool `form:"affilation"`
}
