package functools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	f := func(a string) bool {
		return a > "0"
	}
	convey.Convey("Call Filter on [array]", t, func() {
		s := [3]string{"0", "1", "2"}
		expected := []any{"1", "2"}
		actual, err := Filter(f, s)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	convey.Convey("Call Filter on [slice]", t, func() {
		s := []string{"0", "1", "2", "3"}
		expected := []any{"1", "2", "3"}
		actual, err := Filter(f, s)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	convey.Convey("Call Filter on [string]", t, func() {
		s := "012345"
		expected := []any{"1", "2", "3", "4", "5"}
		actual, err := Filter(f, s)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	convey.Convey("Call Filter on [map]", t, func() {
		m := map[string]string{
			"zero": "0",
			"one":  "1",
			"two":  "2",
		}
		expected := map[any]any{
			"one": "1",
			"two": "2",
		}
		actual, err := Filter(f, m)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}
