package services

import (
	"github.com/nelsonstos/isgchallenge/dto"
	"github.com/nelsonstos/isgchallenge/models"
	"github.com/nelsonstos/isgchallenge/repositories"
)

func SaveStatistic(s dto.StatisticRequest) (*models.Statistic, error) {
	statistic := &models.Statistic{
		CombinedStatistic: models.CombinedStatistic{
			Max: s.Max,
			Min: s.Min,
			Sum: s.Sum,
			Avg: s.Avg,
		},
		IsQMatrixDiagonal: *s.IsQMatrixDiagonal,
		IsRMatrixDiagonal: *s.IsRMatrixDiagonal,
		FactorizationID:   s.FactorizationID,
	}

	err := repositories.CreateStatistic(statistic)
	if err != nil {
		return nil, err
	}
	return statistic, nil
}
