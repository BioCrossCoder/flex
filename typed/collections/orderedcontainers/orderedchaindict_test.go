package orderedcontainers

import (
	"encoding/json"
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedChainDict(t *testing.T) {
	pairs := [][2]any{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{"d", 4},
	}
	d := NewOrderedChainDict[string, int]()
	for _, pair := range pairs {
		_ = d.Set(pair[0].(string), pair[1].(int))
	}
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
		_ = d2.Update(*d)
		assert.True(t, d.Equal(d2))
	})
	convey.Convey("jsonify and stringify", t, func() {
		d2 := d.Copy()
		data, err := json.Marshal(d2)
		assert.Nil(t, err)
		d3 := NewOrderedChainDict[string, int]()
		assert.Nil(t, json.Unmarshal(data, d3))
		assert.Equal(t, fmt.Sprint(d2), fmt.Sprint(d3))
	})
}

func ExampleOrderedChainDict() {
	d := NewOrderedChainDict[string, int]()
	d.Set("c", 1)
	d.Set("a", 2)
	d.Set("b", 3)
	fmt.Println(d)
	// Output: map[c:1 a:2 b:3]
}

func ExampleOrderedChainDict_Set() {
	d := NewOrderedChainDict[string, int]()
	d.Set("c", 1)
	d.Set("a", 2)
	fmt.Println(d)
	d.Set("c", 3)
	fmt.Println(d)
	// Output:
	// map[c:1 a:2]
	// map[c:3 a:2]
}

func ExampleOrderedChainDict_Delete() {
	d := NewOrderedChainDict[string, int]()
	d.Set("c", 1)
	d.Set("a", 2)
	d.Set("b", 3)
	fmt.Println(d)
	d.Delete("a")
	fmt.Println(d)
	// Output:
	// map[c:1 a:2 b:3]
	// map[c:1 b:3]
}

func ExampleOrderedChainDict_PopItem() {
	d := NewOrderedChainDict[string, int]()
	d.Set("c", 1)
	d.Set("a", 2)
	d.Set("b", 3)
	fmt.Println(d)
	key, value, _ := d.PopItem()
	fmt.Println(key, value)
	fmt.Println(d)
	// Output:
	// map[c:1 a:2 b:3]
	// b 3
	// map[c:1 a:2]
}

func ExampleOrderedChainDict_Update() {
	d1 := NewOrderedChainDict[string, int]()
	d1.Set("c", 1)
	d1.Set("a", 2)
	d1.Set("b", 3)
	d2 := NewOrderedChainDict[string, int]()
	d2.Set("d", 5)
	d2.Set("b", 4)
	d2.Set("e", 6)
	fmt.Println(d1)
	d1.Update(*d2)
	fmt.Println(d1)
	// Output:
	// map[c:1 a:2 b:3]
	// map[c:1 a:2 b:4 d:5 e:6]
}

func ExampleOrderedChainDict_Keys() {
	d := NewOrderedChainDict[string, int]()
	d.Set("c", 1)
	d.Set("a", 2)
	d.Set("b", 3)
	fmt.Println(d.Keys())
	// Output: [c a b]
}

func ExampleOrderedChainDict_Values() {
	d := NewOrderedChainDict[string, int]()
	d.Set("c", 1)
	d.Set("a", 2)
	d.Set("b", 3)
	fmt.Println(d.Values())
	// Output: [1 2 3]
}

func ExampleOrderedChainDict_Items() {
	d := NewOrderedChainDict[string, int]()
	d.Set("c", 1)
	d.Set("a", 2)
	d.Set("b", 3)
	for _, item := range d.Items() {
		fmt.Println(*item)
	}
	// Output:
	// {c 1}
	// {a 2}
	// {b 3}
}

func ExampleOrderedChainDict_Copy() {
	d := NewOrderedChainDict[string, int]()
	d.Set("c", 1)
	d.Set("a", 2)
	d.Set("b", 3)
	fmt.Println(d)
	d2 := d.Copy()
	fmt.Println(d2)
	// Output:
	// map[c:1 a:2 b:3]
	// map[c:1 a:2 b:3]
}

func ExampleOrderedChainDict_Equal() {
	d1 := NewOrderedChainDict[string, int]()
	d1.Set("c", 1)
	d1.Set("a", 2)
	d1.Set("b", 3)
	d2 := NewOrderedChainDict[string, int]()
	d2.Set("c", 1)
	d2.Set("b", 3)
	d2.Set("a", 2)
	d3 := NewOrderedChainDict[string, int]()
	d3.Set("c", 1)
	d3.Set("a", 2)
	d3.Set("b", 3)
	fmt.Println(d1.Equal(*d2))
	fmt.Println(d1.Equal(*d3))
	// Output:
	// false
	// true
}

func ExampleOrderedChainDict_KeyAt() {
	d := NewOrderedChainDict[string, int]()
	d.Set("c", 1)
	d.Set("a", 2)
	d.Set("b", 3)
	key, _ := d.KeyAt(1)
	fmt.Println(key)
	// Output: a
}

func ExampleOrderedChainDict_IndexOf() {
	d := NewOrderedChainDict[string, int]()
	d.Set("c", 1)
	d.Set("a", 2)
	d.Set("b", 3)
	index := d.IndexOf("a")
	fmt.Println(index)
	// Output: 1
}

func ExampleOrderedChainDict_MarshalJSON() {
	d := NewOrderedChainDict[string, int]()
	d.Set("c", 1)
	d.Set("a", 2)
	d.Set("b", 3)
	data, _ := json.Marshal(d)
	fmt.Println(string(data))
	// Output: [["c",1],["a",2],["b",3]]
}

func ExampleOrderedChainDict_UnmarshalJSON() {
	d := NewOrderedChainDict[string, int]()
	_ = json.Unmarshal([]byte(`[["c",1],["a",2],["b",3]]`), &d)
	fmt.Println(d)
	// Output: map[c:1 a:2 b:3]
}
