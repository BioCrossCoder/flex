package dict

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	convey.Convey("copy a dict", t, func() {
		d := newTestDict()
		assert.True(t, d.Equal(d.Copy()))
	})
}

func TestFromEntries(t *testing.T) {
	convey.Convey("create a dict from entries", t, func() {
		entries := [][2]any{{"a", 1}, {"b", 2}, {"c", 3}}
		d := FromEntries(entries...)
		assert.Equal(t, 3, d.Size())
		for _, pair := range entries {
			assert.Equal(t, pair[1], d.Get(pair[0]))
		}
	})
}

func ExampleDict_Copy() {
	d := Dict{"a": 1, "b": 2}
	backup := d.Copy()
	fmt.Println(d)
	fmt.Println(backup)
	// Output:
	// map[a:1 b:2]
	// map[a:1 b:2]
}

func ExampleFromEntries() {
	d := FromEntries([2]any{"a", 1}, [2]any{"b", 2})
	fmt.Println(d)
	// Output: map[a:1 b:2]
}
