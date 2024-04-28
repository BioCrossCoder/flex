package set

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestSet() Set {
	return Of([]any{1, 2, 3}...)
}

func TestSet(t *testing.T) {
	convey.Convey("get size of a copy of a set", t, func() {
		s := newTestSet()
		assert.Equal(t, s.Size(), s.Copy().Size())
	})
}
