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

func FindProfileByUserId(id string) (Profile, error) {
	var emp Profile
	err := database.Database.Where("user_id=?", id).Where("is_active=?", true).First(&emp).Error
	return emp, err
}

func FindProfileById(id string) (Profile, error) {
	var emp Profile
	err := database.Database.Where("id=?", id).First(&emp).Error
	return emp, err
}

func FindProfileByName(id string) (Profile, error) {
	var emp Profile
	err := database.Database.Where("username=?", id).First(&emp).Error
	return emp, err
}

func FindProfileByEmail(id string) (Profile, error) {
	var emp Profile
	err := database.Database.Where("email=?", id).First(&emp).Error
	return emp, err
}

func (profile *Profile) ProfileAssignCorporation(id string, corps []schema.Corporation) (Profile, error) {
	err := database.Database.Model(&profile).Association("Corporations").Replace(corps)
	if err != nil {
		return *profile, err
	}
	res, _ := FindProfileById(id)
	return res, nil
}
