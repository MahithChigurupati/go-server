// models.go
package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitializeDB initializes the database connection.
func InitializeDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=goserver port=5432 sslmode=disable"
	// Replace the connection details with your PostgreSQL database connection string.

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("PG Database connection successfully opened")

	// Auto Migrate
	DB.AutoMigrate(&Course{}, &Author{})
}

func IsDuplicateCourseName(courseName string) bool {
	var count int64
	DB.Model(&Course{}).Where("course_name = ?", courseName).Count(&count)
	return count > 0
}
