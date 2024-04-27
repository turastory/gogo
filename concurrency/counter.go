package concurrency

import "sync"

type Counter struct {
	mu    sync.Mutex
	count int
}

func NewCounter() *Counter {
	return &Counter{
		mu: sync.Mutex{},
	}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}
