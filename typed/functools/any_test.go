package functools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	convey.Convey("call Any on int array", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) bool {
			return x > 10
		}
		g := func(x int) bool {
			return x > 3
		}
		assert.False(t, Any(f, arr))
		assert.True(t, Any(g, arr))
		assert.False(t, Any(f, []int{}))
		assert.False(t, Any(g, []int{}))
	})
	convey.Convey("call Any on string array", t, func() {
		arr := []string{"apple", "banana", "orange"}
		f := func(x string) bool {
			return x == "apple"
		}
		g := func(x string) bool {
			return x == "grape"
		}
		assert.True(t, Any(f, arr))
		assert.False(t, Any(f, []string{}))
		assert.False(t, Any(g, arr))
		assert.False(t, Any(g, []string{}))
	})
}
