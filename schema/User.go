package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        string `json:"id" gorm:"primaryKey;size:50;"`
	Email     string `gorm:"size:255;not null;unique" json:"email"`
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:255;not null;" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	IsDeleted bool      `gorm:"type:bool;default:false" json:"delete"`
	IsActive  bool      `gorm:"type:bool;default:true" json:"status"`
	Roles     []Role    `gorm:"many2many:user_roles;"`
	Profiles  []Profile `gorm:"foreignKey:UserId"`
	Accounts  []Account `gorm:"foreignKey:UserId"`
}

func (User) TableName() string {
	return "users"
}

type AssignRole struct {
	Role []string `json:"role"`
}
