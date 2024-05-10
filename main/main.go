package main

import (
	"bellmanzadeh/data"
	"fmt"
	"strings"
)

func main() {

	bellmanZadehData := data.ParseJsonData()

	orderVariants := bellmanZadehData.Variants
	comparisonVariants := make(map[string]*data.Matrix)

	matrixCreation := data.NewConverter(orderVariants)

	for _, value := range bellmanZadehData.ComparisonVariants {
		matrix, err := matrixCreation.CreateComparisonMatrix(value.Variant, value.Comparisons)
		if err != nil {
			panic(err)
		}
		comparisonVariants[value.Criterion] = matrix
	}

	orderCriteria := bellmanZadehData.Criteriea
	for _, value := range orderCriteria {
		if _, ok := comparisonVariants[value]; !ok {
			panic("No comparisons found for criterion - " + value)
		}
	}

	matrixCreation = data.NewConverter(orderCriteria)
	matrix, err := matrixCreation.CreateComparisonMatrix(
		bellmanZadehData.ComparisonCriteria.Criterion,
		bellmanZadehData.ComparisonCriteria.Comparisons)

	if (err != nil) {
		panic("Failed to create criterion pairwise comparison matrix " + err.Error())
	}	

	var result strings.Builder
	for key, value := range comparisonVariants {
		result.WriteString(fmt.Sprintf("%s\n", key))
		result.WriteString(value.ToTextAsTable(orderVariants, orderVariants))
	}
	result.WriteString("Matrix of paired comparisons of criteria:\n\n")
	result.WriteString(matrix.ToTextAsTable(orderCriteria, orderCriteria))
	
	fmt.Println(result.String())

	// solve part

	// ranks of the pairwise comparison matrix
	comparisonCriteriaRanks := make([]float64, 0, matrix.GetNubmerColumns())
	for i := 0; i < matrix.GetNubmerColumns(); i++ {
		var sum float64
		for j := 0; j < matrix.GetNumberRows(); j++ {
			sum += matrix.Get(j, i)
		}
		comparisonCriteriaRanks[i] = 1 / sum
	}

	// equilibrium criteria
	minEquilibriumCriteria := make([]float64, 0, len(orderVariants))
	for i := 0; i < len(orderVariants); i++ {
		minEquilibriumCriteria[i] = 1
	}

	//equilibriumCriteria := make([][]float64, 0, len(orderCriteria))
	


	/*
			server := http.NewServeMux()
			server.HandleFunc("/", baseHandler)
			server.HandleFunc("/favicon.ico", faviconHandler)

			log.Println("Server Started...")
			err := http.ListenAndServe("localhost:8080", server)
			if err != nil {
				log.Fatal(err)
			}

			func baseHandler(writer http.ResponseWriter, request *http.Request) {

		}

		func faviconHandler(writer http.ResponseWriter, request *http.Request) {
			http.ServeFile(writer, request, "resources/favicon.ico")
		}
	*/
}
