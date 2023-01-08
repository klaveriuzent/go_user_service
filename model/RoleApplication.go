package model

import (
	"fmt"
	"userservice/database"
	"userservice/schema"
)

type RoleApplication schema.RoleApplication

func (role *RoleApplication) Save() (*RoleApplication, error) {
	err := database.Database.Create(&role).Error
	return role, err
}

func FindRoleApplicationByName(name string) (RoleApplication, error) {
	var role RoleApplication
	err := database.Database.Where("name=?", name).First(&role).Error
	return role, err
}

func FindRoleApplicationById(id string) (RoleApplication, error) {
	var role RoleApplication
	err := database.Database.Where("id=?", id).First(&role).Error
	fmt.Println(role)
	return role, err
}

func FindRoleApplicationMapById(params []int) ([]RoleApplication, error) {
	var role []RoleApplication
	err := database.Database.Where("id IN ?", params).Find(&role).Error
	fmt.Println(role)
	return role, err
}
func FindRoleApplicationMapByName(params []string) ([]RoleApplication, error) {
	var role []RoleApplication
	err := database.Database.Where("name IN ?", params).Find(&role).Error
	fmt.Println(role)
	return role, err
}
