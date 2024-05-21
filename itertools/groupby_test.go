package itertools

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupBy(t *testing.T) {
	convey.Convey("group by item", t, func() {
		entry := [][2]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}, {3, 6}, {3, 7}}
		g, err := GroupBy(entry, func(x any) any {
			return x.([2]int)[0]
		})
		assert.Nil(t, err)
		assert.Equal(t, []any{
			[]any{[2]int{1, 2}, [2]int{1, 3}},
			[]any{[2]int{2, 4}, [2]int{2, 5}},
			[]any{[2]int{3, 6}, [2]int{3, 7}},
		}, g.Pour())
	})
	convey.Convey("group by condition", t, func() {
		entry := []int{1, 2, 3, 4, 5, 6, 7}
		g, err := GroupBy(entry, func(x any) any {
			return x.(int)%2 == 0
		})
		assert.Nil(t, err)
		assert.Equal(t, []any{
			[]any{1, 3, 5, 7},
			[]any{2, 4, 6},
		}, g.Pour())
	})
}

func ExampleGroupBy() {
	entry := [][2]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}, {3, 6}, {3, 7}}
	f := func(x any) any {
		return x.([2]int)[0]
	}
	g, _ := GroupBy(entry, f)
	for g.Next() {
		fmt.Println(g.Value())
	}
	// Output:
	// [[1 2] [1 3]]
	// [[2 4] [2 5]]
	// [[3 6] [3 7]]
}
