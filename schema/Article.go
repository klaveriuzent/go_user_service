package schema

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	Id        string `json:"id" gorm:"primaryKey;size:50;"`
	Title     string `gorm:"size:50;" json:"title"`
	UserId    string `json:"user_id"`
	Content   string `gorm:"type:text" json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	IsDeleted bool `gorm:"type:bool;default:false" json:"delete"`
	IsActive  bool `gorm:"type:bool;default:true" json:"status"`
}

func (Article) TableName() string {
	return "articles"
}

type UpdateArticle struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"UserId" gorm:"type:uuid"`
}
