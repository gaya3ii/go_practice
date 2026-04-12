package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var users = map[int]User{
	1: {ID: 1, Name: "Alice", Age: 30},
	2: {ID: 2, Name: "Bob", Age: 25},
}
var nextID = 1

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := Person{Name: "James", Age: 20}
	json.NewEncoder(w).Encode(p)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(s)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// convert map to slice
	result := []User{}
	for _, user := range users {
		result = append(result, user)
	}
	json.NewEncoder(w).Encode(result)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// method check
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		return
	}

	// step 2 — decode request body
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//assign ID
	user.ID = nextID
	nextID++

	//save to map
	users[user.ID] = user

	// step 5 — return 201 + created user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// DELETE handler — only accepts DELETE
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// 1. check method
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 2. get id from query param
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3. convert id to int
	userID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 4. check if user exists
	if _, ok := users[userID]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// 5. delete from map
	delete(users, userID)

	// 6. return 204
	w.WriteHeader(http.StatusNoContent)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3. convert id to int
	userID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 4. check if user exists
	if _, ok := users[userID]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.ID = userID
	users[userID] = user

	// step 5 — return 200 + created user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/users", getUsersHandler)
	http.HandleFunc("/users/create", createUserHandler)
	http.HandleFunc("/users/delete", deleteUserHandler)
	http.HandleFunc("/users/update", updateUserHandler)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
