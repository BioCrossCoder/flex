package itertools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter(t *testing.T) {
	convey.Convey("count", t, func() {
		c := NewCounter(10, 2, true)
		assert.Equal(t, 10, c.Count())
		for i := 1; i < 10; i++ {
			assert.Equal(t, 10-i*2, c.Count())
		}
	})
	convey.Convey("reset", t, func() {
		c := NewCounter(6, 3, false)
		assert.Equal(t, 6, c.Count())
		for i := 0; i < 10; i++ {
			_ = c.Count()
		}
		assert.NotEqual(t, 6, c.Count())
		c.Reset()
		assert.Equal(t, 6, c.Count())
	})
	convey.Convey("jump", t, func() {
		c := NewCounter(1, 5, false)
		assert.Equal(t, 1, c.Count())
		assert.Equal(t, 11, c.Jump(2))
	})
}

func ExampleCounter() {
	// reverse counting from 10 with step 2
	c := NewCounter(10, 2, true)
	for i := 0; i < 3; i++ {
		fmt.Println(c.Count())
	}
	// Output:
	// 10
	// 8
	// 6
}

func ExampleCounter_Count() {
	// counting from 1 with step 2
	c := NewCounter(1, 2, false)
	for i := 0; i < 5; i++ {
		fmt.Println(c.Count())
	}
	// Output:
	// 1
	// 3
	// 5
	// 7
	// 9
}

func ExampleCounter_Reset() {
	// counting from 1 with step 2
	c := NewCounter(1, 2, false)
	for i := 0; i < 3; i++ {
		fmt.Println(c.Count())
	}
	c.Reset()
	for i := 0; i < 3; i++ {
		fmt.Println(c.Count())
	}
	// Output:
	// 1
	// 3
	// 5
	// 1
	// 3
	// 5
}

func ExampleCounter_Jump() {
	// counting from 1 with step 2
	c := NewCounter(1, 2, false)
	fmt.Println(c.Count())
	fmt.Println(c.Jump(3))
	fmt.Println(c.Count())
	// Output:
	// 1
	// 7
	// 9
}
