package sorting

import "testing"

func TestSelectionSort(t *testing.T) {
	if val := SelectionSort([]int{1, 32, 8, 9, 4, 0}); !testEq(val, []int{0, 1, 4, 8, 9, 32}) {
		t.Error("Function returned other than expected array: ", val)
	}
	if val := SelectionSort([]int{1, 32, 8, 9, 4, 0, -4}); !testEq(val, []int{-4, 0, 1, 4, 8, 9, 32}) {
		t.Error("Function returned other than expected array: ", val)
	}
	SelectionSort([]int{12, 2, 24, 15, 35, 56, 62, 44, 27})
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		SelectionSort([]int{-4, 0, 1, 4, 8, 9, 32})
	}
}

func TestInsertionSort(t *testing.T) {
	if val := InsertionSort([]int{1, 8, 32, 9, -4, 0}); !testEq(val, []int{-4, 0, 1, 8, 9, 32}) {
		t.Error("Function returned other than expected array: ", val)
	}
	if val := InsertionSort([]int{1, 32, 8, 9, 2, 0}); !testEq(val, []int{0, 1, 2, 8, 9, 32}) {
		t.Error("Function returned other than expected array: ", val)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		InsertionSort([]int{-4, 0, 1, 4, 8, 9, 32})
	}
}

func TestBubbleSort(t *testing.T) {
	if val := BubbleSort([]int{1, 8, 32, 9, -4, -1, 0}); !testEq(val, []int{-4, -1, 0, 1, 8, 9, 32}) {
		t.Error("Function returned other than expected array: ", val)
	}
	if val := BubbleSort([]int{1, 32, 8, 9, 2, 0}); !testEq(val, []int{0, 1, 2, 8, 9, 32}) {
		t.Error("Function returned other than expected array: ", val)
	}

}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		BubbleSort([]int{-4, 0, 1, 4, 8, 9, 32})
	}
}

func TestMergeSort(t *testing.T) {
	if val := MergeSort([]int{1, 8, 32, 9, -4, -1, 0}); !testEq(val, []int{-4, -1, 0, 1, 8, 9, 32}) {
		t.Error("Function returned other than expected array: ", val)
	}
	if val := MergeSort([]int{1, 32, 8, 9, 2, 0}); !testEq(val, []int{0, 1, 2, 8, 9, 32}) {
		t.Error("Function returned other than expected array: ", val)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		MergeSort([]int{-4, 0, 1, 4, 8, 9, 32})
	}
}

func TestHeapSort(t *testing.T) {
	if val := HeapSort([]int{1, 8, 32, 9, -4, -1, 0}); !testEq(val, []int{-4, -1, 0, 1, 8, 9, 32}) {
		t.Error("Function returned other than expected array: ", val)
	}
	if val := HeapSort([]int{1, 32, 8, 9, 2, 0}); !testEq(val, []int{0, 1, 2, 8, 9, 32}) {
		t.Error("Function returned other than expected array: ", val)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		HeapSort([]int{-4, 0, 1, 4, 8, 9, 32})
	}
}

func TestCycleSort(t *testing.T) {
	if val, _ := CycleSort([]int{1, 8, 32, 9, -4, -1, 0}); !testEq(val, []int{-4, -1, 0, 1, 8, 9, 32}) {
		t.Error("Function returned other than expected array: ", val)
	}
	if val, _ := CycleSort([]int{1, 32, 8, 9, 2, 0}); !testEq(val, []int{0, 1, 2, 8, 9, 32}) {
		t.Error("Function returned other than expected array: ", val)
	}
}

func BenchmarkCycleSort(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		CycleSort([]int{-4, 0, 1, 4, 8, 9, 32})
	}
}
