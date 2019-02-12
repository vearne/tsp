package main

import (
	"fmt"
	"math"
	"github.com/vearne/tsp"
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

var m1 map[string]Point

func init() {
	m1 = map[string]Point{}
	m1["A"] = Point{1, 1}
	m1["B"] = Point{2, 1}
	m1["C"] = Point{3, 1}
	m1["D"] = Point{3, 0}
	m1["E"] = Point{2, 0}
	m1["F"] = Point{1, 0}
}

func distance1(a, b string) float64 {
	posMap := m1
	return math.Sqrt(math.Pow(posMap[a].X-posMap[b].X, 2) + math.Pow(posMap[a].Y-posMap[b].Y, 2))
}

func main() {
	positions := []string{"A", "B", "C", "D", "E", "F"}
	maxIterCount := tsp.MaxIterationCountOption(5)
	t := tsp.NewTSP(positions, distance1, maxIterCount)
	fmt.Println(t.Slove())
}
