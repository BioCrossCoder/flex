package functools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	convey.Convey("call All on int array", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) bool {
			return x%2 == 0
		}
		g := func(x int) bool {
			return x > 0
		}
		assert.False(t, All(f, arr))
		assert.True(t, All(g, arr))
		assert.True(t, All(f, []int{}))
		assert.True(t, All(g, []int{}))
	})
	convey.Convey("call All on string array", t, func() {
		arr := []string{"go", "python", "java"}
		f := func(x string) bool {
			return len(x) > 3
		}
		g := func(x string) bool {
			return len(x) > 1
		}
		assert.False(t, All(f, arr))
		assert.True(t, All(f, []string{}))
		assert.True(t, All(g, arr))
		assert.True(t, All(g, []string{}))
	})
}
