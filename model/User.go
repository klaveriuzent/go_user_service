package model

import (
	"html"
	"strings"
	"userservice/database"
	"userservice/schema"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User schema.User

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	return user, err
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?", username).First(&user).Error
	return user, err
}

func FindUserById(id string) (User, error) {
	var user User
	err := database.Database.Preload("ActivityLogs").Preload("Accounts.Application").Preload("Accounts.RoleApplications").Preload(clause.Associations).Where("id = ?", id).Preload("Profiles", "is_active NOT IN (?)", false).First(&user).Error
	return user, err
}

func (user *User) UserGetCount() int64 {
	var result int64
	database.Database.Model(&user).Count(&result)
	return result
}
