package monodb

import (
	// "app/controller"
	// "app/LogError"
	"context"
	"fmt"
	"time"

	"os"

	// "log"

	// "go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongodbConnection() *mongo.Database {

	
	// return Mydb
	uri:= fmt.Sprintf("mongodb://%v:%v@%v:%v/?maxPoolSize=%v&w=majority" ,os.Getenv("MONGO_USER"),os.Getenv("MONGO_PASSWORD"),os.Getenv("MONGO_IP"),os.Getenv("MONGO_PORT"),os.Getenv("MONGO_MAXPOOLSIZE"))
    // fmt.Println("🚀 ~ file: databaseSetup.go ~ line 13 ~ funcDBSetup ~ mongouri : ", mongouri)
	// client ,err :=mongo.NewClient(options.Client().ApplyURI(mongouri))
    // fmt.Println("🚀 ~ file: databaseSetup.go ~ line 17 ~ funcDBSetup ~ err : ",err)
	
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	ctx,cancel:=context.WithTimeout(context.Background(), 10* time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	

	// err =client.Connect(ctx)
	// fmt.Println("❌ ~ file: databaseSetup.go ~ line 25 ~ funcDBSetup ~ err : ", err)
	err =client.Ping(context.TODO(),nil)
    // fmt.Println("❌ ~ file: databaseSetup.go ~ line 27 ~ funcDBSetup ~ err : ", err)
	Mydb := client.Database("userdb")
	if err== nil {

		fmt.Println("⚡😍 sucessfully connected to database")
	}
	return Mydb

}


