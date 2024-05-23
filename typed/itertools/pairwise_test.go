package itertools

import (
	"fmt"
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

func ExamplePairWise() {
	entry := []int{1, 2, 3, 4, 5}
	iter := PairWise(entry)
	for iter.Next() {
		fmt.Println(iter.Value())
	}
	// Output:
	// [1 2]
	// [2 3]
	// [3 4]
	// [4 5]
}
