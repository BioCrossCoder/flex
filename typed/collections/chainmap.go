// Package collections provides several convenient data structures.
package collections

import "github.com/biocrosscoder/flex/typed/collections/dict"

// ChainMap represents a data structure that allows chaining multiple maps together.
type ChainMap[K comparable, V any] struct {
	items  *dict.Dict[K, V]
	parent *ChainMap[K, V]
}

// NewChainMap creates a new ChainMap with the provided maps as parent nodes.
func NewChainMap[K comparable, V any](maps ...*dict.Dict[K, V]) *ChainMap[K, V] {
	cm := &ChainMap[K, V]{
		items: &dict.Dict[K, V]{},
	}
	node := cm
	for _, m := range maps {
		node.parent = &ChainMap[K, V]{
			items: m,
		}
		node = node.parent
	}
	return cm
}

// Set sets the key-value pair in the current ChainMap and returns a pointer to the ChainMap.
func (cm *ChainMap[K, V]) Set(key K, value V) *ChainMap[K, V] {
	_ = cm.items.Set(key, value)
	return cm
}

// Parent returns the parent ChainMap node.
func (cm *ChainMap[K, V]) Parent() *ChainMap[K, V] {
	return cm.parent
}

// Get retrieves the value for the key from the current ChainMap or its parent nodes.
func (cm *ChainMap[K, V]) Get(key K) (value V, ok bool) {
	node := cm
	for node != nil {
		ok = node.items.Has(key)
		if ok {
			value = node.items.Get(key)
			break
		}
		node = node.parent
	}
	return
}

// NewChild creates a new child ChainMap with the current ChainMap as its parent.
func (cm *ChainMap[K, V]) NewChild() *ChainMap[K, V] {
	return &ChainMap[K, V]{
		items:  &dict.Dict[K, V]{},
		parent: cm,
	}
}

// Parents returns a slice of all the parent ChainMap nodes in the chain.
func (cm *ChainMap[K, V]) Parents() []*ChainMap[K, V] {
	parents := make([]*ChainMap[K, V], 0)
	parent := cm.parent
	for parent != nil {
		parents = append(parents, parent)
		parent = parent.parent
	}
	return parents
}

// Maps returns a slice of all the maps in the chain, including the current ChainMap and its parents.
func (cm *ChainMap[K, V]) Maps() []*dict.Dict[K, V] {
	maps := make([]*dict.Dict[K, V], 0)
	node := cm
	for node != nil {
		maps = append(maps, node.items)
		node = node.parent
	}
	return maps
}

// Items returns the dictionary of the current ChainMap node.
func (cm *ChainMap[K, V]) Items() *dict.Dict[K, V] {
	return cm.items
}
