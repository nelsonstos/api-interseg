package repositories

import (
	"github.com/nelsonstos/isgchallenge/config"

	"github.com/nelsonstos/isgchallenge/models"
)

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}
