package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// Global database connection that will be shared across all requests
	globalDB *gorm.DB
)

// InitDB initializes the database connection
func InitDB() (*gorm.DB, error) {
	// Use a shared in-memory database with connection pooling
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Store the database connection globally
	globalDB = db
	return db, nil
}

// GetDB returns the global database connection
func GetDB() *gorm.DB {
	return globalDB
}
