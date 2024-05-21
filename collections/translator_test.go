package collections

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTranslator(t *testing.T) {
	convey.Convey("translator test", t, func() {
		trans := NewTranslator([2]string{"a", "A"}, [2]string{"b", "B"}, [2]string{"c", "C"})
		convey.Convey("overload methods from dict", func() {
			trans2 := trans.Copy()
			assert.Equal(t, *trans, trans2)
			assert.True(t, trans2.Has("a"))
			assert.Equal(t, trans2.Get("a"), "A")
			assert.True(t, trans2.Delete("a"))
			assert.False(t, trans2.Has("a"))
			b, err := trans2.Pop("b")
			assert.Nil(t, err)
			assert.Equal(t, b, "B")
			assert.False(t, trans2.Has("b"))
			assert.True(t, trans2.Has("c"))
			k, v, err := trans2.PopItem()
			assert.Nil(t, err)
			assert.Equal(t, k, "c")
			assert.Equal(t, v, "C")
			assert.False(t, trans2.Has("c"))
			assert.False(t, trans2.Has("d"))
			assert.Equal(t, trans2.Get("d"), "")
			_ = trans2.Set("d", "D")
			assert.True(t, trans2.Has("d"))
			assert.Equal(t, trans2.Get("d"), "D")
			keys := trans2.Keys()
			values := trans2.Values()
			items := trans2.Items()
			for i, item := range items {
				assert.Equal(t, keys[i], item[0])
				assert.Equal(t, values[i], item[1])
			}
		})
		convey.Convey("translate a string", func() {
			trans2 := trans.Copy()
			assert.Equal(t, trans2.Translate("abcdefg"), "ABCdefg")
			_ = trans2.Update(*NewTranslator([2]string{"d", "D"}, [2]string{"e", "E"}))
			assert.Equal(t, trans2.Translate("abcdefg"), "ABCDEfg")
			assert.False(t, trans2.Empty())
			_ = trans2.Clear()
			assert.True(t, trans2.Empty())
			assert.Equal(t, trans2.Translate("abcdefg"), "abcdefg")
		})
	})
}

func ExampleTranslator() {
	trans := NewTranslator([2]string{"a", "A"}, [2]string{"b", "B"}, [2]string{"c", "C"})
	fmt.Println(trans.Translate("abcdefg"))
	// Output: ABCdefg
}
