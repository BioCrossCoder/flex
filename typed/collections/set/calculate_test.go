package set

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	s1 := Of(1, 2, 3, 4)
	s2 := Of(3, 4, 5, 6)
	convey.Convey("s1 - s2", t, func() {
		assert.Equal(t, Of(1, 2), s1.Difference(s2))
	})
	convey.Convey("s2 - s1", t, func() {
		assert.Equal(t, Of(5, 6), s2.Difference(s1))
	})
	convey.Convey("(s1 - s2) | (s2 - s1)", t, func() {
		expected := Of(1, 2, 5, 6)
		assert.Equal(t, expected, s1.SymmetricDifference(s2))
		assert.Equal(t, expected, s2.SymmetricDifference(s1))
	})
	convey.Convey("s1 & s2", t, func() {
		expected := Of(3, 4)
		assert.Equal(t, expected, s1.Intersection(s2))
		assert.Equal(t, expected, s2.Intersection(s1))
	})
	convey.Convey("s1 | s2", t, func() {
		expected := Of(1, 2, 3, 4, 5, 6)
		assert.Equal(t, expected, s1.Union(s2))
		assert.Equal(t, expected, s2.Union(s1))
	})
}

func ExampleSet_Difference() {
	s1 := Of("a", "b", "c")
	s2 := Of("b", "c", "d")
	fmt.Println(s1.Difference(s2))
	fmt.Println(s2.Difference(s1))
	// Output:
	// {a}
	// {d}
}

func ExampleSet_Intersection() {
	s1 := Of("a", "b", "c")
	s2 := Of("e", "c", "d")
	fmt.Println(s1.Intersection(s2))
	// Output: {c}
}

func ExampleSet_Union() {
	s1 := Of("a", "b", "c")
	s2 := Of("b", "c", "d")
	s3 := Of("a", "b", "c", "d")
	union := s1.Union(s2)
	fmt.Println(union.Equal(s3))
	// Output: true
}

func ExampleSet_SymmetricDifference() {
	s1 := Of("a", "b", "c")
	s2 := Of("b", "c", "d")
	sd1 := s1.SymmetricDifference(s2)
	sd2 := s2.SymmetricDifference(s1)
	sd := Of("a", "d")
	fmt.Println(sd1.Equal(sd2))
	fmt.Println(sd1.Equal(sd))
	// Output:
	// true
	// true
}
