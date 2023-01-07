package model

import (
	"fmt"
	"userservice/database"
	"userservice/schema"
)

type Profile schema.Profile

func (emp *Profile) Save() (*Profile, error) {
	err := database.Database.Create(&emp).Error
	return emp, err
}

func FindProfileAll() ([]Profile, error) {
	var emp []Profile
	err := database.Database.Find(&emp).Error
	fmt.Println(emp)
	return emp, err
}

func FindEmployeeById(id string) (Profile, error) {
	var emp Profile
	err := database.Database.Where("id=?", id).First(&emp).Error
	return emp, err
}
