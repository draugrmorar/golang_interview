package main

import (
	"container/list"
	"iter"
)

// ОБЫЧНЫЙ КЭШ

//type Cache[K comparable, V any] interface{
//	Put(key K, value V)
//	Get(key K) (value V, ok bool)
//	Size() int
//}
//
//var _ Cache[any, any] = (*cacheImpl[any, any])(nil)
//
//type cacheImpl[K comparable, V any] struct {
//	data map[K]V
//}
//
//func New[K comparable, V any](cap int) Cache[K, V] {
//	return &cacheImpl[K, V]{
//		data: make(map[K]V, cap),
//	}
//}
//func (c *cacheImpl[K, V]) Put(k K, v V) {
//	c.data[k] = v
//}
//
//func (c *cacheImpl[K, V]) Get(k K) (v V, ok bool) {
//	if v, ok := c.data[k]; ok {
//		return v, true
//	}else {
//		return v, false
//	}
//}
//
//func (c *cacheImpl[K, V]) Size() int {
//	return len(c.data)
//}

// LRU cache
// LRU (Least Recently Used) — удаляем самое давно использованное

type LRUCache[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, ok bool)
	Size() int
	All() iter.Seq2[K, V]
}

var _ LRUCache[any, any] = (*lruCacheImpl[any, any])(nil)

type lruCacheImpl[K comparable, V any] struct {
	linkedList *list.List
	data       map[K]*list.Element
	capacity   int
}

type Node[K comparable, V any] struct {
	key   K
	value V
}

func New[K comparable, V any](cap int) *lruCacheImpl[K, V] {
	return &lruCacheImpl[K, V]{
		linkedList: list.New(),
		data:       make(map[K]*list.Element, cap),
		capacity:   cap,
	}
}

func (c *lruCacheImpl[K, V]) deleteLatest() {
	del := c.linkedList.Back()
	c.linkedList.Remove(del)
	delete(c.data, c.getValueFromEl(del).key)
}

func (c *lruCacheImpl[K, V]) Put(k K, v V) {
	if link, ok := c.data[k]; ok {
		n := c.getValueFromEl(link)
		n.value = v
		c.linkedList.MoveToFront(link)
		return
	}
	if c.Size() == c.capacity {
		c.deleteLatest()
	}
	node := &Node[K, V]{
		key:   k,
		value: v,
	}
	c.data[k] = c.linkedList.PushFront(node)
	return
}

func (c *lruCacheImpl[K, V]) getValueFromEl(el *list.Element) *Node[K, V] {
	switch v := el.Value.(type) {
	case *Node[K, V]:
		return v
	default:
		panic("")
	}
}

func (c *lruCacheImpl[K, V]) Get(k K) (v V, ok bool) {
	if link, ok := c.data[k]; ok {
		c.linkedList.MoveToFront(link) // сдвигаем в начало списка
		return c.getValueFromEl(link).value, true
	}
	var zero V
	return zero, false

}

func (c *lruCacheImpl[K, V]) Size() int {
	return len(c.data)
}

func (c *lruCacheImpl[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		cur := c.linkedList.Front()
		for range c.Size() {
			n := c.getValueFromEl(cur)
			if !yield(n.key, n.value) {
				return
			}
			cur = cur.Next()
		}
	}
}
