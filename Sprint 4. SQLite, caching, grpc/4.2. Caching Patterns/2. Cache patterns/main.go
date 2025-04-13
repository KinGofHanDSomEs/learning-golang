package main

import (
	"container/list"
	"fmt"
	"sync"
)

type MRUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
	mu       sync.RWMutex
}

type kv struct {
	Key, Value string
}

func NewMRUCache(capacity int) *MRUCache {
	if capacity < 1 {
		return nil
	}
	return &MRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

func (c *MRUCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.cache[key]; ok {
		elem.Value.(*kv).Value = value
		c.list.MoveToFront(elem)
	} else {
		if c.capacity <= len(c.cache) {
			backElem := c.list.Back()
			if backElem != nil {
				c.list.Remove(backElem)
				delete(c.cache, backElem.Value.(*kv).Key)
			}
		}
		c.cache[key] = c.list.PushFront(&kv{Key: key, Value: value})
	}
}

func (c *MRUCache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		return elem.Value.(*kv).Value, true
	}
	return "", false
}

func main() {
	cache := NewMRUCache(2)
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	cache.Set("key1", "new value1")
	value, ok := cache.Get("key1")
	if !ok || value != "new value1" {
		fmt.Printf("Expected value for key1 to be 'new value1', but got '%s'\n", value)
	}

	// Тест получения элемента из кэша
	value, ok = cache.Get("key2")
	if !ok || value != "value2" {
		fmt.Printf("Expected value for key2 to be 'value2', but got '%s'\n", value)
	}

	// Тест получения элемента из кэша
	value, ok = cache.Get("key2")
	if !ok || value != "value2" {
		fmt.Printf("Expected value for key2 to be 'value2', but got '%s'\n", value)
	}

	// Тест получения элемента из кэша
	value, ok = cache.Get("key2")
	if !ok || value != "value2" {
		fmt.Printf("Expected value for key2 to be 'value2', but got '%s'\n", value)
	}

	// Тест заполнения кэша до максимальной ёмкости
}
