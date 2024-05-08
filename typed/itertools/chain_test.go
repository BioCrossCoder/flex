package itertools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChain(t *testing.T) {
	convey.Convey("link iterable entries together", t, func() {
		entry1 := []int{1, 2, 3}
		entry2 := []int{4, 5, 6}
		entry3 := []int{7, 8, 9}
		iterator := Chain(entry1, entry2, entry3)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, iterator.Pour())
	})
}
