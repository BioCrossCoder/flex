package collections

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestNamedList(t *testing.T) {
	convey.Convey("create a named list", t, func() {
		names := []string{"a", "b", "c"}
		nl := NewNamedList(names)
		assert.Equal(t, 3, nl.Len())
		assert.Equal(t, 3, nl.Elements().Len())
		assert.Equal(t, nl.Fields(), names)
		for _, name := range names {
			assert.True(t, nl.Contains(name))
		}
		convey.Convey("modify fields", func() {
			assert.False(t, nl.Add("a", 1))
			assert.Nil(t, nl.SetByIndex(0, 666))
			zero, err := nl.GetByIndex(0)
			assert.Nil(t, err)
			a, err := nl.GetByName("a")
			assert.Nil(t, err)
			assert.Equal(t, zero, a)
			assert.False(t, nl.Contains("d"))
			assert.Equal(t, -1, nl.Index("d"))
			assert.True(t, nl.Add("d", 1))
			assert.True(t, nl.Contains("d"))
			d, err := nl.GetByName("d")
			assert.Nil(t, err)
			assert.Equal(t, 1, d)
			assert.Nil(t, nl.SetByName("d", 2))
			d, err = nl.GetByIndex(nl.Index("d"))
			assert.Nil(t, err)
			assert.Equal(t, 2, d)
			assert.True(t, nl.Remove("a"))
			assert.False(t, nl.Contains("a"))
			assert.False(t, nl.Remove("a"))
			for i := 0; i < nl.Len(); i++ {
				assert.Nil(t, nl.SetByIndex(i, 6))
			}
			assert.Equal(t, nl.Len(), nl.Count(6))
		})
		convey.Convey("operate the list", func() {
			nl2 := nl.Copy()
			assert.True(t, nl2.Equal(*nl))
			assert.False(t, nl2.Empty())
			assert.Equal(t, nl2.Clear(), &nl2)
			assert.True(t, nl2.Empty())
			nl3 := nl2.With("f", 50)
			assert.True(t, nl3.Contains("f"))
			f, err := nl3.GetByName("f")
			assert.Nil(t, err)
			assert.Equal(t, 50, f)
			assert.False(t, nl2.Equal(nl3))
			assert.Equal(t, nl2.Update(nl3), &nl2)
			assert.True(t, nl2.Equal(nl3))
		})
	})
}
