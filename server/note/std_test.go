package note

import (
	"errors"
	"testing"
)

// Testing有分T(Test)跟B(Benchmark)，前者測Bug，後者測速度
func TestIsPositive(t *testing.T) {
	err := errors.New("Is Negative")
	if IsPositive(-1) {
		t.Log("OK")
	} else {
		t.Error(err)
	}

	if IsPositive(1) {
		t.Log("OK")
	} else {
		t.Error(err)
	}
}
