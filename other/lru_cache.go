package main

import (
	"container/list"
	"fmt"
)

// Least Recently Used — алгоритм кэширования, при котором вытесняются значения,
// которые дольше всего не запрашивались. Алгоритмическая сложность O(1),
// а потому кеш работает очень быстро и используется в memcached.
// Кеш имеет очередь фиксированного размера. Когда новый элемент попадает в кеш, то добавляется в начало очереди.
// При запросе элемента очередь выталкивает элемент в начало, а если нужно освободить место,
// то из кеша вытесняется последний элемент.

type Item struct {
	Key   string
	Value string
}

type LRU struct {
	capacity int
	items    map[string]*list.Element
	queue    *list.List
}

func NewLru(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

func (c *LRU) Set(key string, value string) bool {
	if element, ok := c.items[key]; ok { // если значение в мапе уже существует
		c.queue.MoveToFront(element)        // смещаем в списке на первое место
		element.Value.(*Item).Value = value // перезаписываем значение для ключа
		return true
	}
	if c.queue.Len() == c.capacity { // если наш кеш уже достиг своей емкости, то удаляем последний элемент в очереди
		if element := c.queue.Back(); element != nil { // достаем последний элемент
			item := c.queue.Remove(element).(*Item) // удаляем элемент из списка
			delete(c.items, item.Key)               // удаляем элемент из мапы
		}
	}
	c.items[key] = c.queue.PushFront(&Item{key, value}) // записываем в мапу новый элемент
	return true
}

func (c *LRU) Get(key string) string {
	element, ok := c.items[key]
	if !ok {
		return ""
	}
	c.queue.MoveToFront(element) // продвигаем искомый элемент в начало очереди
	return element.Value.(*Item).Value
}

func (c *LRU) PrintAll() {
	for range len(c.items) {
		fmt.Printf(" %s:%s ", c.queue.Front().Value.(*Item).Key, c.queue.Front().Value.(*Item).Value)
		c.queue.MoveToBack(c.queue.Front())
	}
	fmt.Println()
}

func main() {
	lru := NewLru(3)
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")
	lru.Set("key3", "value3")
	lru.Get("key1")
	lru.PrintAll()
	lru.Set("key4", "value4")
	lru.Set("key5", "value5")
	lru.Get("key2")
	lru.Get("key1")
	lru.Set("key6", "value6")
	lru.Get("key1")
	lru.PrintAll()
}
