package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	Id             uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Email          string    `json:"email" gorm:"primaryKey;type:varchar(100)"`
	Username       string    `json:"username" gorm:"primaryKey;type:varchar(100)"`
	NickName       string    `gorm:"size:255;" json:"nickname"`
	FullName       string    `gorm:"size:255;" json:"fullname"`
	Picture        string    `gorm:"size:255;" json:"picture"`
	PhoneNumber    string    `gorm:"size:13;" json:"phone_number"`
	Address        string    `gorm:"type:text;" json:"address"`
	Token          string    `gorm:"size:255;" json:"token"`
	TokenExpired   time.Time
	CompanyCode    string `gorm:"size:10;" json:"company_code"`
	CompanyName    string `gorm:"size:10;" json:"company_name"`
	CompanyAddress string `gorm:"type:text;" json:"company_address"`
	CompanyDomain  string `gorm:"type:text;" json:"domain"`
	Department     string `gorm:"type:text;" json:"department"`
	OfficeNumber   string `gorm:"type:text;" json:"office_number"`
	Expired        int    `gorm:"type:int;" json:"expired_time"`
	CreatedBy      string `gorm:"size:13;" json:"created_by"`
	ExpiredAt      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Deleted        bool `gorm:"type:bool;default:false" json:"deleted"`
	UserId         string
	IsActive       bool `gorm:"type:bool;default:true" json:"status"`
}

func (Profile) TableName() string {
	return "profiles"
}
