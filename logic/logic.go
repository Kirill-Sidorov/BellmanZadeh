package logic

import (
	"bellmanzadeh/data"
	"fmt"
	"math"
)

type BellmanZadehLogicData struct {
	OrderVariants      []string
	OrderCriteria      []string
	ComparisonVariants map[string]*Matrix
	ComparisonCriteria *Matrix
}

type BellmanZadehResult struct {
	ComparisonCriteriaRanks      []float64
	EquilibriumCriteriaMatrix    *Matrix
	MinEquilibriumCriteria       []float64
	NotEquilibriumCriteriaMatrix *Matrix
	MinNotEquilibriumCriteria    []float64
	*BellmanZadehLogicData
}

func SolveTask(bellmanZadehData *data.BellmanZadehData) (*BellmanZadehResult, error) {

	logicData, err := prepareData(bellmanZadehData)
	if err != nil {
		return nil, err
	}

	resultData := findBellmanZadehResult(logicData)

	return resultData, nil
}

func prepareData(bellmanZadehData *data.BellmanZadehData) (*BellmanZadehLogicData, error) {

	orderVariants := bellmanZadehData.Variants
	comparisonVariants := make(map[string]*Matrix)

	for _, value := range bellmanZadehData.ComparisonVariants {
		matrix, err := createComparisonMatrix(orderVariants, value.Variant, value.Comparisons)
		if err != nil {
			return nil, err
		}
		comparisonVariants[value.Criterion] = matrix
	}

	orderCriteria := bellmanZadehData.Criteriea
	for _, value := range orderCriteria {
		if _, ok := comparisonVariants[value]; !ok {
			return nil, fmt.Errorf("no comparisons found for criterion - " + value)
		}
	}

	comparisonCriteria, err := createComparisonMatrix(
		orderCriteria,
		bellmanZadehData.ComparisonCriteria.Criterion,
		bellmanZadehData.ComparisonCriteria.Comparisons)

	if err != nil {
		return nil, fmt.Errorf("no comparisons found for criterion - " + err.Error())
	}

	return &BellmanZadehLogicData{
		orderVariants,
		orderCriteria,
		comparisonVariants,
		comparisonCriteria,
	}, nil
}

func findBellmanZadehResult(logicData *BellmanZadehLogicData) *BellmanZadehResult {

	// ranks of the pairwise comparison matrix
	comparisonCriteriaRanks := make([]float64, logicData.ComparisonCriteria.GetNubmerColumns())
	for i := 0; i < logicData.ComparisonCriteria.GetNubmerColumns(); i++ {
		var sum float64
		for j := 0; j < logicData.ComparisonCriteria.GetNumberRows(); j++ {
			sum += logicData.ComparisonCriteria.Get(j, i)
		}
		comparisonCriteriaRanks[i] = 1 / sum
	}

	// equilibrium criteria
	minEquilibriumCriteria := make([]float64, len(logicData.OrderVariants))
	for i := 0; i < len(logicData.OrderVariants); i++ {
		minEquilibriumCriteria[i] = 1
	}

	equilibriumCriteria := make([][]float64, 0, len(logicData.OrderCriteria))
	for _, criterion := range logicData.OrderCriteria {
		variantMatrix := logicData.ComparisonVariants[criterion]
		rowEquilibriumCriterion := make([]float64, 0, variantMatrix.GetNubmerColumns())

		for j := 0; j < variantMatrix.GetNubmerColumns(); j++ {

			var sum float64
			for i := 0; i < variantMatrix.GetNumberRows(); i++ {
				sum += variantMatrix.Get(i, j)
			}

			value := 1 / sum
			if value < minEquilibriumCriteria[j] {
				minEquilibriumCriteria[j] = value
			}
			rowEquilibriumCriterion = append(rowEquilibriumCriterion, value)
		}
		equilibriumCriteria = append(equilibriumCriteria, rowEquilibriumCriterion)
	}

	equilibriumCriteriaMatrix := NewMatrix(equilibriumCriteria)

	// nonequilibrium criteria
	minNotEquilibriumCriteria := make([]float64, len(logicData.OrderVariants))
	for i := 0; i < len(logicData.OrderVariants); i++ {
		minNotEquilibriumCriteria[i] = 1
	}

	notEquilibriumCriteria := make([][]float64, 0, equilibriumCriteriaMatrix.GetNumberRows())
	for i := 0; i < equilibriumCriteriaMatrix.GetNumberRows(); i++ {

		notEquilibriumCriteriaRow := make([]float64, 0, equilibriumCriteriaMatrix.GetNubmerColumns())
		for j := 0; j < equilibriumCriteriaMatrix.GetNubmerColumns(); j++ {

			value := math.Pow(equilibriumCriteriaMatrix.Get(i, j), comparisonCriteriaRanks[i])
			if value < minNotEquilibriumCriteria[j] {
				minNotEquilibriumCriteria[j] = value
			}
			notEquilibriumCriteriaRow = append(notEquilibriumCriteriaRow, value)
		}
		notEquilibriumCriteria = append(notEquilibriumCriteria, notEquilibriumCriteriaRow)
	}

	notEquilibriumCriteriaMatrix := NewMatrix(notEquilibriumCriteria)

	return &BellmanZadehResult{
		comparisonCriteriaRanks,
		equilibriumCriteriaMatrix,
		minEquilibriumCriteria,
		notEquilibriumCriteriaMatrix,
		minNotEquilibriumCriteria,
		logicData,
	}
}
