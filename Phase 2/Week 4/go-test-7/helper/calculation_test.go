package helper

// assert -> 3, 1 fail next.. 2, 3
// require -> 3, 1 fail, stop.

// func TestSumAssert(t *testing.T) {
// 	t.Log("Test with assert start")
// 	res := Sum(1, 2)
// 	assert.Equal(t, 4, res, "Result should be 3")
// 	assert.NotNil(t, res, "Result should not be nil")
// 	t.Log("Test with assert end")
// }

// func TestSumRequire(t *testing.T) {
// 	t.Log("Test with require start")
// 	res := Sum(1, 2)
// 	require.Equal(t, 4, res, "Result should be 3")
// 	require.NotNil(t, res, "Result should not be nil")
// 	t.Log("Test with require end")
// }

// func TestSum(t *testing.T) {
// 	tests := []struct {
// 		Expected int
// 		Actual   int
// 		ErrMsg   string
// 	}{
// 		{Actual: Sum(2, 2), Expected: 4, ErrMsg: "Invalid Sum Result"},
// 		{Actual: Sum(1, 3), Expected: 5, ErrMsg: "Invalid Sum Result"},
// 		{Actual: Sum(5, 12), Expected: 17, ErrMsg: "Invalid Sum Result"},
// 		{Actual: Sum(12, 20), Expected: 32, ErrMsg: "Invalid Sum Result"},
// 	}

// 	for _, test := range tests {
// 		t.Run(fmt.Sprintf("Test Sum %d : ", test.Actual), func(t *testing.T) {
// 			assert.Equal(t, test.Expected, test.Actual, test.ErrMsg)
// 		})
// 	}
// }
