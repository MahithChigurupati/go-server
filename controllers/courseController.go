// courseController.go
package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/MahithChigurupati/go-server/models"
	"github.com/MahithChigurupati/go-server/utils"
	"github.com/gorilla/mux"
)

// GetAllCourses retrieves all courses.
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetCourses")
	w.Header().Set("Content-Type", "application/json")
	var courses []models.Course
	models.DB.Find(&courses)
	json.NewEncoder(w).Encode(courses)
}

// GetCourse retrieves a specific course by ID.
func GetCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetCourse")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var course models.Course
	result := models.DB.Where("course_id = ?", params["id"]).First(&course)
	if result.Error != nil {
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(course)
}

// CreateCourse creates a new course.
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateCourse")
	w.Header().Set("Content-Type", "application/json")

	var newCourse models.Course
	err := json.NewDecoder(r.Body).Decode(&newCourse)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if utils.IsCourseEmpty(newCourse) {
		http.Error(w, "Course data is empty", http.StatusBadRequest)
		return
	}

	if models.IsDuplicateCourseName(newCourse.CourseName) {
		http.Error(w, "Course name already exists", http.StatusBadRequest)
		return
	}

	newCourse.CourseID = strconv.Itoa(rand.Intn(100000000))
	models.DB.Create(&newCourse)

	json.NewEncoder(w).Encode(newCourse)
}

// UpdateCourse updates a course.
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateCourse")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var course models.Course
	result := models.DB.Where("course_id = ?", params["id"]).First(&course)
	if result.Error != nil {
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	models.DB.Save(&course)

	json.NewEncoder(w).Encode(course)
}

// DeleteCourse deletes a course.
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteCourse")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	result := models.DB.Where("course_id = ?", params["id"]).Delete(&models.Course{})
	if result.RowsAffected == 0 {
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode("Course deleted")
}

// DeleteAllCourses deletes all courses.
func DeleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteAllCourses")
	w.Header().Set("Content-Type", "application/json")

	result := models.DB.Delete(&models.Course{})
	if result.RowsAffected == 0 {
		http.Error(w, "Courses not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode("All courses deleted")
}
