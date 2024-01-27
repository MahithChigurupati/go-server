// courseRoutes.go
package routes

import (
	controller "github.com/MahithChigurupati/go-server/controllers"
	"github.com/gorilla/mux"
)

func CourseRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/courses", controller.GetAllCourses).Methods("GET")
	router.HandleFunc("/api/courses/{id}", controller.GetCourse).Methods("GET")
	router.HandleFunc("/api/courses", controller.CreateCourse).Methods("POST")
	router.HandleFunc("/api/courses/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/courses/{id}", controller.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/api/courses", controller.DeleteAllCourses).Methods("DELETE")

	return router
}
