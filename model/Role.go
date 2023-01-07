package model

import (
	"fmt"
	"userservice/database"
	"userservice/schema"
)

type Role schema.Role

func (role *Role) Save() (*Role, error) {
	err := database.Database.Create(&role).Error
	return role, err
}

func FindRoleByName(name string) (Role, error) {
	var role Role
	err := database.Database.Where("name=?", name).First(&role).Error
	return role, err
}

func FindRoleById(id string) (Role, error) {
	var role Role
	err := database.Database.Where("id=?", id).First(&role).Error
	fmt.Println(role)
	return role, err
}

func FindRoleMapById(params []int) ([]Role, error) {
	var role []Role
	err := database.Database.Where("id IN ?", params).Find(&role).Error
	fmt.Println(role)
	return role, err
}
func FindRoleMapByName(params []string) ([]Role, error) {
	var role []Role
	err := database.Database.Where("name IN ?", params).Find(&role).Error
	fmt.Println(role)
	return role, err
}
