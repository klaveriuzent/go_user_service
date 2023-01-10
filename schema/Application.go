package schema

import (
	"time"

	"gorm.io/gorm"
)

type Application struct {
	Id              string `json:"id" gorm:"primaryKey;size:50;"`
	Name            string `gorm:"size:100;" json:"name"`
	Description     string `gorm:"type:text;" json:"description"`
	UpdatedNote     string `gorm:"type:text;" json:"updated_note"`
	Version         string `gorm:"size:50;" json:"version"`
	AppPackage      string `gorm:"size:100;" json:"app_package"`
	AppPackageClass string `gorm:"size:100;" json:"app_package_class"`
	AssetApk        string `gorm:"size:100;" json:"asset_apk"`
	AssetIcon       string `gorm:"size:100;" json:"asset_icon"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	IsDeleted       bool      `gorm:"type:bool;default:false" json:"delete"`
	IsActive        bool      `gorm:"type:bool;default:true" json:"status"`
	Accounts        []Account `gorm:"foreignKey:ApplicationId"`
}

func (Application) TableName() string {
	return "applications"
}

type UpdateApplication struct {
	Name            string `gorm:"size:100;" json:"name"`
	Description     string `gorm:"size:255;" json:"description"`
	UpdatedNote     string `gorm:"type:text;" json:"updated_note"`
	Version         string `gorm:"size:50;" json:"version"`
	AppPackage      string `gorm:"size:100;" json:"app_package"`
	AppPackageClass string `gorm:"size:100;" json:"app_package_class"`
	AssetApk        string `gorm:"size:100;" json:"asset_apk"`
	AssetIcon       string `gorm:"size:100;" json:"asset_icon"`
}
