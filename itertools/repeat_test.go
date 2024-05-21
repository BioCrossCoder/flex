package itertools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	convey.Convey("repeat", t, func() {
		entry := "233"
		r := Repeat(entry)
		assert.Equal(t, []any{"233", "233", "233"}, r.Repeat(3))
	})
}

func ExampleRepeat() {
	r := Repeat("233")
	for i := 0; i < 3; i++ {
		fmt.Println(r.Next())
	}
	fmt.Println(r.Repeat(3))
	// Output:
	// 233
	// 233
	// 233
	// [233 233 233]
}
