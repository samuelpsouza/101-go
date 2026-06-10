package array

import (
	"math/rand/v2"
	"testing"
)

func TestNewSortedArray(t *testing.T) {
	tests := []struct {
		name    string
		maxSize int
	}{
		{"Zero Size", 0},
		{"Single element", 1},
		{"Normal Size", 10},
		{"Large Size", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewSortedArray[int](tt.maxSize)
			if arr.MaxSize != tt.maxSize {
				t.Errorf("Expected MaxSize=%d, got %d", tt.maxSize, arr.MaxSize)
			}

			if arr.Size != 0 {
				t.Errorf("Expected Size=0, got %d", arr.Size)
			}

			if len(arr.Array) != tt.maxSize {
				t.Errorf("Expected underlying array length=%d, got %d", tt.maxSize, len(arr.Array))
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	t.Run("Successful inserts", func(t *testing.T) {
		const arrSize = 5
		arr := NewSortedArray[int](arrSize)

		for i := 0; i < arrSize; i++ {
			randNum := rand.IntN(100) + 1
			if err := arr.insert(randNum); err != nil {
				t.Fatalf("Operation arr.insert(%d) aborted at iteration index %d: %v", randNum, i, err)
			}
		}

		if arr.Size != arrSize {
			t.Errorf("Structural size invariant violation: expected %d, evaluated %d", arrSize, arr.Size)
		}

		for v := range arr.Size - 1 {
			if arr.Array[v] > arr.Array[v+1] {
				t.Errorf("Sorting order assertion failed at index %d: element %d exceeds subsequent element %d", v, arr.Array[v], arr.Array[v+1])
			}
		}
	})
}
