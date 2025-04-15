package main

import (
	"errors"
	"sync"
	"time"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func getProduct(productId int, db map[int]Product, cache *Cache) (Product, error) {
	if product, ok := cache.Get(productId); ok {
		return product, nil
	}
	product, ok := db[productId]
	if !ok {
		return Product{}, errors.New("product not found in database")
	}
	cache.Set(productId, product)
	return product, nil
}

func updateProduct(productId int, newProduct Product, db map[int]Product) error {
	db[productId] = newProduct
	return nil
}

type Cache struct {
	products  map[int]Product
	ttl       time.Duration
	startTime time.Time
	mu        sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		products:  make(map[int]Product),
		ttl:       time.Minute,
		startTime: time.Now(),
	}
}

func (c *Cache) Get(productId int) (Product, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	product, ok := c.products[productId]
	if !ok || time.Since(c.startTime) > c.ttl {
		return Product{}, false
	}
	return product, true
}

func (c *Cache) Set(productId int, product Product) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.products[productId] = product
}

func (c *Cache) Invalidate(productId int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.products, productId)
}
