package array

import (
	"fmt"
	"testing"
)

func TestUnsortedArray(t *testing.T) {
	t.Run("NewUnsortedArray", testNewUnsortedArray)
	t.Run("Insert", testInsert)
	t.Run("Delete", testDelete)
	t.Run("Find", testFind)
	t.Run("Traverse", testTraverse)
	t.Run("TraverseEmptyArray", testTraverseEmptyArray)
	t.Run("MaxInArray", testMaxInArray)
	t.Run("MaxInEmptyArray", testMaxInEmptyArray)
	t.Run("MinInArray", testMinInArray)
	t.Run("MinAndMaxInArray", testMinAndMaxInArray)
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

func testDelete(t *testing.T) {
	t.Run("Normal deletes", func(t *testing.T) {
		arr := NewUnsortedArray[int](5)
		arr.insert(10)
		arr.insert(20)
		arr.insert(30)
		arr.insert(40)

		if err := arr.delete(1); err != nil {
			t.Fatalf("Delete failed: %v", err)
		}

		if arr.Size != 3 {
			t.Errorf("Expected size 3, got %d", arr.Size)
		}

	})

	t.Run("Delete the last element", func(t *testing.T) {
		arr := NewUnsortedArray[int](3)
		arr.insert(10)
		arr.insert(20)

		if err := arr.delete(1); err != nil {
			t.Fatalf("Delete failed: %v", err)
		}

		if arr.Size != 1 || arr.Array[0] != 10 {
			t.Errorf("Delete last element failed")
		}
	})

	t.Run("Delete only element", func(t *testing.T) {
		arr := NewUnsortedArray[int](1)
		arr.insert(10)

		if err := arr.delete(0); err != nil {
			t.Fatalf("Delete failed: %v", err)
		}

		if arr.Size != 0 {
			t.Errorf("Size should be 0 after deleting only element")
		}
	})

	t.Run("Error cases", func(t *testing.T) {
		arr := NewUnsortedArray[int](5)

		if err := arr.delete(0); err == nil {
			t.Errorf("Expected error when deleting from empty array")
		}

		arr.insert(10)
		arr.insert(20)

		if err := arr.delete(-1); err == nil {
			t.Errorf("Expected error for negative index")
		}

		if err := arr.delete(5); err == nil {
			t.Errorf("Expected error for index >= Size")
		}
	})
}

func testFind(t *testing.T) {
	arr := NewUnsortedArray[int](10)
	arr.insert(5)
	arr.insert(10)
	arr.insert(15)
	arr.insert(10)

	tests := []struct {
		target   int
		expected int
	}{
		{5, 0},
		{10, 1},
		{15, 2},
		{99, -1},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Find %d", tt.target), func(t *testing.T) {
			got := arr.find(tt.target)
			if got != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, got)
			}
		})
	}

	arr.delete(1)
	if idx := arr.find(10); idx != 1 {
		t.Errorf("Find after delete failed, got %d", idx)
	}
}

func testTraverse(t *testing.T) {
	arr := NewUnsortedArray[int](5)
	arr.insert(1)
	arr.insert(2)
	arr.insert(3)

	collected := []int{}
	arr.traverse(func(v int) {
		collected = append(collected, v)
	})

	expected := []int{1, 2, 3}
	if len(collected) != len(expected) {
		t.Fatalf("Expected %d elements, got %d", len(expected), len(collected))
	}

	for i, v := range expected {
		if collected[i] != v {
			t.Errorf("Traverse mismatch at index %d", i)
		}
	}
}

func testTraverseEmptyArray(t *testing.T) {
	empty := NewUnsortedArray[int](5)
	called := false

	empty.traverse(func(v int) {
		called = true
	})

	if called {
		t.Errorf("Traverse should not call func on empty array")
	}
}

func testMaxInArray(t *testing.T) {
	arr := NewUnsortedArray[int](5)
	arr.insert(-5)
	arr.insert(10)
	arr.insert(3)
	arr.insert(7)

	err, idx, val := arr.maxInArray()
	if err != nil {
		t.Fatal(err)
	}

	if val != 10 || idx != 1 {
		t.Errorf("Expected max=10 at index 1, got %d at %d", val, idx)
	}
}

func testMaxInEmptyArray(t *testing.T) {
	arr := NewUnsortedArray[int](5)

	err, _, _ := arr.maxInArray()
	if err == nil {
		t.Errorf("Expected error for empty array")
	}
}

func testMinInArray(t *testing.T) {
	arr := NewUnsortedArray[int](5)
	arr.insert(-5)
	arr.insert(10)
	arr.insert(3)
	arr.insert(7)

	err, idx, val := arr.minInArray()
	if err != nil {
		t.Fatal(err)
	}

	if val != -5 || idx != 0 {
		t.Errorf("Expected max=10 at index 1, got %d at %d", val, idx)
	}
}

func testMinAndMaxInArray(t *testing.T) {
	arr := NewUnsortedArray[int](5)
	arr.insert(-5)
	arr.insert(10)
	arr.insert(3)
	arr.insert(7)

	err, min, max := arr.minAndMaxInArray()
	if err != nil {
		t.Fatal(err)
	}

	if min != -5 || max != 10 {
		t.Errorf("Expected min=-5, max=10, got min=%v, max=%v", min, max)
	}
}
