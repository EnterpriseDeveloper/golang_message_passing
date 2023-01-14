package storage

import (
	"fmt"
	"time"

	"github.com/elliotchance/orderedmap/v2"
)

var m = orderedmap.NewOrderedMap[string, any]()

func Distribute(msg string, key string) {
	switch msg {
	case "addItem":
		addItem(key)
	case "getItem":
		getItem(key)
	case "getAllItems":
		getAllItems()
	case "removeItem":
		removeItem(key)
	default:
		fmt.Println("Message do not exist")
	}
}

func addItem(key string) {
	_, exist := m.Get(key)
	if !exist {
		m.Set(key, time.Now().Unix())
		fmt.Println(key, " item added!")
	} else {
		fmt.Println("Item already exist, try another one")
	}
}

func getItem(key string) {
	value, exist := m.Get(key)
	if !exist {
		fmt.Println("Item not found.")
	} else {
		fmt.Println(key, value)
	}
}

func getAllItems() {
	for el := m.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}
}

func removeItem(key string) {
	_, exist := m.Get(key)
	if !exist {
		fmt.Println("Item not found.")
	} else {
		m.Delete(key)
		fmt.Println(key, " item deleted!")
	}
}
