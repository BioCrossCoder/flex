package collections

import (
	"fmt"
	"github.com/biocrosscoder/flex/typed/collections/arraylist"
	"github.com/biocrosscoder/flex/typed/collections/set"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCounter(t *testing.T) {
	l := arraylist.ArrayList[int]{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	convey.Convey("count elements in a list", t, func() {
		c := NewCounter(l, -1)
		nums := c.Elements()
		assert.Equal(t, set.Of(nums...), set.Of(l...))
		assert.Equal(t, c.Total(), len(l))
		maxCount := 0
		minCount := math.MaxInt
		for _, num := range nums {
			count := c.Get(num)
			if count > maxCount {
				maxCount = count
			}
			if count < minCount {
				minCount = count
			}
			assert.Equal(t, count, l.Count(num))
		}
		for _, num := range c.MostCommon() {
			assert.Equal(t, c.Get(num), maxCount)
		}
		for _, num := range c.LeastCommon() {
			assert.Equal(t, c.Get(num), minCount)
		}
	})
	convey.Convey("modify a counter", t, func() {
		c := NewCounter(l, -1)
		assert.Equal(t, c.Get(0), -1)
		_ = c.Set(0, 10)
		assert.Equal(t, c.Get(0), 10)
		_ = c.Increment(0, 2)
		assert.Equal(t, c.Get(0), 12)
		_ = c.Subtract(0, 1)
		assert.Equal(t, c.Get(0), 11)
		assert.True(t, c.Remove(0))
		assert.Equal(t, c.Get(0), -1)
		_ = c.SetDefault(6)
		assert.Equal(t, c.Get(0), 6)
	})
	convey.Convey("reset a counter", t, func() {
		c := NewCounter(l, -1)
		nums := c.Elements()
		rest := arraylist.ArrayList[int]{}
		for i, num := range nums {
			assert.Equal(t, c.Get(num), l.Count(num))
			if i%2 == 0 {
				assert.True(t, c.Remove(num))
				assert.Equal(t, c.Get(num), -1)
			} else {
				_ = rest.Push(num)
			}
		}
		_ = c.Reset()
		for _, num := range rest {
			assert.Equal(t, c.Get(num), -1)
			assert.True(t, c.Remove(num))
		}
		assert.Equal(t, len(c.Clear().Elements()), 0)
	})
	convey.Convey("merge several counters", t, func() {
		c1 := NewCounter(l, -1)
		c2 := c1.Copy()
		assert.True(t, c1.Equal(c2))
		mc := MergeCounts(c1, &c2)
		for _, element := range c2.Elements() {
			assert.Equal(t, mc.Get(element), c1.Get(element)+c2.Get(element))
		}
	})
}

func ExampleCounter() {
	l := arraylist.ArrayList[int]{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	c := NewCounter(l, -1)
	for i := 0; i < 10; i++ {
		fmt.Println(i, c.Get(i))
	}
	// Output:
	// 	0 -1
	// 1 2
	// 2 1
	// 3 2
	// 4 1
	// 5 3
	// 6 1
	// 7 -1
	// 8 -1
	// 9 1
}

func ExampleCounter_MostCommon() {
	l := arraylist.ArrayList[int]{1, 2, 2, 2, 3, 3, 4, 4, 5}
	c := NewCounter(l, -1)
	fmt.Println(c.MostCommon())
	// Output: [2]
}

func ExampleCounter_LeastCommon() {
	l := arraylist.ArrayList[int]{1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 5}
	c := NewCounter(l, -1)
	fmt.Println(c.LeastCommon())
	// Output: [5]
}

func ExampleCounter_Increment() {
	l := arraylist.ArrayList[int]{1, 2, 2, 2, 3, 3, 4, 4, 5}
	c := NewCounter(l, -1)
	fmt.Println(c.Get(2))
	c.Increment(2, 2)
	fmt.Println(c.Get(2))
	// Output:
	// 3
	// 5
}

func ExampleCounter_Subtract() {
	l := arraylist.ArrayList[int]{1, 2, 2, 2, 3, 3, 4, 4, 5}
	c := NewCounter(l, -1)
	fmt.Println(c.Get(2))
	c.Subtract(2, 2)
	fmt.Println(c.Get(2))
	// Output:
	// 3
	// 1
}

func ExampleCounter_Remove() {
	l := arraylist.ArrayList[int]{1, 2, 2, 2, 3, 3, 4, 4, 5}
	c := NewCounter(l, -1)
	fmt.Println(c.Get(2))
	c.Remove(2)
	fmt.Println(c.Get(2))
	// Output:
	// 3
	// -1
}

func ExampleCounter_Total() {
	l := arraylist.ArrayList[int]{1, 2, 2, 2, 3, 3, 4, 4, 5}
	c := NewCounter(l, -1)
	fmt.Println(c.Total())
	// Output: 9
}

func ExampleCounter_Reset() {
	l := arraylist.ArrayList[int]{1, 2, 2, 2, 3, 3, 4, 4, 5}
	c := NewCounter(l, -1)
	fmt.Println(c.Get(2))
	c.Reset()
	fmt.Println(c.Get(2))
	// Output:
	// 3
	// -1
}

func ExampleMergeCounts() {
	l1 := arraylist.ArrayList[int]{1, 2, 2, 2, 3, 3, 4, 4, 5}
	c1 := NewCounter(l1)
	l2 := arraylist.ArrayList[int]{1, 2, 7, 8, 9, 10}
	c2 := NewCounter(l2)
	mc := MergeCounts(c1, c2)
	fmt.Println(c1.Get(2))
	fmt.Println(c2.Get(2))
	fmt.Println(mc.Get(2))
	// Output:
	// 3
	// 1
	// 4
}
