package main

import (
	"sort"
)

// MEDIAN FUNCTION
func median(data []float64) float64 {
	dataCopy := make([]float64, len(data))
	copy(dataCopy, data)

	sort.Float64s(dataCopy)

	var median float64
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}

	return median
}

// REVERSE FUNCTION
func reverse(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// FIBONACCI FUNCTION
func fibonacci(n int) int {
	var result int
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			result = i
		} else {
			result = fibonacci(n-1) + fibonacci(n-2)
		}
	}
	return result
}
