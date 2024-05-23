package set

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestSet() Set[int] {
	return Of([]int{1, 2, 3}...)
}

func TestSet(t *testing.T) {
	convey.Convey("get size of a copy of a set", t, func() {
		s := newTestSet()
		assert.True(t, s.Equal(s.Copy()))
	})
}

func ExampleSet() {
	s := Of(1, 2, 3, 4, 5, 5, 4, 3, 2, 1)
	fmt.Println(s.Size())
	// Output: 5
}

func ExampleSet_Size() {
	s := Of(1, 2, 3, 4, 5)
	fmt.Println(s.Size())
	// Output: 5
}

func ExampleSet_Copy() {
	s1 := Of("apple", "banana", "orange")
	s2 := s1.Copy()
	fmt.Println(s1.Equal(s2))
	// Output: true
}
