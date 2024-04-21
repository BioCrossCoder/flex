package itertools

import (
	"flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPairWise(t *testing.T) {
	convey.Convey("pair array", t, func() {
		entry := []int{1, 2, 3}
		iter, err := PairWise(entry)
		assert.Nil(t, err)
		assert.Equal(t, []any{
			[]any{1, 2},
			[]any{2, 3},
		}, iter.Pour())
	})
	convey.Convey("pair string", t, func() {
		entry := "hello"
		iter, err := PairWise(entry)
		assert.Nil(t, err)
		assert.Equal(t, []any{
			[]any{"h", "e"},
			[]any{"e", "l"},
			[]any{"l", "l"},
			[]any{"l", "o"},
		}, iter.Pour())
	})
	convey.Convey("not a sequence", t, func() {
		entry := 123
		iter, err := PairWise(entry)
		assert.Equal(t, common.ErrNotSeq, err)
		assert.Nil(t, iter)
	})
}
