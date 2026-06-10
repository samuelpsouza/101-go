package array

import "cmp"

type SortedArray[T cmp.Ordered] struct {
	Array   []T
	MaxSize int
	Size    int
}

func NewSortedArray[T cmp.Ordered](maxSize int) SortedArray[T] {
	return SortedArray[T]{
		Array:   make([]T, maxSize),
		MaxSize: maxSize,
		Size:    0,
	}
}
