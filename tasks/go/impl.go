package ordcol

import (
	"sort"
)

type item struct {
    key int
	value int
}

type collection struct {
	storage map[int]Item
	currentOrder int
	orderStorage map[int]Item
	keys map[int]bool
}

type collectionIterator struct {
	index int
	keys []int
	collection *map[int]Item
}

func NewItem(key int, value int) *item {
    return &item{
		key: key,
		value: value,
	}
}

func NewCollection() *collection {
	return &collection{
		storage: make(map[int]Item),
		orderStorage: make(map[int]Item),
		currentOrder: 0,
		keys: make(map[int]bool),
	}
}

func (item item) Key() int {
	return item.key
}

func (item item) Value() int {
	return item.value
}

func (iter *collectionIterator) HasNext() bool {
	if iter.index < len(iter.keys) {
		return true
	}

	return false
}

func (iter *collectionIterator) Next() (Item, error) {
	if iter.HasNext() {
		itm := (*iter.collection)[iter.keys[iter.index]]

		iter.index++

		return itm, nil
	}

	return nil, ErrEmptyIterator
}

func (collection *collection) At(key int) (Item, bool) {
	if collection.keys[key] {
		return collection.storage[key], true
	}
	return nil, false
}

func (collection *collection) Add(item Item) error {
	key := item.Key()
	if collection.keys[key] {
		return ErrDuplicateKey
	}

	collection.keys[key] = true
	collection.storage[key] = item
	collection.orderStorage[collection.currentOrder] = item
	collection.currentOrder++

	return nil
}

func (collection *collection) IterateBy (order IterationOrder) Iterator {
	mp := collection.storage
	if order == ByInsertion {
		keys := make([]int, len(mp))

		for i := 0; i < len(mp); i++ {
			keys[i] = i
		}

		return &collectionIterator{
			index: 0,
			keys: keys,
			collection: &collection.orderStorage,
		}

	} else if order == ByKey {
		keys := make([]int, len(mp))

		i := 0
		for k := range mp {
			keys[i] = k
			i++
		}

		sort.Ints(keys)
		return &collectionIterator{
			index: 0,
			keys: keys,
			collection: &collection.storage,
		}
	} else {
		panic("no such order")
	}
}
