package orderedcontainers

import (
	"encoding/json"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedDict(t *testing.T) {
	pairs := [][2]any{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{"d", 4},
	}
	d := NewOrderedDict(pairs...)
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
			assert.Equal(t, [2]any{keys[i], values[i]}, items[i])
		}
	})
	convey.Convey("clear and update a dict", t, func() {
		d2 := d.Copy()
		assert.True(t, d.Equal(d2))
		assert.False(t, d2.Empty())
		_ = d2.Clear()
		assert.True(t, d2.Empty())
		_ = d2.Update(*d)
		assert.True(t, d.Equal(d2))
	})
	convey.Convey("jsonify and stringify", t, func() {
		d2 := d.Copy()
		data, err := json.Marshal(d2)
		assert.Nil(t, err)
		d3 := NewOrderedDict()
		assert.Nil(t, json.Unmarshal(data, d3))
		assert.Equal(t, fmt.Sprint(d2), fmt.Sprint(d3))
	})
}

func ExampleOrderedDict() {
	d := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	fmt.Println(d)
	// Output: map[c:1 a:2 b:3]
}

func ExampleOrderedDict_Set() {
	d := NewOrderedDict()
	d.Set("c", 1)
	d.Set("a", 2)
	fmt.Println(d)
	d.Set("c", 3)
	fmt.Println(d)
	// Output:
	// map[c:1 a:2]
	// map[c:3 a:2]
}

func ExampleOrderedDict_Delete() {
	d := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	fmt.Println(d)
	d.Delete("a")
	fmt.Println(d)
	// Output:
	// map[c:1 a:2 b:3]
	// map[c:1 b:3]
}

func ExampleOrderedDict_PopItem() {
	d := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	fmt.Println(d)
	key, value, _ := d.PopItem()
	fmt.Println(key, value)
	fmt.Println(d)
	// Output:
	// map[c:1 a:2 b:3]
	// b 3
	// map[c:1 a:2]
}

func ExampleOrderedDict_Update() {
	d1 := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	d2 := NewOrderedDict([2]any{"d", 5}, [2]any{"b", 4}, [2]any{"e", 6})
	fmt.Println(d1)
	d1.Update(*d2)
	fmt.Println(d1)
	// Output:
	// map[c:1 a:2 b:3]
	// map[c:1 a:2 b:4 d:5 e:6]
}

func ExampleOrderedDict_Keys() {
	d := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	fmt.Println(d.Keys())
	// Output: [c a b]
}

func ExampleOrderedDict_Values() {
	d := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	fmt.Println(d.Values())
	// Output: [1 2 3]
}

func ExampleOrderedDict_Items() {
	d := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	fmt.Println(d.Items())
	// Output: [[c 1] [a 2] [b 3]]
}

func ExampleOrderedDict_Copy() {
	d := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	fmt.Println(d)
	d2 := d.Copy()
	fmt.Println(d2)
	// Output:
	// map[c:1 a:2 b:3]
	// map[c:1 a:2 b:3]
}

func ExampleOrderedDict_Equal() {
	d1 := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	d2 := NewOrderedDict([2]any{"c", 1}, [2]any{"b", 3}, [2]any{"a", 2})
	d3 := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	fmt.Println(d1.Equal(*d2))
	fmt.Println(d1.Equal(*d3))
	// Output:
	// false
	// true
}

func ExampleOrderedDict_KeyAt() {
	d := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	key, _ := d.KeyAt(1)
	fmt.Println(key)
	// Output: a
}

func ExampleOrderedDict_IndexOf() {
	d := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	index := d.IndexOf("a")
	fmt.Println(index)
	// Output: 1
}

func ExampleOrderedDict_MarshalJSON() {
	d := NewOrderedDict([2]any{"c", 1}, [2]any{"a", 2}, [2]any{"b", 3})
	data, _ := json.Marshal(d)
	fmt.Println(string(data))
	// Output: [["c",1],["a",2],["b",3]]
}

func ExampleOrderedDict_UnmarshalJSON() {
	d := NewOrderedDict()
	_ = json.Unmarshal([]byte(`[["c",1],["a",2],["b",3]]`), &d)
	fmt.Println(d)
	// Output: map[c:1 a:2 b:3]
}
