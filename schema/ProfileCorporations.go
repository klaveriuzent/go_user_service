package schema

type ProfileCorporations struct {
	CorporationId uint   `gorm:"primaryKey"`
	ProfileId     string `gorm:"primaryKey"`
}

func (ProfileCorporations) TableName() string {
	return "profile_corporations"
}
