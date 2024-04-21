package itertools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	convey.Convey("repeat", t, func() {
		entry := "233"
		r := Repeat(entry)
		assert.Equal(t, []any{"233", "233", "233"}, r.Repeat(3))
	})
}
