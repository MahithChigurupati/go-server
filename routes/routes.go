// routes.go
package routes

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	mainRouter := mux.NewRouter()
	mainRouter.PathPrefix("/api/movies").Handler(MovieRouter())
	mainRouter.PathPrefix("/api/courses").Handler(CourseRouter())
	return mainRouter
}
