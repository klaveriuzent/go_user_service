package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	Id        string    `json:"id" gorm:"primaryKey;size:50;"`
	Title     string    `gorm:"size:50;" json:"title"`
	UserId    uuid.UUID `json:"user_id"`
	Content   string    `gorm:"size:255;type:text" json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	IsActive  bool `gorm:"type:bool;default:true" json:"status"`
}

func (Article) TableName() string {
	return "articles"
}

type UpdateArticle struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"UserId" gorm:"type:uuid"`
}
