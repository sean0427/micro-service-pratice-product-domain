package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name           string             `json:"name" bson:"name"`
	Created        time.Time          `json:"created" bson:"created"`
	Updated        time.Time          `json:"updated" bson:"updated"`
	ManufacturerID primitive.ObjectID `json:"manufacturer_id" bson:"manufacturer_id"`
	CreatedBy      string             `json:"created_by" bson:"created_by"`
}

type Manufacturer struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Created time.Time          `json:"created" bson:"created"`
	Updated time.Time          `json:"updated" bson:"updated"`
}
