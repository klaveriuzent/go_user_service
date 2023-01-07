package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	Id            uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	UserId        string    `json:"user_id"`
	ApplicationId uuid.UUID `json:"application_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	IsActive      bool              `gorm:"type:bool;default:true" json:"status"`
	Roles         []RoleApplication `gorm:"many2many:account_role_applications;"`
}

func (Account) TableName() string {
	return "accounts"
}
