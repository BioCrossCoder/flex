package itertools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPairWise(t *testing.T) {
	convey.Convey("pair array", t, func() {
		entry := []int{1, 2, 3}
		iter := PairWise(entry)
		assert.Equal(t, [][2]int{{1, 2}, {2, 3}}, iter.Pour())
	})
}
