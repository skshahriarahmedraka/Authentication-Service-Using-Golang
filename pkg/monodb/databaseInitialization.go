package monodb

import (
	"fmt"
	"github.com/skshahriarahmedraka/Authentication-Service-Using-Golang/handler"

	"go.mongodb.org/mongo-driver/mongo"
)

func DatabaseInitialization(db1 *mongo.Database) handler.DatabaseCollections {
	fmt.Println("🚀 ~ file: database.go ~ line 15 ~ funcDatabaseInitialization ~ mongodb Database : ", db1)
	return handler.DatabaseCollections{Mongo: db1}
}
