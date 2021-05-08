## Sort Library

Sort Library is a library written in Go (Golang), If you need to sort, you can use this library

## Included sorting method

+ Select sort

  ```go
  func SelectSort (){}
  ```

+ Insertion sort

  ```
  func InsertSort (){}
  ```

+ Bubble Sort

  ```go
  func BubbleSort (){}
  ```

+ Merge sort

  ```go
  func MergeSort (){}
  ```

+ Quick sort

  ```go
  func QuickSort (){}
  ```

+ Hill sort

  ```go
  func ShellSort (){}
  ```

+ Heap sort

  ```
  func HeapSort (){}
  ```

## Installation

To installpackage, you need to install Go and set your Go workspace first.

```shell
$ go get github.com/Fish-pro/sort-library
```

Import it in your code:

```shell
import sort_go "github.com/Fish-pro/sort-library"
```

## Quick Start

```shell
$ mkdir test-sort
$ cd test-srot
$ vim main.go
```

```go
package main

import (
	"fmt"
	sort_go "github.com/Fish-pro/sort-library"
)

func main(){
	array := []int{232,3,234,23,43,343}
	sort_go.ShellSort(array)
	fmt.Println(array)
}
```

```shell
$ go run main.go
[3 23 43 232 234 343]
```



