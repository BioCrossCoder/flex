package dict

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	convey.Convey("copy a dict", t, func() {
		d := newTestDict()
		assert.True(t, d.Equal(d.Copy()))
	})
}

func TestFromEntries(t *testing.T) {
	convey.Convey("create a dict from entries", t, func() {
		entries := [][2]any{{"a", 1}, {"b", 2}, {"c", 3}}
		d := FromEntries(entries...)
		assert.Equal(t, 3, d.Size())
		for _, pair := range entries {
			assert.Equal(t, pair[1], d.Get(pair[0]))
		}
	})
}
