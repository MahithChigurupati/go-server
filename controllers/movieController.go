// movieController.go
package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/MahithChigurupati/go-server/models"
	"github.com/MahithChigurupati/go-server/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllMovies retrieves all movies.
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies := utils.GetAllMovies()
	json.NewEncoder(w).Encode(movies)
}

// GetMovie retrieves a specific movie by ID.
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	filter := bson.M{"_id": id}
	var movie models.Movie
	err := utils.Collection.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(movie)
}

// CreateMovie creates a new movie.
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	utils.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

// MarkAsWatched updates a movie to mark it as watched.
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	utils.UpdateMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteMovie deletes a movie.
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	utils.DeleteMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteAllMovies deletes all movies.
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	count := utils.DeleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
