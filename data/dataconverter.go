package data

import "fmt"

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

func (c *Converter) CreateComparisonMatrix() (*Matrix, error) {

	return nil, fmt.Errorf("error")
}