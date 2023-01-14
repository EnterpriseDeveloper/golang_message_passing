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
		logs.WriteToLogFile(key + " item added!")
	} else {
		logs.WriteToLogFile("Item already exist, try another one")
	}
}

func getItem(key string) {
	value, exist := m.Get(key)
	if !exist {
		logs.WriteToLogFile("Item not found.")
	} else {
		logs.WriteToLogFile(key + ", " + strconv.FormatInt(value, 10))
	}
}

func getAllItems() {
	for el := m.Front(); el != nil; el = el.Next() {
		logs.WriteToLogFile(el.Key + ", " + strconv.FormatInt(el.Value, 10))
	}
}

func removeItem(key string) {
	_, exist := m.Get(key)
	if !exist {
		logs.WriteToLogFile("Item not found.")
	} else {
		m.Delete(key)
		logs.WriteToLogFile(key + " item deleted!")
	}
}
