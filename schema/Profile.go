package schema

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	Id           string    `json:"id" gorm:"primaryKey;size:50;"`
	Email        string    `gorm:"size:255;not null;unique" json:"email"`
	Username     string    `gorm:"size:255;not null;unique" json:"username"`
	Nik          string    `gorm:"size:20;" json:"nik"`
	NickName     string    `gorm:"size:255;" json:"nickname"`
	FullName     string    `gorm:"size:255;" json:"fullname"`
	Picture      string    `gorm:"size:255;" json:"picture"`
	PhoneNumber  string    `gorm:"size:15;" json:"phone_number"`
	Addresses    []Address `gorm:"foreignKey:ProfileId"`
	Token        string    `gorm:"size:255;" json:"token"`
	TokenExpired time.Time
	Department   string `gorm:"type:text;" json:"department"`
	Expired      int    `gorm:"type:int;" json:"expired_time"`
	CreatedBy    string `gorm:"size:13;" json:"created_by"`
	ExpiredAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	UserId       string        `json:"user_id"`
	Deleted      bool          `gorm:"type:bool;default:false" json:"delete"`
	IsActive     bool          `gorm:"type:bool;default:true" json:"status"`
	Corporations []Corporation `gorm:"many2many:profile_corporations;"`
}

func (Profile) TableName() string {
	return "profiles"
}

type AssignCorporation struct {
	Corporation []string `json:"corporation"`
}
