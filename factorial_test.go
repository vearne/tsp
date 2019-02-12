package tsp

import "testing"

func TestFactorial1(t *testing.T) {
	value := Factorial(5)
	target := 120

	if target == value {
		t.Logf("success, expect:%v, result:%v, %v", target, value, Memorandum)
	} else {
		t.Errorf("error, expect:%v, result:%v", target, value)
	}
}
