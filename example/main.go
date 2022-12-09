package main

import (
	"fmt"
	heap_sort "github.com/golang-infrastructure/go-heap-sort"
)

func main() {

	slice := []int{100, 2, 1, 10}
	heap_sort.Sort(slice)
	fmt.Println(slice)
	// Output:
	// [1 2 10 100]

}
