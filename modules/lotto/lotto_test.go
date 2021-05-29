package lotto

import "testing"

func TestPrintLottoSet(t *testing.T) {
	r := PrintLottoSet() // Sum 함수를 테스트
	if r != true {
		t.Errorf("test fail")
	}

}

func TestLotsNumber(t *testing.T) {
	tests := []struct {
		value int
	}{
		// TODO: Add test cases.
		{1000},
		{1},
		{-1406},
		{1000000},
		{93482374},
	}
	for _, tt := range tests {
		r := LotsNumber(tt.value) // Sum 함수를 테스트
		if r < 0 || r > 45 {
			t.Errorf("test fail: %d", r)
		}
	}
}
