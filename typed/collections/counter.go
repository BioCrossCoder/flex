package collections

import (
	"github.com/biocrosscoder/flex/typed/collections/arraylist"
	"github.com/biocrosscoder/flex/typed/collections/dict"
	"github.com/biocrosscoder/flex/typed/collections/set"
	"math"
)

type Counter[T comparable] struct {
	records      dict.Dict[T, int]
	groups       dict.Dict[int, set.Set[T]]
	defaultCount int
}

func NewCounter[T comparable](items arraylist.ArrayList[T], defaultCounts ...int) *Counter[T] {
	records := make(dict.Dict[T, int])
	for _, item := range items {
		_ = records.Set(item, records.Get(item, 0)+1)
	}
	groups := make(dict.Dict[int, set.Set[T]])
	for item, count := range records {
		members := groups.Get(count, set.Set[T]{})
		_ = members.Add(item)
		_ = groups.Set(count, members)
	}
	defaultCount := 0
	if len(defaultCounts) > 0 {
		defaultCount = defaultCounts[0]
	}
	return &Counter[T]{
		records,
		groups,
		defaultCount,
	}
}

func (c *Counter[T]) Get(item T) int {
	return c.records.Get(item, c.defaultCount)
}

func (c *Counter[T]) Set(item T, count int) *Counter[T] {
	_ = c.records.Set(item, count)
	members := c.groups.Get(count, set.Set[T]{})
	_ = c.groups.Set(count, *members.Add(item))
	return c
}

func (c *Counter[T]) Increment(item T, counts ...int) *Counter[T] {
	count := 1
	if len(counts) > 0 {
		count = counts[0]
	}
	return c.Set(item, c.Get(item)+count)
}

func (c *Counter[T]) Subtract(item T, counts ...int) *Counter[T] {
	count := 1
	if len(counts) > 0 {
		count = counts[0]
	}
	return c.Set(item, c.Get(item)-count)
}

func (c *Counter[T]) Remove(item T) (exist bool) {
	exist = c.records.Has(item)
	if exist {
		count := c.Get(item)
		group := c.groups.Get(count)
		if group.Discard(item) && group.Empty() {
			_ = c.groups.Delete(count)
		}
		_ = c.records.Delete(item)
	}
	return
}

func (c *Counter[T]) SetDefault(count int) *Counter[T] {
	c.defaultCount = count
	return c
}

func (c *Counter[T]) MostCommon() arraylist.ArrayList[T] {
	maxCount := 0
	for count := range c.groups {
		if count > maxCount {
			maxCount = count
		}
	}
	items := make(arraylist.ArrayList[T], 0)
	group := c.groups.Get(maxCount, set.Set[T]{})
	for item := range group {
		_ = items.Push(item)
	}
	return items
}

func (c *Counter[T]) LeastCommon() arraylist.ArrayList[T] {
	minCount := math.MaxInt
	for count := range c.groups {
		if count < minCount {
			minCount = count
		}
	}
	items := make(arraylist.ArrayList[T], 0)
	group := c.groups.Get(minCount, set.Set[T]{})
	for item := range group {
		_ = items.Push(item)
	}
	return items
}

func (c *Counter[T]) Total() (total int) {
	for _, count := range c.records {
		total += count
	}
	return
}

func (c *Counter[T]) Elements() arraylist.ArrayList[T] {
	elements := make(arraylist.ArrayList[T], len(c.records))
	i := 0
	for k := range c.records {
		elements[i] = k
		i++
	}
	return elements
}

func (c *Counter[T]) Reset() *Counter[T] {
	items := make(set.Set[T])
	for k := range c.records {
		_ = c.records.Set(k, c.defaultCount)
		_ = items.Add(k)
	}
	c.groups = dict.Dict[int, set.Set[T]]{
		c.defaultCount: items,
	}
	return c
}

func (c *Counter[T]) Clear() *Counter[T] {
	_ = c.records.Clear()
	_ = c.groups.Clear()
	return c
}

func (c Counter[T]) Equal(another Counter[T]) bool {
	return c.records.Equal(another.records)
}

func (c Counter[T]) Copy() Counter[T] {
	return Counter[T]{
		c.records.Copy(),
		c.groups.Copy(),
		c.defaultCount,
	}
}

func MergeCounts[T comparable](counters ...*Counter[T]) *Counter[T] {
	counter := &Counter[T]{
		records:      make(dict.Dict[T, int]),
		groups:       make(dict.Dict[int, set.Set[T]]),
		defaultCount: 0,
	}
	for _, c := range counters {
		elements := c.Elements()
		elements.ForEach(func(a T) T {
			_ = counter.Increment(a, c.Get(a))
			return a
		})
	}
	return counter
}
