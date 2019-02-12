package tsp

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
	假定有 A、B、C、D 四个地点，对于旅行商问题而言，出发点和终止点都是A
	走法只能有以下6种情况
	A BCD A
	A BDC A
	A CBD A
	A CDB A
	A DBC A
	A DCB A

	使用遗传算法解决旅行商问题，只需要考虑B、C、D排列组合所构成的基因(二进制串)
    1. 交叉技术 采用单点杂交

	2. 变异    点突变(其中1位由 0 -> 1 或者 1 -> 0)
    3. 选择规则
	3.1 交叉和变异产生的 不合法的个体(二进制串) 直接淘汰
	3.2 种群个体数量没有达到最大个体数量限制前，所有的父母个体，直接加入下一代，超过数量后
	3.3 超过种群最大个体数量限制时，开始选择，默认采用 档次轮盘法
	4. 种群最大个体数量为10,000

*/

type jobOption struct {
	positions          []string
	distance           DistanceFunc
	initPopulationSize int
	// 种群的的最大个体数量不能设置的过小，种群的个体数量要达到一定的规模，
	// 才能保存基因的多样性，以及有足够的能力(速度)进行进化
	maxPopulationSize    int
	maxIterationCount    int
	selectAlg            string
	crossProbability     float64
	variationProbability float64
}

// JobOption specifies an option
type JobOption struct {
	f func(*jobOption)
}

// 初始时种群的数量
func InitPopulationSizeOption(populationSize int) JobOption {
	return JobOption{func(do *jobOption) {
		do.initPopulationSize = populationSize
	}}

}

// 设置种群个体的最大数量
func MaxPopulationSizeOption(populationSize int) JobOption {
	return JobOption{func(do *jobOption) {
		do.maxPopulationSize = populationSize
	}}
}

// 设置最大迭代次数
func MaxIterationCountOption(iterationCount int) JobOption {
	return JobOption{func(do *jobOption) {
		do.maxIterationCount = iterationCount
	}}
}

// 设置选择函数
func SelectFuncOption(f string) JobOption {
	return JobOption{func(do *jobOption) {
		do.selectAlg = f
	}}
}

// 交叉概率
func CrossProbability(p float64) JobOption {
	return JobOption{func(do *jobOption) {
		do.crossProbability = p
	}}
}

// 变异概率
func VariationProbability(p float64) JobOption {
	return JobOption{func(do *jobOption) {
		do.variationProbability = p
	}}
}

type DistanceFunc func(a, b string) float64

type TSP struct {
	option *jobOption
	// 基因(二进制串)对应数值的最大取值(闭区间)
	maxValue int
	// 基因的bit数
	geneLength int
	// 剔除初始点剩下的位置
	combinPostions []string
	// 起始点
	startPoint string
}

// positions 多个地点 positions的第1个元素为起始点和终止点
// distance  2个点之间的距离
func NewTSP(positions []string, distance DistanceFunc, options ...JobOption) *TSP {
	res := &TSP{}
	option := &jobOption{positions: positions, distance: distance}
	// 设置默认值
	option.initPopulationSize = len(positions)
	option.maxPopulationSize = 1000
	option.maxIterationCount = 200
	option.crossProbability = 0.7
	option.variationProbability = 0.01
	option.selectAlg = RouletteWheelSelection

	res.option = option
	for _, option := range options {
		option.f(res.option)
	}

	//所有基因的取值都在这个范围内 [0, maxValue]
	res.maxValue = res.MaxValue() - 1
	res.geneLength = int(math.Ceil(math.Log2(float64(Factorial(len(positions) - 1)))))
	res.combinPostions = res.option.positions[1:]
	res.startPoint = res.option.positions[0]
	return res
}

// 除去起止点，其它位置能够构成的情况数量
// 0， MaxValue -1
func (t *TSP) MaxValue() int {
	return Factorial(len(t.option.positions)-1) - 1
}

