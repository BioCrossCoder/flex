package functools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	f := func(a string) string {
		return a + a
	}
	convey.Convey("Call Map on [map]", t, func() {
		m := map[int]string{
			1: "one",
			2: "two",
		}
		result, err := MapResult(f, m)
		assert.Nil(t, err)
		output, isMap := result.(map[any]any)
		assert.True(t, isMap)
		assert.Equal(t, len(output), len(m))
		for k, v := range output {
			assert.Equal(t, f(m[k.(int)]), v)
		}
	})
	convey.Convey("Call Map on [slice]", t, func() {
		s := []string{"1", "2", "3"}
		result, err := MapResult(f, s)
		assert.Nil(t, err)
		output, isSlice := result.([]any)
		assert.True(t, isSlice)
		assert.Equal(t, len(output), len(s))
		assert.Equal(t, cap(output), cap(s))
		for i, v := range output {
			assert.Equal(t, f(s[i]), v)
		}
	})
	convey.Convey("Call Map on [array]", t, func() {
		s := [3]string{"1", "2", "3"}
		result, err := MapResult(f, s)
		assert.Nil(t, err)
		output, isSlice := result.([]any)
		assert.True(t, isSlice)
		assert.Equal(t, len(output), len(s))
		assert.Equal(t, cap(output), cap(s))
		for i, v := range output {
			assert.Equal(t, f(s[i]), v)
		}
	})
	convey.Convey("Call Map on [string]", t, func() {
		s := "hello"
		result, err := MapResult(f, s)
		assert.Nil(t, err)
		output, isSlice := result.([]any)
		assert.True(t, isSlice)
		assert.Equal(t, len(output), len([]rune(s)))
		for i, r := range s {
			assert.Equal(t, f(string(r)), output[i])
		}
	})
}
