package api

// Item is a single record returned by the API.
type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// items is the static in-memory dataset served by the API.
// Replace or extend it with the data this app should expose.
var items = []Item{
	{ID: 1, Name: "Item Satu", Description: "Contoh data pertama"},
	{ID: 2, Name: "Item Dua", Description: "Contoh data kedua"},
	{ID: 3, Name: "Item Tiga", Description: "Contoh data ketiga"},
}

func findItem(id int) (Item, bool) {
	for _, it := range items {
		if it.ID == id {
			return it, true
		}
	}
	return Item{}, false
}
