package main

import (
	"bellmanzadeh/customrender"
	"bellmanzadeh/data"
	"bellmanzadeh/logic"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
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
	*components.Page
}

func baseHandler(writer http.ResponseWriter, request *http.Request) {
	bellmanZadehData := data.ParseJsonData()
	result, err := logic.SolveTask(bellmanZadehData)
	if err != nil {
		panic(err)
	}
	stringResult := createStringResult(result)

	criterieaNumbers := make([]int, 0, len(result.OrderCriteria))
	var criterieaSB strings.Builder
	for i, value := range result.OrderCriteria {
		num := i + 1
		criterieaSB.WriteString(fmt.Sprintf("%s - %d\n", value, num))
		criterieaNumbers = append(criterieaNumbers, num)
	}
	criterieaString := criterieaSB.String()

	equilibriumCriteriaLine := charts.NewLine()
	equilibriumCriteriaLine.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Равновесные критерии",
			Subtitle: criterieaString,
		}))

	equilibriumCriteriaLine = equilibriumCriteriaLine.SetXAxis(criterieaNumbers)

	for j, variant := range result.OrderVariants {
		matrix := result.EquilibriumCriteriaMatrix
		items := make([]opts.LineData, 0, len(result.OrderCriteria))
		for i := range result.OrderCriteria {
			items = append(items, opts.LineData{Value: matrix.Get(i, j)})
		}
		equilibriumCriteriaLine = equilibriumCriteriaLine.AddSeries(variant, items)
	}

	notEquilibriumCriteriaLine := charts.NewLine()
	notEquilibriumCriteriaLine.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "Неравновесные критерии",
			Subtitle: criterieaString,
		}))

	notEquilibriumCriteriaLine = notEquilibriumCriteriaLine.SetXAxis(criterieaNumbers)

	for j, variant := range result.OrderVariants {
		matrix := result.NotEquilibriumCriteriaMatrix
		items := make([]opts.LineData, 0, len(result.OrderCriteria))
		for i := range result.OrderCriteria {
			items = append(items, opts.LineData{Value: matrix.Get(i, j)})
		}
		notEquilibriumCriteriaLine = notEquilibriumCriteriaLine.AddSeries(variant, items)
	}

	page := components.NewPage()
	page.Renderer = customrender.NewCustomRender(&mainPageData{stringResult, page}, page.Validate)
	page.AddCharts(equilibriumCriteriaLine, notEquilibriumCriteriaLine)

	err = page.Render(writer)

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