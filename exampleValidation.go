package main

import (
	"fmt"
	"regexp"

	"gorm.io/gorm"
)

func IsValidEmail(email string) bool {
	// Simple email validation using regular expression
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}

func SaveUser(db *gorm.DB, newUser *User) error {
	fmt.Println(" --- Hello ---")
	// Validate name
	if newUser.Name == "" {
		return fmt.Errorf("Name cannot be empty")
	}

	// Validate email
	if newUser.Email == "" || !IsValidEmail(newUser.Email) {
		return fmt.Errorf("Invalid email")
	}

	// Save the user in the database
	result := db.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ExampleWithValidation() {
	// Connect to the database
	db, err := ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	// Create a new user
	newUser := &User{
		Name:  "JohnV Doe",
		Email: "johnV@example.com",
	}

	// Save the user in the database
	err = SaveUser(db, newUser)
	if err != nil {
		panic("Failed to save user: " + err.Error())
	}

	fmt.Println("User saved successfully")
}
