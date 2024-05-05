package collections

import (
	"flex/typed/collections/dict"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultDict(t *testing.T) {
	convey.Convey("get default value by non-exist key", t, func() {
		d := dict.Dict[string, int]{}
		df := 6
		dd := NewDefaultDict(d, df)
		assert.False(t, dd.Has("a"))
		assert.Equal(t, df, dd.Pop("a"))
		assert.False(t, dd.Has("a"))
		assert.Equal(t, df, dd.Get("a"))
		assert.True(t, dd.Has("a"))
		dd2 := dd.Copy()
		df2 := -1
		assert.Equal(t, dd2.SetDefault(df2), &dd2)
		assert.True(t, dd2.Has("a"))
		assert.Equal(t, df, dd2.Get("a"))
		assert.Equal(t, df, dd2.Pop("a"))
		assert.False(t, dd2.Has("a"))
		assert.Equal(t, df2, dd2.Get("a"))
		assert.True(t, dd2.Has("a"))
		assert.Equal(t, df2, dd2.Pop("a"))
		assert.False(t, dd2.Has("a"))
	})
	convey.Convey("inherit methods of dict", t, func() {
		d := dict.Dict[string, int]{"a": 2}
		df := 6
		dd := NewDefaultDict(d, df)
		dd2 := dd.Copy()
		assert.True(t, dd.Equal(dd2))
		assert.Equal(t, dd.Set("a", 1), dd)
		assert.Equal(t, dd.Update(dd2), dd)
		assert.Equal(t, dd.Get("a"), dd2.Get("a"))
		assert.NotZero(t, dd.Size())
		assert.Equal(t, dd.Clear(), dd)
		assert.Zero(t, dd.Size())
	})
}
