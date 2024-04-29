package data

import (
	"encoding/json"
	"io"
	"log"
	"os"
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

func ParseJsonData() *BellmanZadehData {
	filename, err := os.Open("resources/BellmanZadehData.json")
	if err != nil {
		log.Fatal(err)
	}
	defer filename.Close()

	data, err := io.ReadAll(filename)
	if err != nil {
		log.Fatal(err)
	}

	var result BellmanZadehData
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}
