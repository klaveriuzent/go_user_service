package model

import (
	"userservice/database"
	"userservice/schema"
)

type Applications schema.Application
type UpdateApplication schema.UpdateApplication

func (entry *Applications) Save() (*Applications, error) {
	err := database.Database.Create(&entry).Error
	return entry, err
}

func (update_data *Applications) ChangeData(id string, ua UpdateApplication) (Applications, error) {
	err := database.Database.Model(Applications{}).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return *update_data, err
	}
	res, _ := ApplicationFindById(id)
	return res, nil
}

func ApplicationFindById(id string) (Applications, error) {
	var ar Applications
	err := database.Database.Where("id = ?", id).First(&ar).Error
	return ar, err
}
