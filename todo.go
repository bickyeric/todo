package main

// Todo hold data for todo
type Todo struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Notes string `json:"notes"`
}
