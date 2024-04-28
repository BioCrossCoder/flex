package dict

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestDict() Dict {
	return Dict{"a": 1, "b": 2, "c": 3}
}

func TestSize(t *testing.T) {
	convey.Convey("get size of a dict", t, func() {
		d := newTestDict()
		assert.Equal(t, len(d), d.Size())
	})
}

func TestGet(t *testing.T) {
	m := newTestDict()
	convey.Convey("get value by key from a dict", t, func() {
		assert.Equal(t, m["a"], m.Get("a"))
	})
	convey.Convey("get value by default of a key not exist in dict", t, func() {
		assert.Equal(t, 0, m.Get("d", 0))
	})
}
