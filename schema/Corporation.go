package schema

import (
	"time"

	"gorm.io/gorm"
)

type Corporation struct {
	Id          string `json:"id" gorm:"primaryKey;size:50;"`
	Code        string `json:"code" gorm:"size:10;unique"`
	Name        string `gorm:"size:100;" json:"name"`
	Description string `gorm:"type:text;" json:"description"`
	PhoneNumber string `gorm:"size:15;" json:"phone_number"`
	Fax         string `gorm:"size:15;" json:"fax"`
	Sector      string `gorm:"type:text;" json:"sector"`
	Domain      string `gorm:"size:20;" json:"domain"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	IsDeleted   bool      `gorm:"type:bool;default:false" json:"delete"`
	IsActive    bool      `gorm:"type:bool;default:true" json:"status"`
	Profiles    []Profile `gorm:"many2many:profile_corporations;"`
}

func (Corporation) TableName() string {
	return "corporations"
}

type UpdateCorporation struct {
	Code        string `json:"code" `
	Name        string `json:"name"`
	Description string `json:"description"`
	PhoneNumber string `json:"phone_number"`
	Fax         string `json:"fax"`
	Sector      string `json:"sector"`
	Domain      string `json:"domain"`
}
