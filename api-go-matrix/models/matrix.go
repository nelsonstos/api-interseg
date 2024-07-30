package models

import (
	"time"
)

type Matrix struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Matrix    string    `gorm:"type:jsonb;not null" json:"matrix"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// TableName establece el nombre de la tabla para GORM
func (Matrix) TableName() string {
	return "matrixs"
}

type Factorization struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	MatrixID  uint      `gorm:"not null;index" json:"matrix_id"`
	QMatrix   string    `gorm:"type:jsonb;not null" json:"q_matrix"`
	RMatrix   string    `gorm:"type:jsonb;not null" json:"r_matrix"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// TableName establece el nombre de la tabla para GORM
func (Factorization) TableName() string {
	return "factorizations"
}
