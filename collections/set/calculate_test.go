package set

import (
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
