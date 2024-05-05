package collections

import "flex/typed/collections/dict"

type ChainMap[K comparable, V any] struct {
	items  *dict.Dict[K, V]
	parent *ChainMap[K, V]
}

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

func (cm *ChainMap[K, V]) Set(key K, value V) *ChainMap[K, V] {
	_ = cm.items.Set(key, value)
	return cm
}

func (cm *ChainMap[K, V]) Parent() *ChainMap[K, V] {
	return cm.parent
}

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

func (cm *ChainMap[K, V]) NewChild() *ChainMap[K, V] {
	return &ChainMap[K, V]{
		items:  &dict.Dict[K, V]{},
		parent: cm,
	}
}

func (cm *ChainMap[K, V]) Parents() []*ChainMap[K, V] {
	parents := make([]*ChainMap[K, V], 0)
	parent := cm.parent
	for parent != nil {
		parents = append(parents, parent)
		parent = parent.parent
	}
	return parents
}

func (cm *ChainMap[K, V]) Maps() []*dict.Dict[K, V] {
	maps := make([]*dict.Dict[K, V], 0)
	node := cm
	for node != nil {
		maps = append(maps, node.items)
		node = node.parent
	}
	return maps
}

func (cm *ChainMap[K, V]) Items() *dict.Dict[K, V] {
	return cm.items
}
