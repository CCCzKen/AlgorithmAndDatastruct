package lrucache

import (
	"sync"
	"container/list"
	"errors"
	"fmt"
)

type LruCache struct {
	capacity int
	mx       sync.RWMutex
	list     *list.List
	cache    map[interface{}]*list.Element
}

type entry struct {
	key   interface{}
	value interface{}
}

func InitLruCache(cap int) *LruCache {
	return &LruCache{
		capacity: cap,
		list:     list.New(),
		cache:    make(map[interface{}]*list.Element),
	}
}

func (lru *LruCache) Get(key interface{}) interface{} {
	lru.mx.RLock()
	defer lru.mx.RUnlock()
	element := lru.cache[key]
	lru.list.MoveToFront(element)
	return element.Value.(*entry).value
}

func (lru *LruCache) Set(key interface{}, value interface{}) error {
	lru.mx.Lock()
	defer lru.mx.Unlock()

	if lru.list == nil {
		return errors.New("LruCache 结构初始化错误")
	}

	// 插入已存在的元素
	if element, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(element)
		element.Value.(*entry).value = value
		return nil
	}

	// 插入新元素
	element := lru.list.PushFront(&entry{key, value})
	lru.cache[key] = element

	// 若超出容量
	if lru.capacity != 0 && lru.list.Len() > lru.capacity {
		lru.RemoveOldest()
	}
	return nil
}


func (lru *LruCache) Delete(key interface{}) {
	lru.mx.Lock()
	defer lru.mx.Unlock()
	if lru.cache == nil {
		return
	}

	if element, ok := lru.cache[key]; ok {
		lru.list.Remove(element)
		delete(lru.cache, key)
		return
	}
}


func (lru *LruCache) RemoveOldest() {
	if lru.cache == nil {
		return
	}
	element := lru.list.Back()
	if element != nil {
		lru.list.Remove(element)
		key := element.Value.(*entry).key
		delete(lru.cache, key)
	}
}


func (lru *LruCache) Size() int {
	return lru.list.Len()
}


func main() {
	lru := InitLruCache(2)
	lru.Set("realname", "chenzhaokang")
	lru.Set("age", 25)
	fmt.Println(lru.list.Len())
	lru.Set("sex", "male")
	fmt.Println(lru.list.Len())
	fmt.Println(lru.list.Front().Value.(*entry).value)
	lru.Delete("sex")
	fmt.Println(lru.list.Len())
	fmt.Println(lru.list.Front().Value.(*entry).value)
}
