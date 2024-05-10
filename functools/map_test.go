package functools

import (
	"flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	f := func(a string) string {
		return a + a
	}
	convey.Convey("Call Map on [slice]", t, func() {
		s := []string{"1", "2", "3"}
		result, err := Map(f, s)
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
		result, err := Map(f, s)
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
		result, err := Map(f, s)
		assert.Nil(t, err)
		output, isSlice := result.([]any)
		assert.True(t, isSlice)
		assert.Equal(t, len(output), len([]rune(s)))
		for i, r := range s {
			assert.Equal(t, f(string(r)), output[i])
		}
	})
}

func TestMaps(t *testing.T) {
	f := func(a, b, c int) int {
		return a + b + c
	}
	convey.Convey("Call Maps on 3 arrays", t, func() {
		result, err := Maps(f, []int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
		assert.Nil(t, err)
		assert.Equal(t, result, []any{12, 15, 18})
	})
	convey.Convey("not enough entries", t, func() {
		result, err := Maps(f, []int{1, 2, 3})
		assert.Equal(t, common.ErrIllegalParamCount, err)
		assert.Nil(t, result)
	})
	convey.Convey("invalid entry type", t, func() {
		result, err := Maps(f, 666)
		assert.Equal(t, common.ErrNotList, err)
		assert.Nil(t, result)
	})
	convey.Convey("entry length not equal", t, func() {
		result, err := Maps(f, []int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8})
		assert.Equal(t, common.ErrListLengthMismatch, err)
		assert.Nil(t, result)
	})
}
