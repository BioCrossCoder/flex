package dict

import (
	"fmt"
	"github.com/biocrosscoder/flex/collections/arraylist"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHas(t *testing.T) {
	convey.Convey("check the value matches the key in the dict", t, func() {
		d := newTestDict()
		for _, item := range d.Items() {
			assert.Equal(t, item[1], d.Get(item[0]))
		}
	})
	convey.Convey("check if the values are included in the dict", t, func() {
		d := newTestDict()
		values := d.Values()
		for _, k := range d.Keys() {
			assert.True(t, d.Has(k))
			assert.True(t, arraylist.ArrayList(values).Includes(d.Get(k)))
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

func ExampleDict_Has() {
	d := Dict{"key1": "value1", "key2": "value2"}
	fmt.Println(d.Has("key1"))
	fmt.Println(d.Has("key3"))
	// Output:
	// true
	// false
}

func ExampleDict_Empty() {
	d := Dict{}
	fmt.Println(d.Empty())
	d.Set("key", "value")
	fmt.Println(d.Empty())
	// Output:
	// true
	// false
}

func ExampleDict_Equal() {
	d1 := Dict{"key1": "value1", "key2": "value2"}
	d2 := Dict{"key1": "value1", "key2": "value2"}
	d3 := Dict{"key1": "value1", "key2": "value3"}
	d4 := Dict{"key1": "value1"}
	fmt.Println(d1.Equal(d2))
	fmt.Println(d1.Equal(d3))
	fmt.Println(d1.Equal(d4))
	// Output:
	// true
	// false
	// false
}
