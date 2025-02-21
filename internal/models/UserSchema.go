package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID            bson.ObjectID   `json:"_id" bson:"_id,omitempty"`
	Email         *string         `json:"email" bson:"email" validate:"required"`
	Password      *string         `json:"password" bson:"password" validate:"required"`
	Image         *string         `json:"image" bson:"image"`
	SearchHistory []SearchHistory `json:"searchHistory" bson:"searchHistory"`
}

type SearchHistory struct {
	ID         int            `json:"id" bson:"id"`
	Image      *string        `json:"image" bson:"image"`
	Title      *string        `json:"title" bson:"title"`
	SearchType *string        `json:"searchType" bson:"searchType"`
	CreatedAt  *bson.DateTime `json:"createdAt" bson:"createdAt"`
}
