package array

import (
	"cmp"
	"fmt"
)

type UnsortedArray[T cmp.Ordered] struct {
	Array   []T
	MaxSize int
	Size    int
}

func NewUnsortedArray[T cmp.Ordered](maxSize int) UnsortedArray[T] {
	return UnsortedArray[T]{
		Array:   make([]T, maxSize),
		MaxSize: maxSize,
		Size:    0,
	}
}

func (u *UnsortedArray[T]) insert(value T) error {
	if u.Size >= u.MaxSize {
		return fmt.Errorf("The array is already full")
	}

	u.Array[u.Size] = value
	u.Size += 1

	return nil
}

func (u *UnsortedArray[T]) delete(index int) error {
	if u.Size == 0 {
		return fmt.Errorf("Delete from an empty array")
	} else if index < 0 || index >= u.Size {
		return fmt.Errorf("Index %d out of range", index)
	}

	u.Array[index] = u.Array[u.Size-1]
	u.Size -= 1

	return nil
}

func (u *UnsortedArray[T]) find(target T) int {
	for index := range u.Size {
		if u.Array[index] == target {
			return index
		}
	}

	return -1
}

func (u *UnsortedArray[T]) traverse(op func(v T)) {
	for index := range u.Size {
		op(u.Array[index])
	}
}

func (u *UnsortedArray[T]) maxInArray() (error, int, T) {
	var maximumValue T
	if u.Size == 0 {
		return fmt.Errorf("Max of an empty array"), -1, maximumValue
	}

	maximum := 0
	for index := range u.Size {
		if u.Array[index] > u.Array[maximum] {
			maximum = index
		}
	}

	return nil, maximum, u.Array[maximum]
}

func (u *UnsortedArray[T]) minInArray() (error, int, T) {
	var minimumValue T
	if u.Size == 0 {
		return fmt.Errorf("Min of an empty array"), -1, minimumValue
	}

	minimum := 0
	for index := range u.Size {
		if u.Array[index] < u.Array[minimum] {
			minimum = index
		}
	}

	return nil, minimum, u.Array[minimum]
}

func (u *UnsortedArray[T]) minAndMaxInArray() (error, T, T) {
	var minimumValue, maximumValue T
	if u.Size == 0 {
		return fmt.Errorf("Min and Max of an empty array."), minimumValue, maximumValue
	}

	minimum, maximum := 0, 0
	for index := range u.Size {
		indexValue := u.Array[index]
		if indexValue < u.Array[minimum] {
			minimum = index
		}

		if indexValue > u.Array[maximum] {
			maximum = index
		}
	}

	return nil, u.Array[minimum], u.Array[maximum]
}
