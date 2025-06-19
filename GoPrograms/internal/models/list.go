package models

type List[T any] struct {
	Value T
	Next  *List[T]
}
