package model

import (
	"fmt"
	"strconv"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int
	memo  string
}

func NewItem(id int, name string, price int, memo string) *Item {
	return &Item{
		ID:    id,
		Name:  name,
		Price: price,
		memo:  memo,
	}
}

func GetItem(id int) *Item {
	// 疑似データ 作成
	return NewItem(id, "name-"+strconv.Itoa(id), id*100, "memo-"+strconv.Itoa(id))
}

func GetItems(startId, limit int) []*Item {
	var items []*Item
	for i := startId; i <= (startId + limit - 1); i++ {
		item := GetItem(i)
		items = append(items, item)
	}
	return items
}

func (i Item) String() string {
	return fmt.Sprintf("Item(id=%d, name=%s, price=%d, memo=%s)", i.ID, i.Name, i.Price, i.memo)
}
