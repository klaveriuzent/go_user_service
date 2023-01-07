package schema

type AccountRoleApplications struct {
	RoleId    uint   `gorm:"primaryKey"`
	AccountId string `gorm:"primaryKey"`
}

func (AccountRoleApplications) TableName() string {
	return "account_role_applications"
}
