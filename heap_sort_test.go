package heap_sort

import "testing"

func TestSort(t *testing.T) {
	slice := []int{100, 2, 1, 10}
	Sort(slice)
	t.Log(slice)
}
