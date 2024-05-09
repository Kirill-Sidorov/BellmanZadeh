package data

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Converter struct {
	elementsOrder []string
	matrixSize    int
}

func NewConverter(elementsOrder []string) *Converter {
	return &Converter{
		elementsOrder,
		len(elementsOrder),
	}
}

func (c *Converter) CreateComparisonMatrix(variant string, comparisons []Comparison) (*Matrix, error) {

	if len(comparisons) != c.matrixSize-1 {
		return nil, fmt.Errorf("error: comparisons size not equals matrix size")
	}

	comparisonMap := make(map[string]string)

	for _, value := range comparisons {
		comparisonMap[value.Name] = value.Value
	}

	knownRowNumber := 0
	knownRow := make([]float64, 0, c.matrixSize)
	for i, value := range c.elementsOrder {

		if value == variant {
			knownRow = append(knownRow, 0)
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
					return nil, fmt.Errorf("error")
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

	if (len(knownRow) != c.matrixSize) && (len(comparisonMap) != 0) {
		return nil, fmt.Errorf("error")
	}

	matrix := make([][]float64, 0, c.matrixSize)
	for i := 0; i < c.matrixSize; i++ {
		if (knownRowNumber != i) {
			row := make([]float64, 0, c.matrixSize)
			for j := 0; j < c.matrixSize; j++ {
				var value float64
				if (i == j) {
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
