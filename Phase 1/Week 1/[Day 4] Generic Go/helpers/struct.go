package helpers

// generic struct
type Person[T int | float64] struct {
	Name   string
	Scores []T
}

// generic Interface
type GetterSetter[T comparable] interface {
	Get() T
	Set(newVal T)
}

type Redis[T comparable] struct {
	Value T
}
