package mongo_package 
import (
	"log" 
	"fmt"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" 
	"go.mongodb.org/mongo-driver/bson"
)
type person struct {
	 id int `bson:"name"`
	 name string `bson:"id"`
}

func SelectIntoDataBase(){
	 fmt.Println("Connecting into the database")
	 clientoptions := options.Client().ApplyURI("mongodb://localhost:27017")
	 ctx , cancel := context.WithTimeout(context.Background(),10*time.Second)
	 fmt.Println("Conmnecting into the collection")
	 client , err:= mongo.Connect(ctx,clientoptions)
	 if err!= nil {
		log.Fatal(err)
	 }
	 var result person
	 fmt.Println("The connection id is",cancel)
	 collection := client.Database("testdb").Collection("Students")
	 fmt.Println("Retreibving single data from the database")
	 new_err:= collection.FindOne(ctx,bson.D{{Key:"name",Value:"Harsha"}}).Decode(&result)
	 if new_err != nil {
		log.Fatal(err)
	 }
	 fmt.Println(result)
	 fmt.Println("Retreiving the multiple data from the database")
	 cursor , err:= collection.Find(ctx ,bson.D{})
	if err!= nil {
		log.Fatal(err)
	}
	var results[] person
	for cursor.Next(ctx){
		 var new_string person
		 cursor.Decode(&new_string)
		 results = append(results, new_string)
	}
	fmt.Println("Retreiving multiple documents from the database")
	fmt.Println(results)
}