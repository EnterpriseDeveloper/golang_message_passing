package storage

import (
	"fmt"
	"time"

	"github.com/elliotchance/orderedmap/v2"
)

var m = orderedmap.NewOrderedMap[string, any]()

func Distribute(msg string, value string) {
	switch msg {
	case "addItem":
		m.Set(value, time.Now().Unix())
		fmt.Println(value, " item added!")
	case "getItem":
		printItem(value)
	case "getAllItems":
		printAllItems()
	case "removeItem":
		m.Delete("qux")
		fmt.Println(value, " item deleted!")
	default:
		fmt.Println("Message do not exist")
	}
}

func printAllItems() {
	for el := m.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}
}

func printItem(value string) {
	for el := m.Front(); el != nil; el = el.Next() {
		if el.Key == value {
			fmt.Println(el.Key, el.Value)
			break
		}
	}
}
