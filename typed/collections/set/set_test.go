package set

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestSet() Set[int] {
	return Of([]int{1, 2, 3}...)
}

func TestSet(t *testing.T) {
	convey.Convey("get size of a copy of a set", t, func() {
		s := newTestSet()
		assert.True(t, s.Equal(s.Copy()))
	})
}
