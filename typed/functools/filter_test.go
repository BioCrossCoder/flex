package functools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	convey.Convey("Call Filter on int array", t, func() {
		arr := []int{1, 2, 3, 4, 5}
		f := func(x int) bool {
			return x%2 == 0
		}
		expected := []int{2, 4}
		actual := Filter(f, arr)
		assert.Equal(t, expected, actual)
	})
	convey.Convey("Call Filter on string array", t, func() {
		arr := []string{"apple", "banana", "cherry", "date", "elderberry"}
		f := func(x string) bool {
			return strings.HasSuffix(x, "erry")
		}
		expected := []string{"cherry", "elderberry"}
		actual := Filter(f, arr)
		assert.Equal(t, expected, actual)
	})
}
