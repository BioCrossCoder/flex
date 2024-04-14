package functools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReduce(t *testing.T) {
	f := func(a, b string) string {
		return a + b
	}
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
