package schema

import (
	"time"
)

type Account struct {
	Id          string `json:"id" gorm:"primaryKey;size:50;"`
	UserId      string `json:"user_id"`
	CreatedAt   time.Time
	IsActive    bool `gorm:"type:bool;default:true" json:"status"`
	LastLoginAt *time.Time
}

func (Account) TableName() string {
	return "accounts"
}
