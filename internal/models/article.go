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
	FilePathId     string    `json:"filePathId" bson:"filePathId" binding:"required"`
	VideoPathId    string    `json:"videoPathId" bson:"videoPathId" binding:"required"`
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
	FilePathId     string    `json:"filePathId" bson:"filePathId" binding:"required"`
	VideoPathId    string    `json:"videoPathId" bson:"videoPathId" binding:"required"`
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
	FilePathId     string    `json:"filePathId" bson:"filePathId,omitempty" binding:"-"`
	Literature     []string  `json:"literature" bson:"literature,omitempty" binding:"-"`
	VideoPathId    string    `json:"videoPathId" bson:"videoPathId,omitempty" binding:"-"`
	Reference      Text      `json:"reference" bson:"reference,omitempty" binding:"-"`
	DateReceipt    time.Time `json:"dateReceipt" bson:"dateReceipt,omitempty" binding:"-"`
	DateAcceptance time.Time `json:"dateAcceptance" bson:"dateAcceptance,omitempty" binding:"-"`
	DOI            string    `json:"doi" bson:"doi,omitempty" binding:"-"`
}
