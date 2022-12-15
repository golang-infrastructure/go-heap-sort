# 堆排序（Heap Sort）

# 一、安装

```bash
go get -u github.com/golang-infrastructure/go-heap-sort
```

# 二、示例代码

```go
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
```

# 三、API

对元素类型为可比较类型的切片进行正序排序

```go
func Sort[T gtypes.Ordered](slice []T)
```

对元素类型为可比较类型的切片进行逆向排序

```go
func ReverseSort[T gtypes.Ordered](slice []T) 
```

使用自定义的比较器对切片进行排序，如果需要逆序可以自行控制比较器的比较规则

```go
func SortByComparator[T any](slice []T, comparator compare_anything.Comparator[T])
```





