package tsp

var Memorandum map[int]int

func init() {
	Memorandum = make(map[int]int, 10)
}

func Factorial(n int) int {
	if value, ok := Memorandum[n]; ok {
		return value
	}

	if n == 0 {
		return 1
	} else {
		value := n * Factorial(n-1)
		Memorandum[n] = value
		return value
	}
}
