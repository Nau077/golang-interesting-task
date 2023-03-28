package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type UserProfile struct {
	UUID string
	Name string
	// Orders []Order
}

// type Order struct {
// 	UUID      string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

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
		ClearGCInterval:   10 * time.Minute,
		DefaultItemExpire: 5 * time.Minute,
	}

	go cache.startGC()

	return cache
}

func (c *Cache) startGC() {
	for {
		select {
		case <-time.After(c.ClearGCInterval):
			c.clear()
		}
	}
}

func (c *Cache) clear() {
	c.Lock()
	defer c.Unlock()

	for k, v := range c.BucketMap {
		if v.Expire > time.Now().UnixNano() {
			delete(c.BucketMap, k)
		}
	}
}

func (c *Cache) Set(id string, b Bucket) {
	expiration := time.Now().Add(c.DefaultItemExpire).UnixNano()

	c.Lock()
	c.BucketMap[id] = Bucket{
		UserProfile: b.UserProfile,
		Expire:      expiration,
	}

	c.Unlock()
}

func (c *Cache) Get(id string) (Bucket, error) {
	c.Lock()
	b, ok := c.BucketMap[id]
	c.Unlock()

	if !ok {
		return Bucket{}, errors.New("no data")
	}

	return b, nil
}

func main() {
	c := NewCache()

	c.Set(
		"kgkgk",
		Bucket{
			UserProfile: UserProfile{
				UUID: "fkfkfk",
				Name: "fkfkfk",
			},
		},
	)

	fmt.Println(c.Get("kgkgk"))
}
