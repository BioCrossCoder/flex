package functools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCount(t *testing.T) {
	convey.Convey("count element satisfying condition", t, func() {
		assert.Equal(t, CountBy([]int{1, 2, 3}, func(a int) bool {
			return fmt.Sprint(a) == "1"
		}), 1)
		assert.Equal(t, Count([]int{1, 2, 3}, 1), 1)
	})
}
