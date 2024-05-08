package common

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestIsInputFuncValid(t *testing.T) {
	entry := func(a int) int {
		return a
	}
	convey.Convey("valid", t, func() {
		err := IsInputFuncValid(entry, 1, 1)
		assert.Nil(t, err)
	})
	convey.Convey("not a func", t, func() {
		err := IsInputFuncValid("func", 1, 1)
		assert.Equal(t, ErrNotFunc, err)
	})
	convey.Convey("invalid func", t, func() {
		err := IsInputFuncValid(entry, 0, 1)
		assert.Equal(t, ErrIllegalParamCount, err)
		err = IsInputFuncValid(entry, 1, 0)
		assert.Equal(t, ErrIllegalReturnCount, err)
	})
}

func TestIsJudgeFunc(t *testing.T) {
	convey.Convey("valid", t, func() {
		entry := func(a int) bool {
			return a > 0
		}
		err := IsJudgeFunc(entry)
		assert.Nil(t, err)
	})
	convey.Convey("invalid", t, func() {
		entry := func(a int) int {
			return a
		}
		err := IsJudgeFunc(entry)
		assert.NotNil(t, err)
	})
}

func TestIsList(t *testing.T) {
	convey.Convey("array", t, func() {
		entry := [3]int{3, 0, 6}
		err := IsList(entry)
		assert.Nil(t, err)
	})
	convey.Convey("slice", t, func() {
		entry := []int{6, 0, 2}
		err := IsList(entry)
		assert.Nil(t, err)
	})
	convey.Convey("not a list", t, func() {
		entry := "hello"
		err := IsList(entry)
		assert.Equal(t, ErrNotList, err)
	})
}

func TestIsSequence(t *testing.T) {
	convey.Convey("array", t, func() {
		entry := [3]int{1, 2, 3}
		err := IsSequence(entry)
		assert.Nil(t, err)
	})
	convey.Convey("slice", t, func() {
		entry := []int{}
		err := IsSequence(entry)
		assert.Nil(t, err)
	})
	convey.Convey("string", t, func() {
		entry := ""
		err := IsSequence(entry)
		assert.Nil(t, err)
	})
	convey.Convey("not a sequence", t, func() {
		entry := 6
		err := IsSequence(entry)
		assert.Equal(t, ErrNotSeq, err)
	})
}

func TestIsIterable(t *testing.T) {
	convey.Convey("slice", t, func() {
		err := IsIterable([]int{})
		assert.Nil(t, err)
	})
	convey.Convey("array", t, func() {
		err := IsIterable([3]int{1, 2, 3})
		assert.Nil(t, err)
	})
	convey.Convey("map", t, func() {
		err := IsIterable(map[string]int{})
		assert.Nil(t, err)
	})
	convey.Convey("string", t, func() {
		err := IsIterable("hello")
		assert.Nil(t, err)
	})
	convey.Convey("not iterable", t, func() {
		err := IsIterable(123)
		assert.Equal(t, ErrNotIterable, err)
	})
}

func TestCopyMap(t *testing.T) {
	convey.Convey("copy map", t, func() {
		entry := map[string]int{
			"one": 1,
			"two": 2,
		}
		actual := CopyMap(reflect.ValueOf(entry), len(entry))
		for k, v := range actual {
			assert.Equal(t, entry[k.(string)], v)
		}
	})
}

func TestCopyList(t *testing.T) {
	convey.Convey("copy array", t, func() {
		entry := [3]int{1, 2, 3}
		actual := CopyList(reflect.ValueOf(entry), len(entry))
		for i, v := range actual {
			assert.Equal(t, entry[i], v)
		}
	})
	convey.Convey("copy slice", t, func() {
		entry := []int{1, 2, 3}
		actual := CopyList(reflect.ValueOf(entry), len(entry))
		for i, v := range actual {
			assert.Equal(t, entry[i], v)
		}
	})
}

func TestConvertStringToList(t *testing.T) {
	convey.Convey("convert string to list", t, func() {
		entry := "hello"
		acutal := ConvertStringToList(entry)
		for i, r := range entry {
			assert.Equal(t, string(r), acutal[i])
		}
	})
}

func TestConvertMapToLists(t *testing.T) {
	convey.Convey("convert map to list", t, func() {
		entry := map[any]any{
			1: "one",
			2: "two",
		}
		keys, values, length := ConvertMapToLists(entry)
		assert.Equal(t, len(entry), length)
		for i := 0; i < length; i++ {
			assert.Equal(t, entry[keys[i]], values[i])
		}
	})
}

func TestEqual(t *testing.T) {
	convey.Convey("compare normal values", t, func() {
		assert.True(t, Equal(1, 1))
		assert.True(t, Equal(1.0, 1.0))
		assert.True(t, Equal("hello", "hello"))
		assert.True(t, Equal(true, true))
	})
	convey.Convey("compare special values", t, func() {
		assert.True(t, Equal(nil, nil))
		assert.True(t, Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	})
}

func TestIndex(t *testing.T) {
	convey.Convey("convert index out of range to be in range", t, func() {
		assert.Equal(t, ParseIndex(8, 3), 2)
		assert.Equal(t, ParseIndex(-2, 6), 4)
		assert.Equal(t, ParseIndex(-10, 3), 0)
	})
	convey.Convey("verify invalid range", t, func() {
		assert.Equal(t, CheckRange(1, 3, 0, 6), ErrZeroStep)
		assert.Equal(t, CheckRange(1, 3, -1, 6), ErrInvalidRange)
		assert.Equal(t, CheckRange(10, 12, 1, 6), ErrOutOfRange)
	})
}

type mockDict map[any]any

func (m mockDict) Size() int {
	return len(m)
}

type mockList []any

func (m mockList) Len() int {
	return len(m)
}

func TestLen(t *testing.T) {
	convey.Convey("get length", t, func() {
		assert.Equal(t, Len([]int{1, 2, 3}), 3)
		assert.Equal(t, Len("hello"), 5)
		assert.Equal(t, Len(map[string]int{"one": 1, "two": 2}), 2)
		assert.Equal(t, Len(123), -1)
		assert.Equal(t, Len(nil), -1)
		assert.Equal(t, Len(&mockDict{"one": 1, "two": 2}), 2)
		assert.Equal(t, Len(&mockList{1, 2, 3}), 3)
	})
}
