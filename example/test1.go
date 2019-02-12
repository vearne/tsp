package main

import (
	"fmt"
	"math"
	"tsp"
)

type Point struct {
	X float64
	Y float64
}

/*
		A--B--C
		|  |  |
        F--E--D

*/

var posMap1 map[string]Point

func init() {
	posMap1 = map[string]Point{}
	posMap1["A"] = Point{1, 1}
	posMap1["B"] = Point{2, 1}
	posMap1["C"] = Point{3, 1}
	posMap1["D"] = Point{3, 0}
	posMap1["E"] = Point{2, 0}
	posMap1["F"] = Point{1, 0}
}

func distance1(a, b string) float64 {
	posMap := posMap1
	return math.Sqrt(math.Pow(posMap[a].X-posMap[b].X, 2) + math.Pow(posMap[a].Y-posMap[b].Y, 2))
}

func main() {
	positions := []string{"A", "B", "C", "D", "E", "F"}
	maxIterCount := tsp.MaxIterationCountOption(5)
	maxPopulationSize := tsp.MaxPopulationSizeOption(100)
	t := tsp.NewTSP(positions, distance1, maxIterCount, maxPopulationSize)

	fmt.Println("----result-----:", t.Slove())
}
