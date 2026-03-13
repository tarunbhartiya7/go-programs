package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var productList = []Product{
	{ID: 1, Name: "Laptop", Price: 1000},
	{ID: 2, Name: "Phone", Price: 500},
	{ID: 3, Name: "Tablet", Price: 300},
}

func findProductByID(id int) *Product {
	for i := range productList {
		if productList[i].ID == id {
			return &productList[i]
		}
	}
	return nil
}

func removeProductByID(id int) bool {
	for i, p := range productList {
		if p.ID == id {
			productList = append(productList[:i], productList[i+1:]...)
			return true
		}
	}
	return false
}

func getNextID() int {
	max := 0
	for _, p := range productList {
		if p.ID > max {
			max = p.ID
		}
	}
	return max + 1
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(productList)
		case http.MethodPost:
			var p Product
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}
			p.ID = getNextID()
			productList = append(productList, p)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(p)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		idStr := strings.TrimPrefix(r.URL.Path, "/api/")
		if idx := strings.Index(idStr, "/"); idx != -1 {
			idStr = idStr[:idx]
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			p := findProductByID(id)
			if p == nil {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(p)
		case http.MethodPut:
			p := findProductByID(id)
			if p == nil {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			}
			var updated Product
			if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}
			p.Name = updated.Name
			p.Price = updated.Price
			json.NewEncoder(w).Encode(p)
		case http.MethodDelete:
			if !removeProductByID(id) {
				http.Error(w, "Product not found", http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}
