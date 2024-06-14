package mongo_package

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type new_person struct {
	 id int `bson:"id"`
	 name string `bson:"name"`
}
func InsertedIntoCollection() {
	 fmt.Println("Connecting into the mongoDriver using the go lang")
	 clientoptions := options.Client().ApplyURI("mongodb://localhost:27017")

	 ctx , cancelfunc := context.WithTimeout(context.Background(),10*time.Second)
	
	 client , err := mongo.Connect(ctx ,clientoptions)

	 if err!= nil {
		log.Fatal(err)
	 }
	fmt.Println("Connected to the MongoDriver succesffully",cancelfunc)

	fmt.Println("Connecting to the collection using MongoDB")
	collection := client.Database("test.db").Collection("students")
	another_array :=new_person{id : 1 , name : "Harsha"}
	id , err:= collection.InsertOne(ctx , another_array)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The last inserted id in the collection is ",id.InsertedID)
}