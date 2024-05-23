package arraylist

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLen(t *testing.T) {
	convey.Convey("check length", t, func() {
		l := ArrayList[int]{1, 2, 3}
		assert.Equal(t, len(l), l.Len())
	})
}

func TestCount(t *testing.T) {
	convey.Convey("count specific element", t, func() {
		l := ArrayList[int]{1, 2, 3, 2, 1, 4, 5, 4}
		assert.Equal(t, 2, l.Count(2))
		assert.Equal(t, 2, l.Count(4))
		assert.Equal(t, 0, l.Count(6))
	})
}

func TestIncludes(t *testing.T) {
	convey.Convey("check if element is included", t, func() {
		l := ArrayList[int]{1, 2, 3, 4, 5}
		assert.True(t, l.Includes(3))
		assert.False(t, l.Includes(6))
	})
}

func TestEmpty(t *testing.T) {
	convey.Convey("check if list is empty", t, func() {
		l1 := ArrayList[int]{1, 2, 3}
		l2 := ArrayList[any]{}
		assert.False(t, l1.Empty())
		assert.True(t, l2.Empty())
	})
}

func ExampleArrayList() {
	al := ArrayList[int]{1, 2, 3, 4, 5}
	fmt.Println(al) // Output: [1 2 3 4 5]
}

func ExampleArrayList_Len() {
	al := ArrayList[int]{1, 2, 3, 4, 5}
	length := al.Len()
	fmt.Println(length) // Output: 5
}

func ExampleArrayList_Count() {
	al := ArrayList[int]{1, 2, 3, 4, 5, 5, 5}
	fmt.Println(al.Count(5))
	fmt.Println(al.Count(6))
	// Output:
	// 3
	// 0
}

func ExampleArrayList_Includes() {
	al := ArrayList[int]{1, 2, 3, 4, 5}
	fmt.Println(al.Includes(3))
	fmt.Println(al.Includes(6))
	// Output:
	// true
	// false
}

func ExampleArrayList_Empty() {
	al := ArrayList[int]{}
	fmt.Println(al.Empty())
	al2 := ArrayList[int]{1, 2, 3}
	fmt.Println(al2.Empty())
	// Output:
	// true
	// false
}

func ExampleArrayList_Equal() {
	al1 := ArrayList[int]{1, 2, 3}
	al2 := ArrayList[int]{1, 2, 3}
	al3 := ArrayList[int]{1, 2, 3, 4}
	al4 := ArrayList[int]{1, 2, 4}
	fmt.Println(al1.Equal(al2))
	fmt.Println(al1.Equal(al3))
	fmt.Println(al1.Equal(al4))
	// Output:
	// true
	// false
	// false
}
