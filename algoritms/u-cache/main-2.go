package main

import (
	"sync"
	"time"
)

type UserProfile struct {
	UUID string
	Name string
}

type Bucket struct {
	UserProfile UserProfile
	Expire      int64
}

type Cache struct {
	sync.Mutex
	BucketMap         map[string]Bucket
	ClearGCInterval   time.Duration
	DefaultItemExpire time.Duration
}

func NewCache() *Cache {
	m := make(map[string]Bucket)

	cache := &Cache{
		BucketMap:         m,
		ClearGCInterval:   1 * time.Minute,
		DefaultItemExpire: 2 * time.Minute,
	}

	return cache
}

func (c *Cache) StartGarbageCollector() {
	for {
		select {
		case <-time.After(c.ClearGCInterval):
			for k, v := range c.BucketMap {
				if v.Expire > time.Now().UnixNano() {
					delete(c.BucketMap, k)
				}
			}
		}
	}
}
