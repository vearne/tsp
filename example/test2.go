package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"tsp"
)

type Point struct {
	X float64
	Y float64
}

var posMap2 map[string]Point

func init() {
	posMap2 = map[string]Point{}

	fi, err := os.Open("/tmp/data.csv")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		panic(err)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		buf, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		line := string(buf)
		line = strings.TrimSpace(line)
		itemList := strings.Split(line, ";")
		// 北京、上海
		pos := itemList[0]
		p := Point{}
		// 经度
		p.X, err = strconv.ParseFloat(itemList[1], 64)
		// 维度
		p.Y, _ = strconv.ParseFloat(itemList[2], 64)
		posMap2[pos] = p
	}
	fmt.Println("positions:", posMap2, "size:", len(posMap2))
}

func distance2(a, b string) float64 {
	posMap := posMap2
	//fmt.Println("ok", posMap[a], len(posMap))
	value := math.Sqrt(math.Pow(posMap[a].X-posMap[b].X, 2) + math.Pow(posMap[a].Y-posMap[b].Y, 2))
	//fmt.Printf("distance: %v and %v, value:%v\n", a, b, value)
	return value
}

func main() {
	// 数据量如果过大，会溢出产生负数
	//positions := []string{"北京", "天津", "上海", "重庆", "拉萨", "乌鲁木齐",
	//	"银川", "呼和浩特", "南宁", "哈尔滨", "长春", "沈阳", "石家庄",
	//	"太原", "西宁", "济南", "郑州", "南京", "合肥", "杭州", "福州", "南昌",
	//	"长沙", "武汉", "广州", "台北", "海口", "兰州", "西安", "成都", "贵阳",
	//	"昆明", "香港", "澳门"}
	positions := []string{"北京", "天津", "上海", "重庆", "拉萨", "乌鲁木齐",
		"银川", "呼和浩特", "南宁", "哈尔滨", "长春", "沈阳", "石家庄",
		"太原", "西宁"}
	maxIterCount := tsp.MaxIterationCountOption(200)
	maxPopulationSize := tsp.MaxPopulationSizeOption(1000)
	selectFunc := tsp.SelectAlgOption(tsp.RandomTraverseSampleSelection)
	//selectFunc := tsp.SelectAlgOption(tsp.RouletteWheelSelection)
	//selectFunc := tsp.SelectAlgOption(tsp.TournamentSelection)

	t := tsp.NewTSP(positions, distance2, maxIterCount, maxPopulationSize, selectFunc)
	result := t.Slove()
	fmt.Println("---result---:", result, "route distance", t.RouteDistance(result))
	// 输入到文件，以便绘图观察
	write2File(result, "/tmp/1.csv")

}

func write2File(seq []string, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	fmt.Fprintln(w, "posX,posY")

	for i := 0; i < len(seq); i++ {
		point := posMap2[seq[i]]
		lineStr := fmt.Sprintf("%v,%v", point.X, point.Y)
		fmt.Fprintln(w, lineStr)
	}
	return w.Flush()
}
