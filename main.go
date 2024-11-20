package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var items = []Item{}

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	json.NewDecoder(r.Body).Decode(&item)
	item.ID = uuid.New().String()
	items = append(items, item)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// Дополнительные функции
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	// Реализация обновления элемента
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// Реализация удаления элемента
}

func main() {
	http.HandleFunc("/items", GetItems)
	http.HandleFunc("/items/create", CreateItem)
	// Добавьте маршруты для обновления и удаления
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
