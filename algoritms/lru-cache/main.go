package main

import (
	"fmt"
	"sync"
)

type LruCache struct {
	sync.Mutex
	items []LruItem
}

type LruItem struct {
	key   string
	value string
}

func (l *LruCache) Get(key string) LruItem {

	for index, el := range l.items {
		if key == el.key {
			l.MoveUp(index, el)
			return el
		}
	}

	return LruItem{}
}

func (l *LruCache) MoveUp(index int, lruItem LruItem) {
	l.Lock()
	// // удаляем элемент
	// l.items = append(l.items[:index], l.items[index+1:]...)
	// // добавляем в последний айтем
	// l.items = append(l.items, lruItem)

	if index == 0 {
		return
	}
	var newItems []LruItem
	newItems = append(newItems, l.items[index])

	for k, _ := range l.items {
		if k != index {
			newItems = append(newItems, l.items[k])
		}
	}
	l.items = newItems

	l.Unlock()

}

func (l *LruCache) Set(lruItem LruItem) {

	for index, el := range l.items {
		if el.key == lruItem.key {
			_ = append(l.items[:index], l.items[index+1:]...)
		}
	}
	l.Lock()
	l.items = append(l.items, lruItem)
	l.Unlock()

}

func main() {
	var users = LruCache{}
	users.Set(
		LruItem{
			key:   "seun",
			value: "james",
		},
	)
	users.Set(
		LruItem{
			key:   "ola",
			value: "kucha",
		},
	)

	users.Set(
		LruItem{
			key:   "dssdf",
			value: "kudsfsdcha",
		},
	)
	fmt.Println(users.items)

	users.Get("ola")

	fmt.Println(users.items)
}
