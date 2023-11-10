package monodb

import (
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

func DatabaseInitialization(db1 *mongo.Database) handler.DatabaseCollections {
	return handler.DatabaseCollections{Mongo: db1}
}
