package internal

import (
	"fmt"
	"sync"
	"time"
)


type Cache struct {
	cacheEntries map[string]*cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache{
	ticker := time.NewTicker(interval)
	cacheEntries := make(map[string]*cacheEntry)
	creationTime := time.Now()
	newCache := &Cache{
		cacheEntries: cacheEntries, 
	}

	go func(){
			for {
			<- ticker.C
			newCache.reapLoop(creationTime)
		}
	}()

	return newCache
}

func (c *Cache) Add(key string,val []byte) {
	fmt.Println("adding a key",key)
	// fmt.Println("adding a value",val)
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheEntries[key]
	//if doesnt exist create a new one and assign it back to the map
	if !ok {
		entry = &cacheEntry{}
		c.cacheEntries[key] = entry
	}

	entry.val = val
	entry.createdAt = time.Now()

	for key, _ := range c.cacheEntries {
		fmt.Println(key)
	}

	fmt.Println("current entries:",c.cacheEntries)
}

func (c *Cache) Get(key string) ([]byte,bool){
	fmt.Println("get called")
	entry,ok := c.cacheEntries[key]
	if !ok {
		return nil,false
	}
	return entry.val,true
}

func (c *Cache) reapLoop(t time.Time) {
	for key, entry := range c.cacheEntries	{
		// fmt.Println("entry",entry.createdAt) 
		// fmt.Println("cache time",t)

		if  entry.createdAt.After(t) {
			delete(c.cacheEntries,key)
		} 
	}
}
