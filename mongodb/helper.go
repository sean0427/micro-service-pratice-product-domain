package mongodb

import (
	"time"

	"github.com/sean0427/micro-service-pratice-product-domain/model"
	"go.mongodb.org/mongo-driver/bson"
)

func ToBson(item model.Product) bson.M {
	if item.Created.IsZero() {
		item.Created = time.Now()
	}

	return bson.M{
		"name":            item.Name,
		"manufacturer_id": item.ManufacturerID,
		"updated":         time.Now(),
		"created":         item.Created,
		"Create_by":       "TODO"}
}
