package itertools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCycle(t *testing.T) {
	convey.Convey("cycle slice", t, func() {
		entry := []int{1, 2, 3}
		iterator := Cycle(entry)
		for i := 0; i < 3; i++ {
			assert.Equal(t, entry, iterator.Pour())
		}
	})
}

func ExampleCycle() {
	entry := []int{1, 2, 3}
	c := Cycle(entry)
	for i := 0; i < 10; i++ {
		c.Next()
		fmt.Println(c.Value())
	}
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
	// 1
}
