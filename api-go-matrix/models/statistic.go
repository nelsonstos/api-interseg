package models

import "time"

type CombinedStatistic struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
	Avg float64 `json:"avg"`
	Sum float64 `json:"sum"`
}

type Statistic struct {
	ID                uint              `json:"id" gorm:"primaryKey"`
	CombinedStatistic CombinedStatistic `json:"combinedStatistic" gorm:"embedded"`
	IsQMatrixDiagonal bool              `json:"isQMatrixDiagonal"`
	IsRMatrixDiagonal bool              `json:"isRMatrixDiagonal"`
	FactorizationID   uint              `json:"factorization_id"`
	CreatedAt         time.Time         `gorm:"autoCreateTime" json:"created_at"`
}

func (Statistic) TableName() string {
	return "statistics"
}
