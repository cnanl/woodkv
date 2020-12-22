package skiplist

import (
	"math/rand"
	"woodkv/e"
	"woodkv/store"
)

func New() *SkipList {
	return &SkipList{
		head:     &Element{next: make([]*Element, defaultMaxLevel)},
		p:        defaultP,
		level:    0,
		maxLevel: defaultMaxLevel,
	}
}

func (sk *SkipList) Get(key store.Key) (store.Value, error) {
	x := sk.head
	for i := sk.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key.Less(key) {
			x = x.next[i]
		}
	}
	//now x is the biggest element which is less than key
	x = x.next[0]
	if x != nil && x.key == key {
		return x.value, nil
	}
	return nil, e.NotFound
}
func (sk *SkipList) Put(key store.Key, val store.Value) error {
	update := make([]*Element, sk.maxLevel)
	x := sk.head
	for i := sk.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key.Less(key) {
			x = x.next[i]
		}
		update[i] = x //last element of each level less than key
	}
	x = x.next[0]
	//Key already exists, just replace the value
	if x != nil && x.key == key {
		x.value = val
		return nil
	}
	level := sk.randomLevel()
	if level > sk.level {
		level = sk.level + 1
		update[sk.level] = sk.head
		sk.level = level
	}

	//insert the element into skiplist
	ele := &Element{
		key:   key,
		value: val,
		next:  make([]*Element, level),
	}
	for i := 0; i < level; i++ {
		ele.next[i] = update[i].next[i]
		update[i].next[i] = ele
	}
	sk.length++
	return nil
}
func (sk *SkipList) Delete(key store.Key) error {
	update := make([]*Element, sk.maxLevel)
	x := sk.head
	for i := sk.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].key.Less(key) {
			x = x.next[i]
		}
		update[i] = x //last element of each level less than key
	}
	x = x.next[0]
	if x != nil && x.key == key {
		for i := 0; i < sk.level; i++ {
			update[i].next[i] = x.next[i]
		}
		sk.length--
		return nil
	}
	return e.NotFound
}
func (sk *SkipList) PrefixScan(n int) []interface{} {
	x := sk.First()
	var res []interface{}
	for n > 0 && x != nil {
		res = append(res, x.key)
		n--
		x = x.next[0]
	}
	return res
}
func (sk *SkipList) randomLevel() int {
	level := 1
	for rand.Float32() < sk.p && sk.level < sk.maxLevel {
		level++
	}
	return level
}

func (sk *SkipList) First() *Element {
	return sk.head.next[0]
}
