package repositories

import (
	"github.com/nelsonstos/isgchallenge/config"
	"github.com/nelsonstos/isgchallenge/models"
)

func CreateStatistic(statistic *models.Statistic) error {
	return config.DB.Create(statistic).Error
}
