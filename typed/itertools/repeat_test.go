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
		r := Repeat(entry, 3)
		assert.Equal(t, []string{"233", "233", "233"}, r.Pour())
	})
}

func ExampleRepeat() {
	iter := Repeat("233", 3)
	for iter.Next() {
		fmt.Println(iter.Value())
	}
	// Output:
	// 233
	// 233
	// 233
}
