// mongoHelper.go
package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/MahithChigurupati/go-server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://mahithchigurupati:mahithtest@cluster0.fiehdem.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Movie"
const collectionName = "watchlist"

var Collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection instance created!")
}

func InsertOneMovie(movie models.Movie) {
	insertResult, err := Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}

func UpdateMovie(movieId string) {
	id := movieId
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	updateResult, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func DeleteMovie(movieId string) {
	id := movieId
	filter := bson.M{"_id": id}
	deleteResult, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func DeleteAllMovies() int64 {
	deleteResult, err := Collection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}

func GetAllMovies() []models.Movie {
	cursor, err := Collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var movies []models.Movie

	for cursor.Next(context.Background()) {
		var movie models.Movie
		if err = cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}
