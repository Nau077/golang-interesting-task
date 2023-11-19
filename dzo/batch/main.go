package main

import (
	"fmt"
	"sync"
)

func main() {
	entitis := GetByHeaders()
	fmt.Println(len(entitis))
}

type header struct {
	id int
}
type headers []header

type mSync struct {
	m  map[int]header
	mx sync.Mutex
}

func NewMSync() *mSync {
	return &mSync{
		m: make(map[int]header),
	}
}

func (c *mSync) Store(key int, value header) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func (c *mSync) Load(key int) (header, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]

	return val, ok
}

func GetByHeaders() map[int]header {
	h := headers{}
	max := 1000
	m := NewMSync()
	threshold := 200
	for i := 0; i < max; i++ {
		h = append(h, header{id: i})
	}

	length := len(h)
	ch := make(chan header)
	w := sync.WaitGroup{}
	w.Add(length / threshold) // это 5 будет в данном случае - 1000/200

	for start := 0; start < length; {
		func(start int) {
			end := start + threshold
			if end > length {
				end = length
			}
			go GetBackupsEntitiesByIDs(h[start:end], ch, &w)

			start = end
		}(start)
	}

	go func() {
		for h := range ch {
			m.Store(h.id, h)
		}
	}()

	w.Wait()

	return m.m
}

func GetBackupsEntitiesByIDs(h headers, ch chan header, w *sync.WaitGroup) {
	defer w.Done()

	for i := range h {
		ch <- header{
			id: h[i].id * 2,
		}
	}
}
