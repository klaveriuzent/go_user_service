package schema

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	Id        string `json:"id" gorm:"primaryKey;size:50;"`
	UserId    string `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	IsDeleted bool `gorm:"type:bool;default:false" json:"delete"`
	IsActive  bool `gorm:"type:bool;default:true" json:"status"`
}

func (Account) TableName() string {
	return "accounts"
}

type AssignRoleApplication struct {
	RoleApplication []string `json:"role_application"`
}
