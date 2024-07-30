package dto

type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix" validate:"required"`
}

type MatrixResponse struct {
	Q [][]float64 `json:"Q"`
	R [][]float64 `json:"R"`
}
