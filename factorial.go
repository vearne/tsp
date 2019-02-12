package tsp

var Memorandum map[int64]int64

func init() {
	Memorandum = make(map[int64]int64, 10)
}

func Factorial(n int64) int64 {
	if value, ok := Memorandum[n]; ok {
		return value
	}

	if n <= 0 {
		return 1
	} else {
		value := n * Factorial(n-1)
		Memorandum[n] = value
		if value < 0 {
			panic("Data overflow int64")
		}
		return value
	}
}
