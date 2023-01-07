package schema

import (
	"time"

	"gorm.io/gorm"
)

type RoleApplication struct {
	Id          uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT;not_null"`
	Name        string `gorm:"size:50;" json:"name"`
	Description string `gorm:"type:text;" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	IsDeleted   bool `gorm:"type:bool;default:false" json:"delete"`
	IsActive    bool `gorm:"type:bool;default:true" json:"status"`
}

func (RoleApplication) TableName() string {
	return "role_applications"
}
