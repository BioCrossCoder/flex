package dict

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestDict() Dict {
	return Dict{"a": 1, "b": 2, "c": 3}
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

func ExampleDict() {
	m := map[any]any{"a": 1, "b": 2}
	d := Dict(m)
	fmt.Println(d)
	// Output: map[a:1 b:2]
}

func ExampleDict_Get() {
	d := Dict{"one": 1, "two": 2, "three": 3}
	fmt.Println(d.Get("two"))
	fmt.Println(d.Get("four"))
	fmt.Println(d.Get("four", "not found"))
	// Output:
	// 2
	// <nil>
	// not found
}

func ExampleDict_Size() {
	d := Dict{"one": 1, "two": 2, "three": 3}
	fmt.Println(d.Size())
	// Output: 3
}
