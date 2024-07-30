package repositories

import (
	"github.com/nelsonstos/isgchallenge/config"
	"github.com/nelsonstos/isgchallenge/models"
)

func CreateMatrix(matrix *models.Matrix) error {
	return config.DB.Create(matrix).Error
}

func CreateFactorizationMatrix(matrix *models.Factorization) error {
	return config.DB.Create(matrix).Error
}
