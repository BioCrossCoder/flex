package collections

import "flex/collections/dict"

type ChainMap struct {
	items  *dict.Dict
	parent *ChainMap
}

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

func (cm *ChainMap) Set(key, value any) {
	_ = cm.items.Set(key, value)
}

func (cm *ChainMap) Parent() *ChainMap {
	return cm.parent
}

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

func (cm *ChainMap) NewChild() *ChainMap {
	return &ChainMap{
		items:  &dict.Dict{},
		parent: cm,
	}
}

func (cm *ChainMap) Parents() []*ChainMap {
	parents := make([]*ChainMap, 0)
	parent := cm.parent
	for parent != nil {
		parents = append(parents, parent)
		parent = parent.parent
	}
	return parents
}

func (cm *ChainMap) Maps() []*dict.Dict {
	maps := make([]*dict.Dict, 0)
	node := cm
	for node != nil {
		maps = append(maps, node.items)
		node = node.parent
	}
	return maps
}

func (cm *ChainMap) Items() *dict.Dict {
	return cm.items
}
