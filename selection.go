package tsp

import (
	"math/rand"
)

const (
	RouletteWheelSelection = "Roulette_Wheel"
)

// 返回存活下来的个体
type SelectFunc func(genes []string, fitnessSlice []float64, maxSize int) []string

var SelectionMap map[string]SelectFunc

func init() {
	SelectionMap = make(map[string]SelectFunc)
	Register(RouletteWheelSelection, RouletteWheel)
}

func Register(key string, f SelectFunc) {
	SelectionMap[key] = f
}

/*
	种群的的最大个体数量不能设置的过小，种群的个体数量要达到一定的规模，
	才能保存基因的多样性，以及有足够的能力(速度)进行进化
*/
func RouletteWheel(genes []string, fitnessSlice []float64, maxSize int) []string {
	// 未超过最大限制，则全部保留
	if len(genes) <= maxSize {
		return genes
	}

	var sum float64 = 0
	var temp float64 = 0
	// 0-->5-->10----->20--------->35
	// 适应度大的生存的机率更大
	var upBounds []float64 = make([]float64, 0, len(genes))

	// 归一化
	for _, fitness := range fitnessSlice {
		sum += fitness
	}
	for _, fitness := range fitnessSlice {
		temp += fitness
		upBounds = append(upBounds, temp/sum)
	}

	//fmt.Println("genes:", len(genes), "fitnessSlice:",
	//len(fitnessSlice), "upBounds:", len(upBounds), "maxSize:", maxSize)
	//fmt.Println("fitnessSlice:", fitnessSlice)
	set := NewStringSet()
	for i := 0; i <= maxSize; i++ {
		for idx, upBound := range upBounds {
			if rand.Float64() < upBound {
				set.Add(genes[idx])
				break
			}
		}
	}

	return set.ToArray()
}
