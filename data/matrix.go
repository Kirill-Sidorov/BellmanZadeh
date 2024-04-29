package data

import (
	"fmt"
	"strings"
)

type Matrix struct {
	nubmerRows    int
	numberColumns int
	matrix        [][]float32
}

func NewMatrix(matrix [][]float32) *Matrix {
	return &Matrix{
		len(matrix),
		len(matrix[0]),
		matrix,
	}
}

func (m *Matrix) Get(i, j int) float32 {
	return m.matrix[i][j]
}

func (m *Matrix) ToTextAsTable(rowNames, columnNames []string) string {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("%10s", ""))

	for _, value := range columnNames {
		result.WriteString(fmt.Sprintf("%10.10s", value))
	}
	result.WriteString("\n")
	
	for i, rowName := range rowNames {
		result.WriteString(fmt.Sprintf("%10.10s", rowName))
		for j := range columnNames {
			result.WriteString(fmt.Sprintf("%10.3f", m.matrix[i][j]))
		}
		result.WriteString("\n")
	}
	result.WriteString("\n")

	return result.String()
}