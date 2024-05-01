package functools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func f(a, b string) string {
	return a + b
}

func TestReduce(t *testing.T) {
	convey.Convey("Call Reduce on [array]", t, func() {
		s := [3]string{"a", "b", "c"}
		expected := "abc"
		actual, err := Reduce(f, s)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	convey.Convey("Call Reduce on [slice]", t, func() {
		s := []string{"h", "e", "l", "l", "o"}
		expected := "hello"
		actual, err := Reduce(f, s)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
	convey.Convey("Call Reduce on [string]", t, func() {
		s := "hello"
		actual, err := Reduce(f, s)
		assert.Nil(t, err)
		assert.Equal(t, s, actual)
	})
}

func TestReduceMap(t *testing.T) {
	convey.Convey("Call Reduce on [map]", t, func() {
		m := map[int]string{
			1: "one",
			2: "two",
		}
		candidates := []string{"onetwo", "twoone"}
		actual, err := ReduceMap(f, m)
		assert.Nil(t, err)
		assert.Contains(t, candidates, actual)
	})
}
