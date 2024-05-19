package collections

import (
	"github.com/biocrosscoder/flex/collections/arraylist"
	"github.com/biocrosscoder/flex/collections/dict"
	"github.com/biocrosscoder/flex/collections/set"
	"math"
)

// Counter struct represents a counter that counts the occurrences of items and provides various operations based on these counts.
type Counter struct {
	records      dict.Dict
	groups       dict.Dict
	defaultCount int
}

// NewCounter creates and initializes a new Counter with the given items and default count value.
func NewCounter(items arraylist.ArrayList, defaultCounts ...int) *Counter {
	records := make(dict.Dict)
	for _, item := range items {
		_ = records.Set(item, records.Get(item, 0).(int)+1)
	}
	groups := make(dict.Dict)
	for k, v := range records {
		count := v.(int)
		members := groups.Get(count, set.Set{}).(set.Set)
		_ = members.Add(k)
		_ = groups.Set(count, members)
	}
	defaultCount := 0
	if len(defaultCounts) > 0 {
		defaultCount = defaultCounts[0]
	}
	return &Counter{
		records,
		groups,
		defaultCount,
	}
}

// Get gets the count of the given item in the Counter. If the item is not in the Counter, the default count value is returned.
func (c *Counter) Get(item any) int {
	return c.records.Get(item, c.defaultCount).(int)
}

// Set sets the count of the given item in the Counter and returns the updated Counter.
func (c *Counter) Set(item any, count int) *Counter {
	_ = c.records.Set(item, count)
	members := c.groups.Get(count, set.Set{}).(set.Set)
	_ = c.groups.Set(count, *members.Add(item))
	return c
}

// Increment increments the count of the given item in the Counter by the specified amount and returns the updated Counter.
func (c *Counter) Increment(item any, counts ...int) *Counter {
	count := 1
	if len(counts) > 0 {
		count = counts[0]
	}
	return c.Set(item, c.Get(item)+count)
}

// Subtract decrements the count of the given item in the Counter by the specified amount and returns the updated Counter.
func (c *Counter) Subtract(item any, counts ...int) *Counter {
	count := 1
	if len(counts) > 0 {
		count = counts[0]
	}
	return c.Set(item, c.Get(item)-count)
}

// Remove removes the given item from the Counter and returns true if the item was in the Counter, false otherwise.
func (c *Counter) Remove(item any) (exist bool) {
	exist = c.records.Has(item)
	if exist {
		count := c.Get(item)
		group := c.groups.Get(count).(set.Set)
		if group.Discard(item) && group.Empty() {
			_ = c.groups.Delete(count)
		}
		_ = c.records.Delete(item)
	}
	return
}

// SetDefault sets the default count value of the Counter and returns the updated Counter.
func (c *Counter) SetDefault(count int) *Counter {
	c.defaultCount = count
	return c
}

// MostCommon returns a list of items with the highest count in the Counter.
func (c *Counter) MostCommon() arraylist.ArrayList {
	maxCount := 0
	for k := range c.groups {
		count := k.(int)
		if count > maxCount {
			maxCount = count
		}
	}
	items := make(arraylist.ArrayList, 0)
	group := c.groups.Get(maxCount, set.Set{}).(set.Set)
	for item := range group {
		_ = items.Push(item)
	}
	return items
}

// LeastCommon returns a list of items with the lowest count in the Counter.
func (c *Counter) LeastCommon() arraylist.ArrayList {
	minCount := math.MaxInt
	for k := range c.groups {
		count := k.(int)
		if count < minCount {
			minCount = count
		}
	}
	items := make(arraylist.ArrayList, 0)
	group := c.groups.Get(minCount, set.Set{}).(set.Set)
	for item := range group {
		_ = items.Push(item)
	}
	return items
}

// Total returns the total count of all items in the Counter.
func (c *Counter) Total() (total int) {
	for _, v := range c.records {
		total += v.(int)
	}
	return
}

// Elements returns a list of all distinct items in the Counter.
func (c *Counter) Elements() arraylist.ArrayList {
	elements := make(arraylist.ArrayList, len(c.records))
	i := 0
	for k := range c.records {
		elements[i] = k
		i++
	}
	return elements
}

// Reset resets the counts of all items in the Counter to the default count value and returns the updated Counter.
func (c *Counter) Reset() *Counter {
	items := make(set.Set)
	for k := range c.records {
		_ = c.records.Set(k, c.defaultCount)
		_ = items.Add(k)
	}
	c.groups = dict.Dict{
		c.defaultCount: items,
	}
	return c
}

// Clear removes all items from the Counter and returns the updated Counter.
func (c *Counter) Clear() *Counter {
	_ = c.records.Clear()
	_ = c.groups.Clear()
	return c
}

// Equal checks if the Counter is equal to another Counter based on their item counts.
func (c Counter) Equal(another Counter) bool {
	return c.records.Equal(another.records)
}

// Copy creates a deep copy of the Counter.
func (c Counter) Copy() Counter {
	return Counter{
		c.records.Copy(),
		c.groups.Copy(),
		c.defaultCount,
	}
}

// MergeCounts merges multiple Counters into a single Counter based on their item counts.
func MergeCounts(counters ...*Counter) *Counter {
	counter := &Counter{
		records:      make(dict.Dict),
		groups:       make(dict.Dict),
		defaultCount: 0,
	}
	for _, c := range counters {
		elements := c.Elements()
		elements.ForEach(func(a any) any {
			_ = counter.Increment(a, c.Get(a))
			return a
		})
	}
	return counter
}
