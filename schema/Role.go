package schema

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id          uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT;not_null"`
	Name        string `gorm:"size:50;" json:"name"`
	Description string `gorm:"type:text;" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	IsDeleted   bool `gorm:"type:bool;default:true" json:"delete"`
	IsActive    bool `gorm:"type:bool;default:true" json:"status"`
}

func (Role) TableName() string {
	return "roles"
}
