package hw04lrucache

import "fmt"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	item, exists := l.items[key]
	if exists {
		l.queue.MoveToFront(item)
		item.Value = value
	} else {
		newListItem := l.queue.PushFront(value)
		l.items[key] = newListItem
		if l.queue.Len() > l.capacity {
			elemToDelete := l.queue.Back()
			fmt.Println(elemToDelete)
			l.queue.Remove(elemToDelete)
		}
	}
	return exists
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	item, exists := l.items[key]
	if exists {
		l.queue.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	for _, v := range l.items {
		l.queue.Remove(v)
	}
	l.items = make(map[Key]*ListItem)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
