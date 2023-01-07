package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRoleApplications struct {
	RoleId    uint      `gorm:"primaryKey"`
	AccountId uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (AccountRoleApplications) TableName() string {
	return "account_role_applications"
}
