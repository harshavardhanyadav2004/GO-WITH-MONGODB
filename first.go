package mongo_package

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func ConnectToMongoDB() {
	 fmt.Println("Connecting into the MongoFDriver")
	clientoptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("Using the cancel function with the time")
	ctx , cancel :=context.WithTimeout(context.Background(),10*time.Second) 
	fmt.Println("The cancel function is",cancel)

	client,err := mongo.Connect(ctx, clientoptions)
	if err!= nil {
		 log.Fatal(err)
	}

	get_connected:= client.Ping(ctx,nil) 
	if get_connected != nil {
		log.Fatal(get_connected) 
	}
	fmt.Println("Connected to the Mongo Driver succesfully")


}

 