# Flex

A Golang library to make development easier and more efficient

## Getting Started

### Prerequisites

Flex requires Go version `1.18` or above.

### Getting Flex

With Go's module support, `go [build|run|test]` automatically fetches the necessary dependencies when you add the import in your code:

```go
import "github.com/biocrosscoder/flex"
```

Alternatively, use `go get`:

```sh
go get -u github.com/biocrosscoder/flex
```

***The package path above is not actually used in the code, use the following pacakge paths when necessary:***

+ github.com/biocrosscoder/flex/itertools
+ github.com/biocrosscoder/flex/functools
+ github.com/biocrosscoder/flex/collections
+ github.com/biocrosscoder/flex/collections/arraylist
+ github.com/biocrosscoder/flex/collections/linkedlist
+ github.com/biocrosscoder/flex/collections/queue
+ github.com/biocrosscoder/flex/collections/set
+ github.com/biocrosscoder/flex/collections/dict
+ github.com/biocrosscoder/flex/collections/orderedcontainers
+ github.com/biocrosscoder/flex/typed/itertools
+ github.com/biocrosscoder/flex/typed/functools
+ github.com/biocrosscoder/flex/typed/collections
+ github.com/biocrosscoder/flex/typed/collections/arraylist
+ github.com/biocrosscoder/flex/typed/collections/linkedlist
+ github.com/biocrosscoder/flex/typed/collections/queue
+ github.com/biocrosscoder/flex/typed/collections/set
+ github.com/biocrosscoder/flex/typed/collections/dict
+ github.com/biocrosscoder/flex/typed/collections/orderedcontainers
+ github.com/biocrosscoder/flex/typed/collections/sortedcontainers
+ github.com/biocrosscoder/flex/typed/collections/sortedcontainers/sortedlist

### Running Flex

Example 1: Remove duplicates from a slice

```go
package main

import (
 "fmt"
 "github.com/biocrosscoder/flex/typed/collections/orderedcontainers"
 "math/rand"
)

func main() {
 arr := make([]int, 20)
 for i := 0; i < 20; i++ {
  arr[i] = rand.Intn(10)
 }
 fmt.Println("Original array:", arr)
 rd := orderedcontainers.NewOrderedSet(arr...).Elements()
 fmt.Println("Removed duplicates: ", rd)
}
```

Example 2: Sort a slice of complicated structures

```go
package main

import (
 "fmt"
 "github.com/biocrosscoder/flex/typed/functools"
 "math/rand"
)

func main() {
 arr := make([][]int, 10)
 for i := 0; i < 10; i++ {
  arr[i] = []int{rand.Intn(5), rand.Intn(5)}
 }
 fmt.Println("Before sorting:", arr)
 f1 := func(x, y []int) int {
  return x[0] - y[0]
 }
 f2 := func(x, y []int) int {
  return y[1] - x[1]
 }
 functools.Sort(arr, f1, f2)
 fmt.Println("After sorting: ", arr)
}
```

## Features

1. `itertools` provide a series of functions to create iterators.
2. `functools` provide a series of functions to support functional programming.
3. `collections` provide a series of powerful and practical containers to store and manipulate data.
4. `typed` provide `itertools`, `functools` and `collections` with `generics` support.
5. This library mainly refers to `Python` and `JavaScript`, considering the feature of `Golang` as well.
