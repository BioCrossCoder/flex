package itertools

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
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
			[2]any{1, 2},
			[2]any{2, 3},
		}, iter.Pour())
	})
	convey.Convey("pair string", t, func() {
		entry := "hello"
		iter, err := PairWise(entry)
		assert.Nil(t, err)
		assert.Equal(t, []any{
			[2]any{"h", "e"},
			[2]any{"e", "l"},
			[2]any{"l", "l"},
			[2]any{"l", "o"},
		}, iter.Pour())
	})
	convey.Convey("not a sequence", t, func() {
		entry := 123
		iter, err := PairWise(entry)
		assert.Equal(t, common.ErrNotSeq, err)
		assert.Nil(t, iter)
	})
}

func ExamplePairWise() {
	entry := []int{1, 2, 3, 4, 5}
	iter, _ := PairWise(entry)
	for iter.Next() {
		fmt.Println(iter.Value())
	}
	// Output:
	// [1 2]
	// [2 3]
	// [3 4]
	// [4 5]
}
