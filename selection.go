package tsp

import (
	"fmt"
	"math/rand"
)

const (
	RouletteWheelSelection = "RouletteWheel"
	TournamentSelection    = "Tournament"
	//随机遍历抽样法
	RandomTraverseSampleSelection = "RandomTraverseSample"
)

// 返回存活下来的个体
type SelectFunc func(genes []string, fitnessSlice []float64, maxSize int) []string

var SelectionMap map[string]SelectFunc

func init() {
	SelectionMap = make(map[string]SelectFunc)
	Register(RouletteWheelSelection, RouletteWheel)
	Register(RandomTraverseSampleSelection, RandomTraverseSample)
	Register(TournamentSelection, Tournament)
}

func Register(key string, f SelectFunc) {
	SelectionMap[key] = f
}

func RandomTraverseSample(genes []string, fitnessSlice []float64, maxSize int) []string {
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

	var step float64 = 1.0 / float64(maxSize)
	var vister float64 = 0.0
	set := NewStringSet()
	vister = 0
	for i := 0; i <= maxSize; i++ {
		for idx, upBound := range upBounds {
			if vister < upBound {
				set.Add(genes[idx])
				break
			}
		}
		vister += step
		//fmt.Println("vister", vister)
	}
	fmt.Println("before selection", len(genes))
	fmt.Println("after selection", set.Size())
	return set.ToArray()

}

func Tournament(genes []string, fitnessSlice []float64, maxSize int) []string {
	// 未超过最大限制，则全部保留
	if len(genes) <= maxSize {
		return genes
	}

	count := len(genes)
	fmt.Println("before selection", len(genes))

	intSlice := make([]int, count)
	for i := 0; i < count; i++ {
		intSlice[i] = i
	}

	choiceCount := maxSize
	var idx1, idx2 int
	for i := 0; i < choiceCount; i++ {
		idx1 = rand.Int()%count + i
		idx2 = rand.Int()%count + i
		target := idx1

		if fitnessSlice[idx1] < fitnessSlice[idx2] {
			target = idx2
		} else {
			target = idx1
		}
		swap(intSlice, i, target)
		count--

	}

	// intSlice[0:choiceCount] 为锦标赛的胜出者
	res := make([]string, 0, choiceCount)
	for i := 0; i < choiceCount; i++ {
		res = append(res, genes[intSlice[i]])
	}

	fmt.Println("after selection", len(res))
	return res
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
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

	fmt.Println("before selection", len(genes))
	fmt.Println("after selection", set.Size())
	return set.ToArray()
}
