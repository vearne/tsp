package tsp

import (
	"testing"
)

func fake(a, b string) float64 {
	return 0
}

func TestCrossAction(t *testing.T) {
	tsp := NewTSP([]string{"xxx"}, fake)
	tsp.geneLength = 8
	tsp.maxValue = Factorial(10)
	res := tsp.CrossAction("00101100", "11101010")
	if len(res) >= 0 {
		t.Logf("success, %v", res)
	} else {
		t.Errorf("error")
	}
}

func TestVariationAction(t *testing.T) {
	tsp := NewTSP([]string{"xxx"}, fake)
	tsp.geneLength = 8
	tsp.maxValue = Factorial(10)
	res := tsp.VariationAction("00101100")
	if len(res) >= 0 {
		t.Logf("success, %v", res)
	} else {
		t.Errorf("error")
	}
}