// 返回迭代指定次数后得到的最优解
func (t *TSP) Evolution() (gene string) {
	// 1.初始化种群
	population := t.InitPopulation()
	// 进入迭代过程
	// TODO 是否多次迭代最优解都没有发生变化，就可以退出迭代
	for counter := 0; counter <= t.option.maxIterationCount; counter++ {
		// 打印每一代的情况
		t.PrintGen(counter, population)
		// 2. 评估适应度
		fitnessSlice := t.Fitness(population)
		// 3. 选择
		population = SelectionMap[t.option.selectAlg](population, fitnessSlice, t.option.maxPopulationSize)
		// 4. 交叉 5. 变异
		population = t.CrossAndVariation(population)

	}

	fitnessSlice := t.Fitness(population)
	bestIdx := 0
	bestFitness := 0.0

	for idx, fitness := range fitnessSlice {
		if fitness > bestFitness {
			bestFitness = fitness
			bestIdx = idx
		}
	}
	return population[bestIdx]
}

func (t *TSP) PrintGen(counter int, genes []string) {
	fmt.Println("--------------")
	fmt.Println("counter:", counter, "size:", len(genes))
	fmt.Println("gene:", genes)
	for _, gene := range genes {
		fmt.Print(Gene2Sequence(t.option.positions, gene))
	}
	fmt.Println()
}

func (t *TSP) Slove() []string {
	seq := Gene2Sequence(t.combinPostions, t.Evolution())
	res := make([]string, 0, 10)
	res = append(res, t.startPoint)
	for _, item := range seq {
		res = append(res, item)
	}
	res = append(res, t.startPoint)
	return res
}

// 直接随机产生
func (t *TSP) InitPopulation() []string {
	set := NewStringSet()
	for i := 0; i < t.option.initPopulationSize; i++ {
		x := rand.Intn(t.maxValue + 1)
		set.Add(Encode(x, t.geneLength))
	}

	return set.ToArray()
}

func (t *TSP) Fitness(itemSlice []string) []float64 {
	res := make([]float64, len(itemSlice))
	for idx, item := range itemSlice {
		val := 0.0
		seq := Gene2Sequence(t.combinPostions, item)
		// 计算旅行需要走过的路程
		for i := 0; i < len(seq); i++ {
			if i > 0 {
				val += t.option.distance(seq[i], seq[i-1])
			}
		}
		// 从起始点出发的开销
		val += t.option.distance(t.startPoint, seq[0])
		// 回到起始点的开销
		val += t.option.distance(seq[len(seq)-1], t.startPoint)
		// 距离越远适应度越差
		res[idx] = 1 / val
	}
	return res
}

/*
   1. 交叉 每次从种群中，随机挑出2个，根据交叉概率，决定是否交叉
   2. 变异 交叉产生的新个体，根据变异概率，决定是否变异
*/
func (t *TSP) CrossAndVariation(population []string) []string {
	var N int
	set := NewStringSet()
	set.AddAll(population)

	N = len(population)
	for i := 0; i < N/2+1; i++ {
		// 选择父亲和母亲
		father := rand.Intn(N)
		mother := rand.Intn(N)
		// 父亲和母亲相同的情况，忽略掉
		if father == mother {
			continue
		}
		if rand.Float64() <= t.option.crossProbability {
			nextGen := t.CrossAction(population[father], population[mother])
			for _, gene := range nextGen {
				if rand.Float64() <= t.option.variationProbability {
					set.AddAll(t.VariationAction(gene))
				} else {
					set.Add(gene)
				}
			}
			set.AddAll(nextGen)
		}
	}
	return set.ToArray()
}

/*
	使用单交叉点算法
*/
func (t *TSP) CrossAction(gene1, gene2 string) []string {
	res := make([]string, 0, 2)
	point := rand.Intn(t.geneLength - 1)
	child1 := gene1[0:point] + gene2[point:]
	child2 := gene2[0:point] + gene1[point:]
	if t.isValid(child1) {
		res = append(res, child1)
	}
	if t.isValid(child2) {
		res = append(res, child2)
	}
	return res
}

/*
	只选择1bit 发生逆转
	返回 0个或1个
*/
func (t *TSP) VariationAction(gene string) []string {
	res := make([]string, 0, 1)
	point := rand.Intn(t.geneLength)
	newBit := '0'
	if gene[point] == '0' {
		newBit = '1'
	} else {
		newBit = '0'
	}
	temp := gene[0:point] + string(newBit)
	if point < len(gene)-1 {
		temp += gene[point+1:]
	}

	if t.isValid(temp) {
		res = append(res, temp)
	}
	return res
}

// 不能超过取值范围
func (t *TSP) isValid(gene string) bool {
	if Decode(gene) <= t.maxValue {
		return true
	}
	return false
}
