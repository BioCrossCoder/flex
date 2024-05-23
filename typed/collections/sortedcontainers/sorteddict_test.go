package sortedcontainers

import (
	"fmt"
	"github.com/biocrosscoder/flex/typed/collections/dict"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedDict(t *testing.T) {
	src := dict.Dict[string, int]{
		"d": 1,
		"a": 2,
		"b": 3,
		"c": 4,
	}
	d := NewSortedDict[string, int](nil, nil)
	for k, v := range src {
		_ = d.Set(k, v)
	}
	assert.Equal(t, d.Keys(), []string{"a", "b", "c", "d"})
	size := d.Size()
	convey.Convey("add a key-value pair to dict", t, func() {
		d2 := d.Copy()
		assert.Equal(t, d2.Size(), size)
		assert.False(t, d2.Has("e"))
		_ = d2.Set("e", 5)
		assert.True(t, d2.Has("e"))
		assert.Equal(t, d2.Size(), size+1)
		assert.Equal(t, d2.Get("e"), 5)
		assert.Equal(t, d2.IndexOf("e"), size)
		key, err := d2.KeyAt(size)
		assert.Nil(t, err)
		assert.Equal(t, key, "e")
	})
	convey.Convey("remove a key from dict", t, func() {
		d2 := d.Copy()
		assert.Equal(t, d2.Size(), size)
		assert.True(t, d2.Has("b"))
		assert.True(t, d2.Delete("b"))
		assert.False(t, d2.Has("b"))
		assert.Equal(t, d2.Size(), size-1)
	})
	convey.Convey("pop a key from dict", t, func() {
		d2 := d.Copy()
		assert.Equal(t, d2.Size(), size)
		assert.True(t, d2.Has("a"))
		expected := d2.Get("a")
		value, err := d2.Pop("a")
		assert.Nil(t, err)
		assert.Equal(t, expected, value)
		assert.Equal(t, d2.Size(), size-1)
		assert.False(t, d2.Has("a"))
	})
	convey.Convey("pop a key-value pair from dict", t, func() {
		d2 := d.Copy()
		assert.Equal(t, d2.Size(), size)
		expectedKey := d2.Keys()[d2.Size()-1]
		expectedValue := d2.Get(expectedKey)
		key, value, err := d2.PopItem()
		assert.Nil(t, err)
		assert.Equal(t, key, expectedKey)
		assert.Equal(t, value, expectedValue)
	})
	convey.Convey("iterate over dict", t, func() {
		d2 := d.Copy()
		keys := d2.Keys()
		values := d2.Values()
		items := d2.Items()
		for i := 0; i < d2.Size(); i++ {
			assert.Equal(t, keys[i], items[i].Key)
			assert.Equal(t, values[i], items[i].Value)
		}
	})
	convey.Convey("clear and update a dict", t, func() {
		d2 := d.Copy()
		assert.True(t, d.Equal(d2))
		assert.False(t, d2.Empty())
		_ = d2.Clear()
		assert.True(t, d2.Empty())
		_ = d2.Update(d.Dict)
		assert.True(t, d.Equal(d2))
	})
}

func ExampleSortedDict() {
	d := NewSortedDict(nil, dict.Dict[string, int]{"c": 1, "a": 2, "b": 3, "d": 4})
	fmt.Println(d.Keys())
	fmt.Println(d.Values())
	for _, item := range d.Items() {
		fmt.Println(*item)
	}
	// Output:
	// [a b c d]
	// [2 3 1 4]
	// {a 2}
	// {b 3}
	// {c 1}
	// {d 4}
}

func ExampleSortedDict_Set() {
	d := NewSortedDict[string, int](nil, nil)
	d.Set("c", 1)
	d.Set("a", 2)
	fmt.Println(d.Keys())
	fmt.Println(d.Values())
	d.Set("c", 3)
	fmt.Println(d.Keys())
	fmt.Println(d.Values())
	// Output:
	// [a c]
	// [2 1]
	// [a c]
	// [2 3]
}

func ExampleSortedDict_Update() {
	d1 := NewSortedDict[string, int](nil, dict.Dict[string, int]{"c": 1, "a": 2, "b": 3})
	fmt.Println(d1.Keys())
	fmt.Println(d1.Values())
	d2 := dict.Dict[string, int]{"d": 5, "b": 4, "e": 6}
	d1.Update(d2)
	fmt.Println(d1.Keys())
	fmt.Println(d1.Values())
	// Output:
	// [a b c]
	// [2 3 1]
	// [a b c d e]
	// [2 4 1 5 6]
}
