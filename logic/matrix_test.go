package logic

import (
	"testing"
)

var testData = [][]float64{
	{0.1, 0.2, 0.3},
	{0.4, 0.5, 0.6},
}

var matrixTestData = NewMatrix(testData)

func TestNewMatrix(t *testing.T) {

	if (matrixTestData.nubmerRows != len(testData)) {
		t.Fatalf("Matrix number rows expect - %d, but actual - %d after create", len(testData), matrixTestData.nubmerRows)
	}
	if (matrixTestData.numberColumns != len(testData[0])) {
		t.Fatalf("Matrix number columns expect - %d, but actual - %d after create", len(testData[0]), matrixTestData.numberColumns)
	}
}

func TestGet(t *testing.T) {
	
	if (matrixTestData.Get(1, 1) != testData[1][1]) {
		t.Fatalf("Matrix item expect - %f, but actual - %f", testData[1][1], matrixTestData.Get(1, 1))
	}
}

func TestGetNumberRows(t *testing.T) {
	if (matrixTestData.GetNumberRows() != len(testData)) {
		t.Fatalf("Matrix number rows expect - %d, but actual - %d", len(testData), matrixTestData.GetNumberRows())
	}
}

func TestGetNumberColumns(t *testing.T) {
	if (matrixTestData.GetNubmerColumns() != len(testData[0])) {
		t.Fatalf("Matrix number columns expect - %d, but actual - %d", len(testData[0]), matrixTestData.GetNubmerColumns())
	}
}

func TestToTextAsTable(t *testing.T) {
	rowNames := []string{"Row1", "Row2"}
	columnNames := []string{"Column1", "Column2", "Column3"}

	actualResult := matrixTestData.ToTextAsTable(rowNames, columnNames)

	if (actualResult == "") {
		t.Fatal("Error: Matrix to text as table - empty result")
	}
}