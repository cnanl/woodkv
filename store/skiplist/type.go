package skiplist

import "woodkv/store"

const (
	defaultMaxLevel int     = 16
	defaultP        float32 = 0.25
)

type Element struct {
	key   store.Key
	value store.Value
	next  []*Element
}

func (e *Element) Next() *Element {
	if e == nil {
		return nil
	}
	return e.next[0]
}

type SkipList struct {
	head     *Element
	length   int
	level    int
	maxLevel int
	p        float32
}
