package main

import "sync"

type Counter struct {
	Count int
	mu    sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Count++
}

func (c *Counter) Value() int {
	return c.Count
}
