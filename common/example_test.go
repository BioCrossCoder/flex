package common

import (
	"fmt"
)

func ExampleEqual() {
	// map
	fmt.Println(Equal(map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2}))
	fmt.Println(Equal(map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 3}))
	// slice
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 4}))
	// struct
	fmt.Println(Equal(struct{ a int }{1}, struct{ a int }{1}))
	fmt.Println(Equal(struct{ a int }{1}, struct{ a int }{2}))
	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
}

func ExampleLen() {
	// string
	fmt.Println(Len("hello"))
	// map
	fmt.Println(Len(map[string]int{"a": 1, "b": 2}))
	// slice
	fmt.Println(Len([]int{1, 2, 3}))
	// Output:
	// 5
	// 2
	// 3
}

func ExampleContains() {
	// string
	fmt.Println(Contains("hello", "l"))
	fmt.Println(Contains("hello", "x"))
	// slice
	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]int{1, 2, 3}, 4))
	// array
	fmt.Println(Contains([3]int{1, 2, 3}, 2))
	fmt.Println(Contains([5]int{1, 2, 3}, 4))
	// map
	fmt.Println(Contains(map[string]int{"a": 1, "b": 2}, "a"))
	fmt.Println(Contains(map[string]int{"a": 1, "b": 2}, 1))
	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
	// false
	// true
}

func ExampleCount() {
	// string
	fmt.Println(Count("hello", "l"))
	// slice
	fmt.Println(Count([]int{1, 2, 3}, 2))
	// array
	fmt.Println(Count([4]int{1, 2, 3, 2}, 2))
	// map
	fmt.Println(Count(map[string]int{"a": 1, "b": 2}, "a"))
	fmt.Println(Count(map[string]int{"a": 1, "b": 2}, 1))
	// Output:
	// 2
	// 1
	// 2
	// 0
	// 1
}
