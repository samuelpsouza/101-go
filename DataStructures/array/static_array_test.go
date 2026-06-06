package array

import (
	"testing"
)

func TestUnsortedArray(t *testing.T) {
	t.Run("NewUnsortedArray", testNewUnsortedArray)
	t.Run("Insert", testInsert)
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

func testInsert(t *testing.T) {
	t.Run("Successful inserts", func(t *testing.T) {
		arrSize := 5
		arr := NewUnsortedArray[int](arrSize)

		for i := 0; i < arrSize; i++ {
			if err := arr.insert((i + 1) * 10); err != nil {
				t.Fatalf("Insert %d failed: %v", (i+1)*10, err)
			}
		}

		if arr.Size != arrSize {
			t.Errorf("Expected size=5, got %d", arr.Size)
		}
	})

	t.Run("Insert when full", func(t *testing.T) {
		arr := NewUnsortedArray[int](1)
		arr.insert(100)

		err := arr.insert(200)
		if err == nil || err.Error() != "The array is already full" {
			t.Errorf("Expected 'THe array is already full' error, got %v", err)
		}

	})
}
