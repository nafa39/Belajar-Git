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
}
