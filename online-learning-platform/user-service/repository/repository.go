package repository

import "github.com/Tes-sudo/online-learning-platform/user-service/models"

func CreateUser(user *models.UserModel) error {
	return DB.Create(user).Error
}

func GetUserByID(id uint) (*models.UserModel, error) {
	var user models.UserModel
	err := DB.First(&user, id).Error
	return &user, err
}

func GetUserByEmail(email string)(*models.UserModel,error){
	var user models.UserModel
	err := DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func UpdateUser(user *models.UserModel) error {
	return DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return DB.Delete(&models.UserModel{}, id).Error
}