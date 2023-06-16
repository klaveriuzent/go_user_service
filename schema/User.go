package schema

import (
	"time"
)

type User struct {
	Id          string `json:"id" gorm:"primaryKey;size:50;"`
	Email       string `gorm:"size:255;not null;unique" json:"email"`
	Username    string `gorm:"size:255;not null;unique" json:"username"`
	Password    string `gorm:"size:255;not null;" json:"-"`
	LastLoginAt *time.Time
	Accounts    []Account `gorm:"foreignKey:UserId"`
}

func (User) TableName() string {
	return "users"
}
