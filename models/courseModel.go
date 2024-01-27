// courseModel.go
package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	CourseID    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	Description string  `json:"description"`
	CoursePrice float32 `json:"course_price"`
	Author      Author  `json:"author" gorm:"foreignkey:AuthorID"`
	AuthorID    uint    // Foreign key
}

type Author struct {
	gorm.Model
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

// Sample JSON data
//
// Request body for POST request
//
// {
// 	"course_id": "c123",
// 	"course_name": "Introduction to Go Programming",
// 	"description": "A comprehensive course on Go programming language.",
// 	"course_price": 49.99,
// 	"author": {
// 		"full_name": "John Doe",
// 		"website": "https://www.johndoe.com"
// 	}
// }
