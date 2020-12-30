package util

import "testing"

func TestMinInt(t *testing.T) {
	paramIntOne := 6
	paramIntTwo := 5
	want := 5
	if got := MinInt(paramIntOne, paramIntTwo); got != want {
		t.Errorf("MinInt(%d, %d) = %d, want %d", paramIntOne, paramIntTwo, got, want)
	}
}
