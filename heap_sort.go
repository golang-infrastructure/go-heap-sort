package heap_sort

import (
	compare_anything "github.com/golang-infrastructure/go-compare-anything"
	"github.com/golang-infrastructure/go-gtypes"
	"github.com/golang-infrastructure/go-heap"
)

func Sort[T gtypes.Ordered](slice []T) {
	SortByComparator(slice, func(a T, b T) int {
		if a == b {
			return 0
		} else if a > b {
			return 1
		} else {
			return -1
		}
	})
}

func SortByComparator[T any](slice []T, comparator compare_anything.Comparator[T]) {
	heap := heap.New[T](heap.Comparator[T](comparator))
	heap.Push(slice...)
	index := 0
	for heap.IsNotEmpty() {
		slice[index] = heap.Pop()
		index++
	}
}
