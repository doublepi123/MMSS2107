package LRUCache

import "container/list"

type item struct {
	key   string
	value string
	pos   *list.Element
}

type LRU struct {
	ul      list.List
	data    map[string]*item
	Maxsize int
}

func (lru *LRU) Get(key string) string {
	i := lru.data[key]
	if i != nil {
		lru.ul.Remove(i.pos)
		lru.ul.PushBack(i)
		return i.value
	} else {
		return ""
	}
}

func (lru *LRU) Set(key string, value string) {
	i := lru.data[key]
	if i == nil {
		i = &item{key: key, value: value}
		lru.data[key] = i
	} else {
		lru.ul.Remove(i.pos)
	}
	lru.ul.PushBack(i)
	if lru.ul.Len() > lru.Maxsize {
		delete(lru.data, lru.ul.Front().Value.(*item).key)
		lru.ul.Remove(lru.ul.Front())
	}
	i.pos = lru.ul.Back()
}

func (lru *LRU) Init() {
	lru.data = make(map[string]*item)
	lru.ul = list.List{}
}
func (lru *LRU) Remove(key string) {
	t := lru.data[key]
	delete(lru.data, key)
	lru.ul.Remove(t.pos)
}
