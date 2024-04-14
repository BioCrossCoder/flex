package functools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	f := func(a int) bool {
		return a > 0
	}
	convey.Convey("true", t, func() {
		entry, _ := Map(f, []int{1, 2, 3})
		result, err := All(entry)
		assert.Nil(t, err)
		assert.True(t, result)
	})
	convey.Convey("false", t, func() {
		entry, _ := Map(f, []int{1, 2, 0})
		result, err := All(entry)
		assert.Nil(t, err)
		assert.False(t, result)
	})
}
