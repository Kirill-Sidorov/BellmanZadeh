package main

import (
	"bellmanzadeh/customrender"
	"bellmanzadeh/data"
	"bellmanzadeh/logic"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/go-echarts/go-echarts/v2/types"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", baseHandler)
	server.HandleFunc("/favicon.ico", faviconHandler)

	log.Println("Server Started...")
	err := http.ListenAndServe("localhost:8080", server)
	if err != nil {
		log.Fatal(err)
	}
}

type mainPageData struct {
	ResultStringData string
	*charts.Line
}

func baseHandler(writer http.ResponseWriter, request *http.Request) {
	bellmanZadehData := data.ParseJsonData()
	result, err := logic.SolveTask(bellmanZadehData)
	if err != nil {
		panic(err)
	}
	stringResult := createStringResult(result)

	line := charts.NewLine()
	line.Renderer = customrender.NewCustomRender(&mainPageData{stringResult, line}, line.Validate)
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Равновесные критерии",
		}))

	line = line.SetXAxis(result.OrderCriteria)

	for j, variant := range result.OrderVariants {
		matrix := result.EquilibriumCriteriaMatrix
		items := make([]opts.LineData, 0, len(result.OrderCriteria))
		for i := range result.OrderCriteria {
			items = append(items, opts.LineData{Value: matrix.Get(i, j)})
		}
		line = line.AddSeries(variant, items)
	}

	err = line.Render(writer)

	if err != nil {
		log.Println(err)
	}
}

func faviconHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "resources/favicon.ico")
}

func createStringResult(result *logic.BellmanZadehResult) string {

	var sb strings.Builder
	for key, value := range result.ComparisonVariants {
		sb.WriteString(fmt.Sprintf("%s\n", key))
		sb.WriteString(value.ToTextAsTable(result.OrderVariants, result.OrderVariants))
	}
	sb.WriteString("Matrix of paired comparisons of criteria:\n\n")
	sb.WriteString(result.ComparisonCriteria.ToTextAsTable(result.OrderCriteria, result.OrderCriteria))

	sb.WriteString("Ранги матрицы парных сравений критериев\n")
	for _, orderCriterion := range result.OrderCriteria {
		sb.WriteString(fmt.Sprintf("%15s", orderCriterion))
	}

	sb.WriteString("\n")
	for _, rank := range result.ComparisonCriteriaRanks {
		sb.WriteString(fmt.Sprintf("%15.3f", rank))
	}

	sb.WriteString("\n\nМатрица равновесных критериев:\n")
	sb.WriteString(result.EquilibriumCriteriaMatrix.ToTextAsTable(
		result.OrderCriteria, result.OrderVariants))

	sb.WriteString(fmt.Sprintf("%15s", "Мин"))
	for _, min := range result.MinEquilibriumCriteria {
		sb.WriteString(fmt.Sprintf("%15.3f", min))
	}

	sb.WriteString("\n\nМатрица неравновесных критериев:\n")
	sb.WriteString(result.NotEquilibriumCriteriaMatrix.ToTextAsTable(
		result.OrderCriteria, result.OrderVariants))

	sb.WriteString(fmt.Sprintf("%15s", "Мин"))
	for _, min := range result.MinNotEquilibriumCriteria {
		sb.WriteString(fmt.Sprintf("%15.3f", min))
	}

	return sb.String()
}

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}
