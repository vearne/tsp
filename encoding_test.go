package tsp

import "testing"

func TestEncode(t *testing.T) {
	target := "00101"
	value := Encode(5, 5)

	if target == value {
		t.Logf("success, expect:%v, result:%v", target, value)
	} else {
		t.Errorf("error, expect:%v, result:%v", target, value)
	}
}

func TestDecode(t *testing.T) {
	var target int64 = 5
	value := Decode("00101")

	if target == value {
		t.Logf("success, expect:%v, result:%v", target, value)
	} else {
		t.Errorf("error, expect:%v, result:%v", target, value)
	}
}

func TestSequence2Gene1(t *testing.T) {
	initSlice := []string{"B", "C", "D"}
	seq := []string{"C", "B", "D"}
	target := "010"
	value := Sequence2Gene(initSlice, seq)
	if target == value {
		t.Logf("success, expect:%v, result:%v", target, value)
	} else {
		t.Errorf("error, expect:%v, result:%v", target, value)
	}
}

func TestSequence2Gene2(t *testing.T) {
	initSlice := []string{"B", "C", "D"}
	seq := []string{"C", "D", "B"}
	target := "011"
	value := Sequence2Gene(initSlice, seq)
	if target == value {
		t.Logf("success, expect:%v, result:%v", target, value)
	} else {
		t.Errorf("error, expect:%v, result:%v", target, value)
	}
}

func TestSequence2Gene3(t *testing.T) {
	initSlice := []string{"B", "C", "D"}
	seq := []string{"B", "D", "C"}
	target := "001"
	value := Sequence2Gene(initSlice, seq)
	if target == value {
		t.Logf("success, expect:%v, result:%v", target, value)
	} else {
		t.Errorf("error, expect:%v, result:%v", target, value)
	}
}

func TestGene2Sequence1(t *testing.T) {
	initSlice := []string{"B", "C", "D"}
	gene := "001"
	seq := Gene2Sequence(initSlice, gene)
	target := []string{"B", "D", "C"}
	if StrSliceEqual(target, seq) {
		t.Logf("success, expect:%v, result:%v", target, seq)
	} else {
		t.Errorf("error, expect:%v, result:%v", target, seq)
	}
}

func TestGene2Sequence2(t *testing.T) {
	initSlice := []string{"B", "C", "D"}
	gene := "010"
	seq := Gene2Sequence(initSlice, gene)
	target := []string{"C", "B", "D"}
	if StrSliceEqual(target, seq) {
		t.Logf("success, expect:%v, result:%v", target, seq)
	} else {
		t.Errorf("error, expect:%v, result:%v", target, seq)
	}
}

func TestGene2Sequence3(t *testing.T) {
	initSlice := []string{"B", "C", "D"}
	gene := "100"
	seq := Gene2Sequence(initSlice, gene)
	target := []string{"D", "B", "C"}
	if StrSliceEqual(target, seq) {
		t.Logf("success, expect:%v, result:%v", target, seq)
	} else {
		t.Errorf("error, expect:%v, result:%v", target, seq)
	}
}

func TestGene2Sequence4(t *testing.T) {
	initSlice := []string{"B", "C", "D"}
	gene := "101"
	seq := Gene2Sequence(initSlice, gene)
	target := []string{"D", "C", "B"}
	if StrSliceEqual(target, seq) {
		t.Logf("success, expect:%v, result:%v", target, seq)
	} else {
		t.Errorf("error, expect:%v, result:%v", target, seq)
	}
}

func StrSliceEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
