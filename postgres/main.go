package main

import (
	"fmt"

	"time"

	"encoding/json"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// define model
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Age       int
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Handler struct {
	db *gorm.DB
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	h.db.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	h.db.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	h.db.Unscoped().Delete(&User{}, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	dsn := "host=localhost user=postgres password=password dbname=goapi port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect:", err)
		return
	}
	fmt.Println("connected to database!")

	// auto create table
	db.AutoMigrate(&User{})

	h := &Handler{db: db}

	mux := http.NewServeMux()
	mux.HandleFunc("/users", h.getUsers)
	mux.HandleFunc("/users/create", h.createUser)
	mux.HandleFunc("/users/delete", h.deleteUser)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", mux)
}
