package handler

import (
	"bellmanzadeh/customrender"
	"bellmanzadeh/data"
	"bellmanzadeh/logic"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var mainPageHtml *template.Template

func Init() {
	var err error
	mainPageHtml, err = template.ParseFiles("resources/mainpage.html")
	if err != nil {
		log.Fatal(err)
	}
}

type mainPageData struct {
	ErrorMessage string
	Json         string
}

func MainPageHandler(response http.ResponseWriter, _ *http.Request) {

	err := mainPageHtml.Execute(response, &mainPageData{})
	if err != nil {
		log.Println(err)
	}
}

type resultPageData struct {
	ResultStringData string
	*components.Page
}

func ResultHandler(response http.ResponseWriter, request *http.Request) {

	jsonData := request.FormValue("jsonData")

	bellmanZadehData, err := data.ParseJsonData(jsonData)
	if err != nil {
		createMainPageWithError(response, err, jsonData)
		return
	}

	result, err := logic.SolveTask(bellmanZadehData)
	if err != nil {
		createMainPageWithError(response, err, jsonData)
		return
	}
	stringResult := createStringViewResult(result)

	criterieaNumbers := make([]int, 0, len(result.OrderCriteria))
	var criterieaSB strings.Builder
	for i, value := range result.OrderCriteria {
		num := i + 1
		criterieaSB.WriteString(fmt.Sprintf("%s - %d\n", value, num))
		criterieaNumbers = append(criterieaNumbers, num)
	}
	criterieaString := criterieaSB.String()

	equilibriumCriteriaLine, notEquilibriumCriteriaLine := createChartLinesData(result, criterieaNumbers, criterieaString)

	page := components.NewPage()
	page.Renderer = customrender.NewCustomRender(&resultPageData{stringResult, page}, page.Validate)
	page.AddCharts(equilibriumCriteriaLine, notEquilibriumCriteriaLine)

	err = page.Render(response)

	if err != nil {
		log.Println(err)
	}
}

func createMainPageWithError(response http.ResponseWriter, errText error, jsonData string) {
	err := mainPageHtml.Execute(response, &mainPageData{errText.Error(), jsonData})
	if err != nil {
		log.Println(err)
	}
}

func createStringViewResult(result *logic.BellmanZadehResult) string {

	var sb strings.Builder
	for key, value := range result.ComparisonVariants {
		sb.WriteString(fmt.Sprintf("\n\n\n%s\n", key))
		sb.WriteString(value.ToTextAsTable(result.OrderVariants, result.OrderVariants))
	}
	sb.WriteString("\n\n\nMatrix of paired comparisons of criteria:\n\n")
	sb.WriteString(result.ComparisonCriteria.ToTextAsTable(result.OrderCriteria, result.OrderCriteria))

	sb.WriteString("\n\n\nRanks of the matrix of paired comparisons of criteria:\n\n")
	for _, orderCriterion := range result.OrderCriteria {
		sb.WriteString(fmt.Sprintf("%25s", orderCriterion))
	}
	sb.WriteString("\n")
	for _, rank := range result.ComparisonCriteriaRanks {
		sb.WriteString(fmt.Sprintf("%25.3f", rank))
	}

	sb.WriteString("\n\n\n\nEquilibrium Criteria Matrix:\n\n")
	sb.WriteString(result.EquilibriumCriteriaMatrix.ToTextAsTable(
		result.OrderCriteria, result.OrderVariants))

	sb.WriteString(fmt.Sprintf("%20s", "Min"))
	for _, min := range result.MinEquilibriumCriteria {
		sb.WriteString(fmt.Sprintf("%20.3f", min))
	}

	sb.WriteString("\n\n\n\nMatrix of nonequilibrium criteria:\n\n")
	sb.WriteString(result.NotEquilibriumCriteriaMatrix.ToTextAsTable(
		result.OrderCriteria, result.OrderVariants))

	sb.WriteString(fmt.Sprintf("%20s", "Min"))
	for _, min := range result.MinNotEquilibriumCriteria {
		sb.WriteString(fmt.Sprintf("%20.3f", min))
	}

	return sb.String()
}

func createChartLinesData(result *logic.BellmanZadehResult, criterieaNumbers []int, criterieaString string) (equilibriumCriteriaLine, notEquilibriumCriteriaLine *charts.Line) {
	equilibriumCriteriaLine = charts.NewLine()
	equilibriumCriteriaLine.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Equilibrium criteria",
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

	notEquilibriumCriteriaLine = charts.NewLine()
	notEquilibriumCriteriaLine.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Nonequilibrium criteria",
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

	return equilibriumCriteriaLine, notEquilibriumCriteriaLine
}
