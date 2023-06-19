package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSaveUser(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Expect the query for the database version
	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.0"))
	// Initialize a new GORM database connection with the mock database
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open GORM connection: %v", err)
	}

	// Mock the expected query and define its behavior
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Create a new user to save
	newUser := &User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	// Call the SaveUser method
	err = SaveUser(gormDB, newUser)
	if err != nil {
		t.Fatalf("Failed to save user: %v", err)
	}

	// Verify that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestSaveUser_InvalidEmail(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Expect the query for the database version
	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.0"))

	// Initialize a new GORM database connection with the mock database
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open GORM connection: %v", err)
	}

	// Create a new user with an invalid email address
	newUser := &User{
		Name: "John Doe",
	}

	// Call the SaveUser method
	err = SaveUser(gormDB, newUser)
	expectedError := "Invalid email"
	assert.Equal(t, expectedError, err.Error(), "In case of invalid email, message is wrong")

	// Verify that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
