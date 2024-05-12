// Package functools provides functional programming tools.
package functools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	f := func(a any) bool {
		return a == 0
	}
	convey.Convey("true", t, func() {
		entry := []int{1, 0, 1}
		result, err := Any(entry, f)
		assert.Nil(t, err)
		assert.True(t, result)
	})
	convey.Convey("false", t, func() {
		entry := []int{1, 3, 2}
		result, err := Any(entry, f)
		assert.Nil(t, err)
		assert.False(t, result)
	})
}
