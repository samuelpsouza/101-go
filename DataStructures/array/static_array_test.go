package array

import (
	"testing"
)

func TestUnsortedArray(t *testing.T) {
	t.Run("NewUnsortedArray", testNewUnsortedArray)
}

func testNewUnsortedArray(t *testing.T) {
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
			arr := NewUnsortedArray[int](tt.maxSize)
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
