package mathutil

import "testing"

func TestAdd(t *testing.T) {
	result := Add(4, 2)
	if result != 6 {
		t.Errorf("wrong Output: got %d , expected  %d", result, 6)
	}
}

func TestMul(t *testing.T) {
	result := Mul(2, 3)
	if result != 6 {
		t.Errorf("Wrong Output : Expected-%d Output -%d", 6, result)
	}
}

func TestDiv(t *testing.T) {
	result := Div(18, 3)
	if result != 6 {
		t.Errorf("Wrong Output : Expected-%d Output -%d", 6, result)
	}
}

func TestSub(t *testing.T) {
	result := Sub(10, 4)
	if result != 6 {
		t.Errorf("Wrong Output : Expected-%d Output -%d", 6, result)
	}
}
