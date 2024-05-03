package collections

import (
	"flex/collections/arraylist"
	"flex/collections/dict"
	"flex/collections/set"
	"math"
)

type Counter struct {
	records      dict.Dict
	groups       dict.Dict
	defaultCount int
}

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

func (c *Counter) Get(item any) int {
	return c.records.Get(item, c.defaultCount).(int)
}

func (c *Counter) Set(item any, count int) *Counter {
	_ = c.records.Set(item, count)
	members := c.groups.Get(count, set.Set{}).(set.Set)
	_ = c.groups.Set(count, *members.Add(item))
	return c
}

func (c *Counter) Increment(item any, counts ...int) *Counter {
	count := 1
	if len(counts) > 0 {
		count = counts[0]
	}
	return c.Set(item, c.Get(item)+count)
}

func (c *Counter) Subtract(item any, counts ...int) *Counter {
	count := 1
	if len(counts) > 0 {
		count = counts[0]
	}
	return c.Set(item, c.Get(item)-count)
}

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

func (c *Counter) SetDefault(count int) *Counter {
	c.defaultCount = count
	return c
}

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

func (c *Counter) Total() (total int) {
	for _, v := range c.records {
		total += v.(int)
	}
	return
}

func (c *Counter) Elements() arraylist.ArrayList {
	elements := make(arraylist.ArrayList, len(c.records))
	i := 0
	for k := range c.records {
		elements[i] = k
		i++
	}
	return elements
}

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

func (c *Counter) Clear() *Counter {
	_ = c.records.Clear()
	_ = c.groups.Clear()
	return c
}

func (c Counter) Equal(another Counter) bool {
	keys1 := set.Of(c.Elements()...)
	keys2 := set.Of(another.Elements()...)
	if !keys1.SymmetricDifference(keys2).Empty() {
		return false
	}
	for k, v := range c.records {
		if v != another.Get(k) {
			return false
		}
	}
	return true
}

func (c Counter) Copy() Counter {
	return Counter{
		c.records.Copy(),
		c.groups.Copy(),
		c.defaultCount,
	}
}

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
