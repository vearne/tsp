package tsp

import "testing"

func TestFactorial1(t *testing.T) {
	value := Factorial(15)
	var target int64 = 1307674368000

	if target == value {
		t.Logf("success, expect:%v, result:%v, %v", target, value, Memorandum)
	} else {
		t.Errorf("error, expect:%v, result:%v, %v", target, value, Memorandum)
	}
}

func TestFactorial2(t *testing.T) {
	value := Factorial(5)
	var target int64 = 120

	if target == value {
		t.Logf("success, expect:%v, result:%v, %v", target, value, Memorandum)
	} else {
		t.Errorf("error, expect:%v, result:%v, %v", target, value, Memorandum)
	}
}
