package main

import (
	"sort"
	"sync"
)

type Cache struct {
	UpperBound int
	LowerBound int
	values     map[string]*Element
	keys       []string
	vals       []Element
	lock       sync.Mutex
}

type Element struct {
	key   string
	value interface{}
	freq  int
}

func New() *Cache {
	return &Cache{}
}

func (c *Cache) Swap(i, j int) {
	c.keys[i], c.keys[j] = c.keys[j], c.keys[i]
	c.vals[i], c.vals[j] = c.vals[j], c.vals[i]
}
func (c *Cache) Less(i, j int) bool {
	return c.vals[i].freq < c.vals[j].freq
}

func (c *Cache) Has(key string) bool {
	_, ok := c.values[key]
	return ok
}

func (c *Cache) increment(elem *Element) {
	elem.freq++
}

func (c *Cache) Get(key string) interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()
	if e, ok := c.values[key]; ok {
		c.increment(e)
		return e.value
	}
	return nil
}

func (c *Cache) Set(key string, value interface{}) {
	c.lock.Lock()
	if elem, ok := c.values[key]; ok {
		c.values[key] = &Element{
			key:   key,
			value: value,
			freq:  elem.freq + 1,
		}
	} else {
		if c.Len() >= c.UpperBound && c.UpperBound != 0 && c.LowerBound != 0 {
			sort.Sort(c)
			for c.Len() != c.LowerBound {
				firstKey := c.keys[0]
				delete(c.values, firstKey)
				c.keys = c.keys[1:]
				c.vals = c.vals[1:]
			}
		}
		elem := Element{
			key:   key,
			value: value,
			freq:  1,
		}
		c.values[key] = &elem
		c.keys = append(c.keys, key)
		c.vals = append(c.vals, elem)
	}
}

func (c *Cache) Len() int {
	return len(c.values)
}

func (c *Cache) GetFrequency(key string) int {
	if elem, ok := c.values[key]; ok {
		return elem.freq
	}
	return 0
}

func (c *Cache) Keys() []string {
	var result []string
	for key := range c.values {
		result = append(result, key)
	}
	return result
}

func (c *Cache) Evict(count int) int {
	sort.Sort(c)
	k := 0
	for k < count && c.Len() != 0 {
		firstKey := c.keys[0]
		delete(c.values, firstKey)
		c.keys = c.keys[1:]
		c.vals = c.vals[1:]
		k++
	}
	return k
}
