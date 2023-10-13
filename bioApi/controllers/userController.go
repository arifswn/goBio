package controllers

import (
	"goBio/config"
	"goBio/model"
)

func CreateUser(user *model.Users) error {
	db := config.DB
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *model.Users, email string) error {
	db := config.DB
	if err := db.Save(user).Where("email = ?", email).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *model.Users) error {
	db := config.DB
	if err := db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func GetOneByEmail(email string) (model.Users, error) {
	var user model.Users
	db := config.GetDB()
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetAll(keyword string) ([]model.Users, error) {
	users := []model.Users{}
	db := config.GetDB()

	if err := db.Where("email LIKE ? OR nama LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
