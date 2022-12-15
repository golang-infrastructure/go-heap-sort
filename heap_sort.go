package heap_sort

import (
	compare_anything "github.com/golang-infrastructure/go-compare-anything"
	"github.com/golang-infrastructure/go-gtypes"
	"github.com/golang-infrastructure/go-heap"
)

// Sort 对元素类型为可比较类型的切片进行正序排序
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

// ReverseSort 对元素类型为可比较类型的切片进行逆向排序
func ReverseSort[T gtypes.Ordered](slice []T) {
	SortByComparator(slice, func(a T, b T) int {
		if a == b {
			return 0
		} else if a > b {
			return -1
		} else {
			return 1
		}
	})
}

// SortByComparator 使用自定义的比较器对切片进行排序，如果需要逆序可以自行控制比较器的比较规则
func SortByComparator[T any](slice []T, comparator compare_anything.Comparator[T]) {
	heapSort := heap.New[T](heap.Comparator[T](comparator))
	heapSort.Push(slice...)
	index := 0
	for heapSort.IsNotEmpty() {
		slice[index] = heapSort.Pop()
		index++
	}
}
