package tsp

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// length 生成的2进制串的长度， 不足的部分在高位补0
func Encode(value int, length int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(length)+"b", value)
}

func Decode(s string) int {
	value, _ := strconv.ParseInt(s, 2, 64)
	return int(value)
}

func Sequence2Gene(initSilce []string, seq []string) string {

	// 1. seq route B-C-D
	// 2. 项对应的编号
	set := NewStringSet()
	for _, item := range initSilce {
		set.Add(item)
	}

	indexSlice := make([]int, len(seq))
	for idx, item := range seq {
		tempSlice := set.ToArray()
		sort.StringSlice(tempSlice).Sort()
		for i := 0; i < len(initSilce); i++ {
			if tempSlice[i] == item {
				indexSlice[idx] = i
				break
			}
		}
		set.Remove(item)
	}

	// 3. 转换成数值
	//fmt.Printf("value:%v\n", indexSlice)
	value := 0
	for idx, num := range indexSlice {
		value += num * Factorial(len(indexSlice)-idx-1)
	}

	//fmt.Printf("value:%v\n", value)
	// 4. 换成2进制(形如: "010")
	N := len(initSilce)
	length := int(math.Ceil(math.Log2(float64(Factorial(N)))))
	// 计算基因序列的最大长度
	return Encode(value, length)
}

func Gene2Sequence(initSilce []string, gene string) (seq []string) {
	set := NewStringSet()
	for _, item := range initSilce {
		set.Add(item)
	}
	// 1. gene 2进制(形如: "010")
	// 2. 转换为数值
	value := Decode(gene)
	// 3. 转换为项对应的编号 & 生成路径(序列)
	N := len(initSilce)
	res := make([]string, N)
	for i := 0; i < len(initSilce); i++ {
		tempSlice := set.ToArray()
		sort.StringSlice(tempSlice).Sort()
		f := Factorial(N - i - 1)

		num := value / f
		//fmt.Printf("value:%v, f:%v, num:%v\n", value, f, num)
		value = value % f
		res[i] = tempSlice[num]

		set.Remove(tempSlice[num])
	}
	return res
}
