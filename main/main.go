package main

import (
	"bellmanzadeh/data"
	"fmt"
)

func main() {

	bellmanZadehData := data.ParseJsonData()

	orderVariants := bellmanZadehData.Variants
	//orderCriteria := bellmanZadehData.Criteriea
	//comparisonVariants := make(map[string]*data.Matrix)

	matrixCreation := data.NewConverter(orderVariants)

	for _, value := range bellmanZadehData.ComparisonVariants {
		matrix, err := matrixCreation.CreateComparisonMatrix(value.Variant, value.Comparisons)
		if (err != nil) {
			panic(err)
		} else {
			fmt.Println(matrix.ToText())
		}
	}



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
