package note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testing有分T(Test)跟B(Benchmark)，前者測Bug，後者測速度
func TestIsPositive(t *testing.T) {
	// 測試負數
	assert.Equal(t, false, IsPositive(-1), "IsPositive(-1) should return false")

	// 測試正數
	assert.Equal(t, true, IsPositive(1), "IsPositive(1) should return true")
}
