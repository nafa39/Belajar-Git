package main

import (
	"fmt"
	"generic-go/helpers"
)

//learn about generic in go

func SumInt(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumFloat(numbers []float32) float32 {
	sum := float32(0)
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//generic function

func Hello[T int | float32 | string](param T) T {
	fmt.Println("Hello", param)
	return param
}

//type set

type Number interface {
	int | float32 | int32 | int64 | float64
}

func Sum[T Number](numbers []T) T {
	var sum T
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// multiple type
func MultipleType[T int | float32, K string | bool](num T, ex K) {
	fmt.Println(num, ex)
}

// type data any
func Any[T any](param T) {
	fmt.Println(param)
}

// type data interface
func Interface[T interface{}](param T) {
	fmt.Println(param)
}

// type data comparable
func Comparable[T comparable](param T) {
	fmt.Println(param)
}

// function generic interface
func ChangeValue[T comparable](entity helpers.GetterSetter[T], value T) {
	entity.Set(value)
	fmt.Println("success change value to", entity.Get())
}

// INCLASS
func halo[T any](params T) {
	fmt.Println(params)
}

func hallo[T int | float64 | bool](params T) {
	fmt.Println(params)
}

func genericsTwo[asd int | float64, qwe int | float64](params1 asd, params2 qwe) {
	fmt.Println(params1)
	fmt.Println(params2)
}

func sum[T int | float64](a T, b T) T {
	return a + b
}

func sumInt(a int, b int) int {
	return a + b
}

func sumFloat64(a float64, b float64) float64 {
	return a + b
}

func merge[T any](a []T, b []T) []T {
	return append(a, b...)
}

func main() {
	var numbers = []int{1, 2, 3, 4, 5}
	var numbersFloat = []float32{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println(SumInt(numbers))
	fmt.Println(SumFloat(numbersFloat))

	//generic function
	Hello[string]("Airell")
	Hello[int](23)

	//type set
	fmt.Println(Sum(numbers))
	fmt.Println(Sum(numbersFloat))

	//multiple type
	MultipleType[int, string](23, "Airell")

	//type data any
	Any[int](23)

	//type data interface
	Interface[int](23)

	//type data comparable
	Comparable[int](23)

	//generic struct
	var personInt = helpers.Person[int]{Name: "Airell", Scores: []int{1, 2, 3}}
	personInt.AddScores(4)
	personInt.AddScores(5)
	fmt.Println(personInt)

	//generic interface
	var redisInt = helpers.Redis[int]{Value: 10}
	ChangeValue(&redisInt, 20)

	//INCLASS
	halo("Hello World")
	halo(123)
	halo(1.23)
	halo(true)

	hallo(123)
	hallo(1.23)
	hallo(true)

	genericsTwo(123, 1.23)

	hallo(10)
	hallo(10.5)
	hallo(true)

	genericsTwo(10, 10.5)
	genericsTwo(10.5, 10)

	fmt.Println(sumInt(10, 10))
	fmt.Println(sumFloat64(10.5, 10.5))

	fmt.Println(sum(10, 10))
	fmt.Println(sum(10.5, 10.5))

	fmt.Println(merge([]int{1, 2, 3}, []int{4, 5, 6}))
}
