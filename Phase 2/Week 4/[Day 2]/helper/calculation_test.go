package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSumRequire(t *testing.T) {
	t.log("TestSumRequire is running")
	result := Sum(1, 2)
	require.Equal(t, 3, result, "The result should be 3")
	require.NotEqual(t, 4, result, "The result should not be 4")
	t.log("TestSumRequire is done")
}

func TestSumAssert(t *testing.T) {
	t.log("TestSumAssert is running")
	result := Sum(1, 2)
	assert.Equal(t, 3, result, "The result should be 3")
	assert.NotEqual(t, 4, result, "The result should not be 4")
	t.log("TestSumAssert is done")
}

// assert --> 3, 1 fail next .. 2, 3
// require --> 3, 1 fail, stop

func TestSum(t *testing.T) {
	tests := []struct {
		Expected int
		Actual   int
		ErrMsg   string
	}{
		{Expected: 3, Actual: Sum(1, 2), ErrMsg: "The result should be 3"},
		{Expected: 4, Actual: Sum(1, 2), ErrMsg: "The result should not be 4"},
		{Expected: 5, Actual: Sum(2, 3), ErrMsg: "The result should be 5"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Test Sum %d ", test.Actual), func(t *testing.T) {
			assert.Equal(t, test.Expected, test.Actual, test.ErrMsg)
		})
	}
}
