package dict

import (
	"fmt"
	"github.com/biocrosscoder/flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClear(t *testing.T) {
	convey.Convey("clear a dict", t, func() {
		d := newTestDict()
		assert.False(t, d.Empty())
		_ = d.Clear()
		assert.True(t, d.Empty())
	})
}

func TestSet(t *testing.T) {
	convey.Convey("set a value by key in a dict", t, func() {
		d := newTestDict()
		_ = d.Set("key1", "value1")
		assert.Equal(t, "value1", d.Get("key1"))
	})
}

func TestDelete(t *testing.T) {
	convey.Convey("delete a key in a dict", t, func() {
		d := newTestDict()
		_ = d.Set("key1", "value1")
		assert.True(t, d.Has("key1"))
		ok := d.Delete("key1")
		assert.True(t, ok)
		assert.False(t, d.Has("key1"))
	})
}

func TestPop(t *testing.T) {
	convey.Convey("pop a key from a dict", t, func() {
		d := newTestDict()
		_ = d.Set("key1", "value1")
		assert.True(t, d.Has("key1"))
		v, err := d.Pop("key1")
		assert.Nil(t, err)
		assert.Equal(t, "value1", v)
		assert.False(t, d.Has("key1"))
	})
	convey.Convey("pop a non-exist key from a dict", t, func() {
		d := newTestDict()
		_ = d.Delete("f")
		assert.False(t, d.Has("f"))
		v, err := d.Pop("f")
		assert.Equal(t, common.ErrKeyNotFound, err)
		assert.Nil(t, v)
	})
	convey.Convey("pop a non-exist key from a dict with default value", t, func() {
		d := newTestDict()
		_ = d.Delete("f")
		assert.False(t, d.Has("f"))
		v, err := d.Pop("f", "default")
		assert.Nil(t, err)
		assert.Equal(t, "default", v)
	})
}

func TestPopItem(t *testing.T) {
	convey.Convey("pop an item from a dict", t, func() {
		d := newTestDict()
		d2 := d.Copy()
		k, v, err := d.PopItem()
		assert.Nil(t, err)
		assert.Equal(t, d2.Get(k), v)
	})
}

func TestUpdate(t *testing.T) {
	convey.Convey("update a dict with another dict", t, func() {
		d := newTestDict()
		d2 := Dict{"key1": "value1", "key2": "value2"}
		_ = d.Update(d2)
		assert.Equal(t, d.Get("key1"), "value1")
		assert.Equal(t, d.Get("key2"), "value2")
	})
}

func ExampleDict_Clear() {
	d := Dict{"a": 1}
	fmt.Println(d)
	d.Clear()
	fmt.Println(d)
	fmt.Println(d.Size())
	fmt.Println(d.Empty())
	// Output:
	// map[a:1]
	// map[]
	// 0
	// true
}

func ExampleDict_Set() {
	d := make(Dict)
	fmt.Println(d.Get("a", nil))
	fmt.Println(d.Has("a"))
	d.Set("a", 5)
	fmt.Println(d.Has("a"))
	fmt.Println(d.Get("a", nil))
	d.Set("a", 10)
	fmt.Println(d.Get("a", nil))
	// Output:
	// <nil>
	// false
	// true
	// 5
	// 10
}

func ExampleDict_Delete() {
	d := Dict{"a": 1}
	fmt.Println(d.Has("a"))
	fmt.Println(d.Delete("a"))
	fmt.Println(d.Has("a"))
	// Output:
	// true
	// true
	// false
}

func ExampleDict_Pop() {
	d := Dict{"a": 1, "b": 2, "c": 3}
	fmt.Println(d.Has("a"))
	fmt.Println(d.Pop("a"))
	fmt.Println(d.Has("a"))
	fmt.Println(d.Pop("a"))
	fmt.Println(d.Pop("a", "default"))
	// Output:
	// true
	// 1 <nil>
	// false
	// <nil> the key is not found in the dict
	// default <nil>
}

func ExampleDict_Update() {
	d1 := Dict{"a": 1, "b": 2}
	d2 := Dict{"b": 3, "c": 4}
	fmt.Println(d1)
	fmt.Println(d2)
	d1.Update(d2)
	fmt.Println(d1)
	fmt.Println(d2)
	// Output:
	// map[a:1 b:2]
	// map[b:3 c:4]
	// map[a:1 b:3 c:4]
	// map[b:3 c:4]
}
