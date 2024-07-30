package dto

type StatisticRequest struct {
	Max               float64 `json:"max" validate:"required"`
	Min               float64 `json:"min" validate:"required"`
	Avg               float64 `json:"avg" validate:"required"`
	Sum               float64 `json:"sum" validate:"required"`
	IsQMatrixDiagonal *bool   `json:"isQMatrixDiagonal" validate:"required"`
	IsRMatrixDiagonal *bool   `json:"isRMatrixDiagonal" validate:"required"`
	FactorizationID   uint    `json:"factorization_id" validate:"required"`
}
