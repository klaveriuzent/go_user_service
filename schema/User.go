package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:255;not null;" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Deleted   bool   `gorm:"type:bool" json:"deleted"`
	IsActive  bool   `gorm:"type:bool;default:true" json:"status"`
	Roles     []Role `gorm:"many2many:user_roles;"`
	Profiles  []Profile
	Accounts  []Account `gorm:"foreignKey:UserId"`
	Articles  []Article `gorm:"foreignKey:UserId"`
}

func (User) TableName() string {
	return "users"
}

type AssignRole struct {
	Role []string `json:"role"`
}
