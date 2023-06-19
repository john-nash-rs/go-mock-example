package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestInsertUser(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	//assert.NoError(t, err)
	mock.ExpectQuery(".*").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.0"))

	// Create a GORM database connection with the mock DB
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	defer db.Close()

	// Create a new user
	newUser := &User{
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	// Set up the expectations
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").
		WithArgs(newUser.Name, newUser.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Call the InsertUser function
	err = InsertUser(gormDB, newUser)
	assert.NoError(t, err)
}
