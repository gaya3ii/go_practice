package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// define model
type User struct {
	gorm.Model
	Name  string
	Email string
	Age   int
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
	fmt.Println("table created!")

	// create a user
	db.Create(&User{Name: "John", Email: "john@example.com", Age: 25})

	var user User
	db.First(&user, 5)
	db.Delete(&user)
	fmt.Println("before:", user.Name)

	// update name
	db.Model(&user).Update("name", "Jane")
	fmt.Println("after:", user.Name)
	fmt.Println("user created!")

	// read all users
	var users []User
	db.Find(&users)
	fmt.Println("users:", users)
}
