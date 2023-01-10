package model

import (
	"fmt"
	"userservice/database"
	"userservice/schema"
)

type Account schema.Account
type UpdateAccount schema.UpdateAccount

func (emp *Account) Save() (*Account, error) {
	err := database.Database.Create(&emp).Error
	return emp, err
}

func FindAccountAll() ([]Account, error) {
	var emp []Account
	err := database.Database.Find(&emp).Error
	fmt.Println(emp)
	return emp, err
}

func FindAccountByUserId(id string) (Account, error) {
	var emp Account
	err := database.Database.Where("user_id=?", id).Where("is_active=?", true).First(&emp).Error
	return emp, err
}

func FindAccountById(id string) (Account, error) {
	var emp Account
	err := database.Database.Where("id=?", id).First(&emp).Error
	return emp, err
}

func (account *Account) AccountAssignRoleApplication(id string, corps []schema.RoleApplication) (Account, error) {
	err := database.Database.Model(&account).Association("RoleApplications").Replace(corps)
	if err != nil {
		return *account, err
	}
	res, _ := FindAccountById(id)
	return res, nil
}

func (update_data *Account) ChangeData(id string, ua UpdateAccount) (Account, error) {
	err := database.Database.Model(&update_data).Where("id = ?", id).Updates(ua).Error
	if err != nil {
		return *update_data, err
	}
	res, _ := FindAccountById(id)
	return res, nil
}
