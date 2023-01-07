package schema

import (
	"time"

	"gorm.io/gorm"
)

type RoleApplication struct {
	Id          uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT;not_null"`
	Name        string `gorm:"size:50;" json:"name"`
	Description string `gorm:"size:255;" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	IsActive    bool `gorm:"type:bool;default:true" json:"status"`
}

func (RoleApplication) TableName() string {
	return "role_applications"
}
