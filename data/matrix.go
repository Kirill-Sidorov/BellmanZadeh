package data

import (
	"fmt"
	"strings"
)

type Matrix struct {
	nubmerRows    int
	numberColumns int
	matrix        [][]float64
}

func NewMatrix(matrix [][]float64) *Matrix {
	return &Matrix{
		len(matrix),
		len(matrix[0]),
		matrix,
	}
}

func (m *Matrix) Get(i, j int) float64 {
	return m.matrix[i][j]
}

func (m *Matrix) GetNumberRows() int {
	return m.nubmerRows
}

func (m *Matrix) GetNubmerColumns() int {
	return m.numberColumns
}

func (m *Matrix) ToText() string {
	var result strings.Builder
	for _, row := range m.matrix {
		for _, value := range row {
			result.WriteString(fmt.Sprintf("%6.2f", value))
		}
		result.WriteString("\n")
	}
	return result.String()
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