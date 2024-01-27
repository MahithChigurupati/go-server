// movieModel.go
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Rating  string             `json:"rating,omitempty" bson:"rating,omitempty"`
	Watched bool               `json:"watched,omitempty" bson:"watched,omitempty"`
}

// Sample JSON data
//
// Request body for POST request
//
// {
// 	"name": "The Shawshank Redemption",
// 	"rating": "9.3",
// 	"watched": false
// }

// Response body for GET request

// {
// 	"_id": "5fecf357c2a54032a04e3ac5",
// 	"name": "The Shawshank Redemption",
// 	"rating": "9.3",
// 	"watched": false
// }
