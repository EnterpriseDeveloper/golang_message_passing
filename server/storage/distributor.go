package storage

import (
	"fmt"
	"strconv"
	"time"

	"github.com/elliotchance/orderedmap/v2"

	"server/logs"
)

var m = orderedmap.NewOrderedMap[string, int64]()

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
		msg := key + " item added!"
		fmt.Println(msg)
		logs.WriteToLogFile(msg)
	} else {
		msg := "Item already exist, try another one"
		fmt.Println(msg)
		logs.WriteToLogFile(msg)
	}
}

func getItem(key string) {
	value, exist := m.Get(key)
	if !exist {
		msg := "Item not found."
		fmt.Println(msg)
		logs.WriteToLogFile(msg)
	} else {
		msg := key + ", " + strconv.FormatInt(value, 10)
		fmt.Println(msg)
		logs.WriteToLogFile(msg)
	}
}

func getAllItems() {
	for el := m.Front(); el != nil; el = el.Next() {
		msg := el.Key + ", " + strconv.FormatInt(el.Value, 10)
		fmt.Println(msg)
		logs.WriteToLogFile(msg)
	}
}

func removeItem(key string) {
	_, exist := m.Get(key)
	if !exist {
		msg := "Item not found."
		fmt.Println(msg)
		logs.WriteToLogFile(msg)
	} else {
		msg := key + " item deleted!"
		m.Delete(key)
		fmt.Println(msg)
		logs.WriteToLogFile(msg)
	}
}
