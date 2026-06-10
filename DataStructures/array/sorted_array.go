package array

import (
	"cmp"
	"fmt"
)

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

func (u *SortedArray[T]) insert(value T) error {
	if u.Size >= u.MaxSize {
		return fmt.Errorf("The array is already full, maximum size: %d", u.MaxSize)
	}

	for i := u.Size; i > 0; i-- {
		if u.Array[i-1] <= value {
			u.Array[i] = value
			u.Size += 1
			return nil
		} else {
			u.Array[i] = u.Array[i-1]
		}
	}

	u.Array[0] = value
	u.Size += 1

	return nil
}
