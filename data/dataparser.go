package data

import (
	"encoding/json"
)

type Comparison struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ComparisonVariant struct {
	Criterion   string       `json:"criterion"`
	Variant     string       `json:"variant"`
	Comparisons []Comparison `json:"comparisons"`
}

type ComparisonCriteria struct {
	Criterion   string       `json:"criterion"`
	Comparisons []Comparison `json:"comparisons"`
}

type BellmanZadehData struct {
	Variants           []string            `json:"variants"`
	Criteriea          []string            `json:"criteria"`
	ComparisonVariants []ComparisonVariant `json:"comparison_variants"`
	ComparisonCriteria ComparisonCriteria  `json:"comparison_criteria"`
}

func ParseJsonData(data string) (*BellmanZadehData, error) {

	var result BellmanZadehData
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
