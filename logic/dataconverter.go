package logic

import (
	"bellmanzadeh/data"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func createComparisonMatrix(elementsOrder []string, variant string, comparisons []data.Comparison) (*Matrix, error) {

	matrixSize := len(elementsOrder)

	if len(comparisons) != matrixSize-1 {
		return nil, fmt.Errorf("error: comparisons size not equals matrix size")
	}

	comparisonMap := make(map[string]string)

	for _, value := range comparisons {
		comparisonMap[value.Name] = value.Value
	}

	knownRowNumber := 0
	knownRow := make([]float64, 0, matrixSize)
	for i, value := range elementsOrder {

		if value == variant {
			knownRow = append(knownRow, 1)
			knownRowNumber = i
			continue
		}

		if stringValue, ok := comparisonMap[value]; ok {
			delete(comparisonMap, value)

			var doubleValue float64
			if strings.Contains(stringValue, "/") {

				fractions := strings.Split(stringValue, "/")

				if len(fractions) > 1 {
					numerator, err1 := strconv.ParseFloat(fractions[0], 64)
					denominator, err2 := strconv.ParseFloat(fractions[1], 64)

					if err1 != nil {
						return nil, err1
					}
					if err2 != nil {
						return nil, err2
					}

					doubleValue = roundFloat(numerator / denominator)
				} else {
					return nil, fmt.Errorf("error parse comparison value for %s: %s", value, stringValue)
				}

			} else {
				var err error
				doubleValue, err = strconv.ParseFloat(stringValue, 64)
				if err != nil {
					return nil, err
				}
				doubleValue = roundFloat(doubleValue)
			}
			knownRow = append(knownRow, doubleValue)
		}
	}

	if (len(knownRow) != matrixSize) && (len(comparisonMap) != 0) {
		return nil, fmt.Errorf("error: input data not valid")
	}

	matrix := make([][]float64, 0, matrixSize)
	for i := 0; i < matrixSize; i++ {
		if knownRowNumber != i {
			row := make([]float64, 0, matrixSize)
			for j := 0; j < matrixSize; j++ {
				var value float64
				if i == j {
					value = 1
				} else {
					value = knownRow[j] / knownRow[i]
				}
				row = append(row, value)
			}
			matrix = append(matrix, row)

		} else {
			matrix = append(matrix, knownRow)
		}
	}

	return NewMatrix(matrix), nil
}

func roundFloat(val float64) float64 {
	ratio := math.Pow(10, float64(3))
	return math.Round(val*ratio) / ratio
}
