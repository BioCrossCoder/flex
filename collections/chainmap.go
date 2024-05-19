// Package collections provides several convenient data structures.
package collections

import "github.com/biocrosscoder/flex/collections/dict"

// ChainMap struct represents a chain of Dict where a new Dict is linked to its parent.
type ChainMap struct {
	items  *dict.Dict
	parent *ChainMap
}

// NewChainMap creates and returns a new ChainMap with the provided dictionaries linked as parents.
func NewChainMap(maps ...*dict.Dict) *ChainMap {
	cm := &ChainMap{
		items: &dict.Dict{},
	}
	node := cm
	for _, m := range maps {
		node.parent = &ChainMap{
			items: m,
		}
		node = node.parent
	}
	return cm
}

// Set sets the key-value pair in the current level's dictionary and returns the ChainMap.
func (cm *ChainMap) Set(key, value any) *ChainMap {
	_ = cm.items.Set(key, value)
	return cm
}

// Parent returns the parent ChainMap.
func (cm *ChainMap) Parent() *ChainMap {
	return cm.parent
}

// Get retrieves the value of the key from the current level or its parent dictionaries.
func (cm *ChainMap) Get(key any) (value any, ok bool) {
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

// NewChild creates and returns a new ChainMap with the current level's dictionary linked as the parent.
func (cm *ChainMap) NewChild() *ChainMap {
	return &ChainMap{
		items:  &dict.Dict{},
		parent: cm,
	}
}

// Parents returns a slice of all the parent ChainMaps in the chain.
func (cm *ChainMap) Parents() []*ChainMap {
	parents := make([]*ChainMap, 0)
	parent := cm.parent
	for parent != nil {
		parents = append(parents, parent)
		parent = parent.parent
	}
	return parents
}

// Maps returns a slice of all the dictionaries in the chain, including the current and parent dictionaries.
func (cm *ChainMap) Maps() []*dict.Dict {
	maps := make([]*dict.Dict, 0)
	node := cm
	for node != nil {
		maps = append(maps, node.items)
		node = node.parent
	}
	return maps
}

// Items returns the dictionary of key-value pairs for the current level.
func (cm *ChainMap) Items() *dict.Dict {
	return cm.items
}
