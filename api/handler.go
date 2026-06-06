package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mizanmahi/types" // ← importing our types module
)

// in-memory store for this demo
var store = map[string]*types.User{
	"1": {ID: "1", Name: "Alice", Email: "alice@example.com", CreatedAt: time.Now()},
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	users := make([]*types.User, 0, len(store))
	for _, u := range store {
		users = append(users, u)
	}
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	u, ok := store[id]
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(u)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u types.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u.CreatedAt = time.Now()
	store[u.ID] = &u
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
