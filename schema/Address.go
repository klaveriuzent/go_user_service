package schema

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	Id        string `json:"id" gorm:"primaryKey;size:50;"`
	Country   string `gorm:"size:50;" json:"country"`
	State     string `gorm:"size:50;" json:"state"`
	City      string `gorm:"size:50;" json:"city"`
	Address   string `gorm:"type:text" json:"address"`
	ZipCode   string `gorm:"size:7;" json:"zip_code"`
	ProfileId string `json:"profile_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	IsDeleted bool `gorm:"type:bool;default:true" json:"delete"`
	IsActive  bool `gorm:"type:bool;default:true" json:"status"`
}

func (Address) TableName() string {
	return "address"
}

type UpdateAddress struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"UserId" gorm:"type:uuid"`
}
