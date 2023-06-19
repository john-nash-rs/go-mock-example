package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

func ConnectDB() (*gorm.DB, error) {
	// Database connection configuration
	dsn := "root:root@tcp(localhost:3306)/nd?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertUser(db *gorm.DB, newUser *User) error {
	// Insert the user into the database
	result := db.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func main() {
	// Create a new user
	newUser := &User{
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	// Insert the user into the database
	err = InsertUser(db, newUser)
	if err != nil {
		panic("Failed to insert user")
	}

	fmt.Println("User inserted successfully")
}
