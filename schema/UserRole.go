package schema

type UserRole struct {
	RoleId uint   `gorm:"primaryKey"`
	UserId string `gorm:"primaryKey"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
