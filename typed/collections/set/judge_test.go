package set

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHas(t *testing.T) {
	convey.Convey("check if a set has the specified element", t, func() {
		s := newTestSet()
		_ = s.Discard(1)
		assert.False(t, s.Has(1))
		_ = s.Add(1)
		assert.True(t, s.Has(1))
	})
}

func TestEmpty(t *testing.T) {
	convey.Convey("check if a set is empty", t, func() {
		s := newTestSet()
		_ = s.Add(1)
		assert.False(t, s.Empty())
		_ = s.Clear()
		assert.True(t, s.Empty())
	})
}

func TestSetRelation(t *testing.T) {
	convey.Convey("check if two sets has common elements", t, func() {
		s1 := Of(1, 2)
		s2 := Of(2, 3)
		s3 := Of(3, 4)
		assert.False(t, s1.IsDisjoint(s2))
		assert.False(t, s2.IsDisjoint(s1))
		assert.False(t, s2.IsDisjoint(s3))
		assert.False(t, s3.IsDisjoint(s2))
		assert.True(t, s1.IsDisjoint(s3))
		assert.True(t, s3.IsDisjoint(s1))
	})
	convey.Convey("check if one set is a subset of another set", t, func() {
		s1 := Of(1, 2)
		s2 := Of(1, 3)
		s3 := Of(1, 2, 3)
		assert.True(t, s1.IsSubset(s3))
		assert.True(t, s2.IsSubset(s3))
		assert.False(t, s1.IsSubset(s2))
		assert.False(t, s2.IsSubset(s1))
		assert.False(t, s3.IsSubset(s1))
		assert.False(t, s3.IsSubset(s2))
	})
	convey.Convey("check if one set is a superset of another set", t, func() {
		s1 := Of(1, 2)
		s2 := Of(1, 3)
		s3 := Of(1, 2, 3)
		assert.False(t, s1.IsSuperset(s2))
		assert.False(t, s2.IsSuperset(s1))
		assert.True(t, s3.IsSuperset(s1))
		assert.True(t, s3.IsSuperset(s2))
		assert.False(t, s1.IsSuperset(s3))
		assert.False(t, s2.IsSuperset(s3))
	})
}
