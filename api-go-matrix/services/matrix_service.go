package services

import (
	"encoding/json"

	"github.com/nelsonstos/isgchallenge/dto"
	"github.com/nelsonstos/isgchallenge/models"
	"github.com/nelsonstos/isgchallenge/repositories"
	"gonum.org/v1/gonum/mat"
)

func CalculateQR(m dto.MatrixRequest) (*models.Factorization, error) {
	rows := len(m.Matrix)
	cols := len(m.Matrix[0])

	A := mat.NewDense(rows, cols, flatten(m.Matrix))

	// Factorizar la matriz
	qr := new(mat.QR)
	qr.Factorize(A)

	// Obtener las matrices Q y R
	var Q, R mat.Dense
	Q, R = *new(mat.Dense), *new(mat.Dense)
	qr.QTo(&Q)
	qr.RTo(&R)

	// Preparar la respuesta
	// Convertir Q y R a listas anidadas
	qList := toNestedList(&Q)
	rList := toNestedList(&R)

	//save matrix
	matrixJSON, _ := json.Marshal(m.Matrix)
	matrix := models.Matrix{Matrix: string(matrixJSON)}
	err := repositories.CreateMatrix(&matrix)
	if err != nil {
		return nil, err
	}

	//save factorization
	QMatrixJSON, _ := json.Marshal(qList)
	RMatrixJSON, _ := json.Marshal(rList)
	factorization := models.Factorization{
		MatrixID: matrix.ID,
		QMatrix:  string(QMatrixJSON),
		RMatrix:  string(RMatrixJSON),
	}

	errFact := repositories.CreateFactorizationMatrix(&factorization)
	if errFact != nil {
		return nil, errFact
	}

	return &factorization, nil
}

// flatten convierte una matriz 2D en un slice 1D
func flatten(matrix [][]float64) []float64 {
	var result []float64
	for _, row := range matrix {
		result = append(result, row...)
	}
	return result
}

// toNestedList convierte una matriz gonum en una lista anidada
func toNestedList(m mat.Matrix) [][]float64 {
	rows, cols := m.Dims()
	nestedList := make([][]float64, rows)
	for r := 0; r < rows; r++ {
		nestedList[r] = make([]float64, cols)
		for c := 0; c < cols; c++ {
			nestedList[r][c] = m.At(r, c)
		}
	}
	return nestedList
}
