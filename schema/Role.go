package schema

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id          uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT;not_null"`
	Name        string `gorm:"size:50;" json:"name"`
	Description string `gorm:"size:255;" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	IsActive    bool `gorm:"type:bool;default:true" json:"status"`
}

func (Role) TableName() string {
	return "roles"
}
