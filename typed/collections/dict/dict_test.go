package dict

import (
	"fmt"
	"github.com/biocrosscoder/flex/typed/collections/arraylist"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestDict() Dict[string, int] {
	return Dict[string, int]{"a": 1, "b": 2, "c": 3}
}

func TestSize(t *testing.T) {
	convey.Convey("get size of a dict", t, func() {
		d := newTestDict()
		assert.Equal(t, len(d), d.Size())
	})
}

func TestGet(t *testing.T) {
	m := newTestDict()
	convey.Convey("get value by key from a dict", t, func() {
		assert.Equal(t, m["a"], m.Get("a"))
	})
	convey.Convey("get value by default of a key not exist in dict", t, func() {
		assert.Equal(t, 0, m.Get("d", 0))
	})
}

func TestCopy(t *testing.T) {
	convey.Convey("copy a dict", t, func() {
		d := newTestDict()
		assert.True(t, d.Equal(d.Copy()))
	})
}

func TestHas(t *testing.T) {
	convey.Convey("check the value matches the key in the dict", t, func() {
		d := newTestDict()
		for _, item := range d.Items() {
			assert.Equal(t, item.Value, d.Get(item.Key))
		}
	})
	convey.Convey("check if the values are included in the dict", t, func() {
		d := newTestDict()
		values := d.Values()
		for _, k := range d.Keys() {
			assert.True(t, d.Has(k))
			assert.True(t, arraylist.Of(values...).Includes(d.Get(k)))
		}
	})
}

func TestEmpty(t *testing.T) {
	convey.Convey("check if the dict is empty", t, func() {
		d := newTestDict()
		assert.False(t, d.Empty())
		_ = d.Clear()
		assert.True(t, d.Empty())
	})
}

func ExampleDict() {
	d := Dict[string, int]{"a": 1, "b": 2}
	fmt.Println(d)
	// Output: map[a:1 b:2]
}

func ExampleDict_Get() {
	d := Dict[string, int]{"one": 1, "two": 2, "three": 3}
	fmt.Println(d.Get("two"))
	fmt.Println(d.Get("four"))
	fmt.Println(d.Get("four", -1))
	// Output:
	// 2
	// 0
	// -1
}

func ExampleDict_Size() {
	d := Dict[string, int]{"one": 1, "two": 2, "three": 3}
	fmt.Println(d.Size())
	// Output: 3
}

func ExampleDict_Has() {
	d := Dict[string, string]{"key1": "value1", "key2": "value2"}
	fmt.Println(d.Has("key1"))
	fmt.Println(d.Has("key3"))
	// Output:
	// true
	// false
}

func ExampleDict_Empty() {
	d := Dict[string, string]{}
	fmt.Println(d.Empty())
	d.Set("key", "value")
	fmt.Println(d.Empty())
	// Output:
	// true
	// false
}

func ExampleDict_Equal() {
	d1 := Dict[string, string]{"key1": "value1", "key2": "value2"}
	d2 := Dict[string, string]{"key1": "value1", "key2": "value2"}
	d3 := Dict[string, string]{"key1": "value1", "key2": "value3"}
	d4 := Dict[string, string]{"key1": "value1"}
	fmt.Println(d1.Equal(d2))
	fmt.Println(d1.Equal(d3))
	fmt.Println(d1.Equal(d4))
	// Output:
	// true
	// false
	// false
}

func ExampleDict_Copy() {
	d := Dict[string, int]{"a": 1, "b": 2}
	backup := d.Copy()
	fmt.Println(d)
	fmt.Println(backup)
	// Output:
	// map[a:1 b:2]
	// map[a:1 b:2]
}
