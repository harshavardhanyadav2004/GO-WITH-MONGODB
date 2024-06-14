package mongo_package
import (
	"fmt"
	"log"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" 
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateIntoTheDatabase(){
	 clientoptions := options.Client().ApplyURI("mongodb://localhost:27017")
	
	 ctx , cancel :=  context.WithTimeout(context.Background(),10*time.Second)
     client , err := mongo.Connect(ctx , clientoptions)
	 if err!= nil {
		log.Fatal(err)
	 }
	 fmt.Println("The connection id is ",cancel)
	 collection := client.Database("test.db").Collection("Students")
	 fmt.Println("Updating a single field in the database")
	 filter := bson.D{{"name","Harsha"}}
	 cons := bson.D{{"$set",bson.D{{"age",12}}}}
	 result , err:= collection.UpdateOne(ctx,filter ,cons)
	 if err!= nil {
		log.Fatal(err)
	 }
	 fmt.Println("The last updated id is ",result.UpsertedID)
	 fmt.Println("Updating multiple columns in the database ")
	 new_filter := bson.D{{"age",bson.D{{"$lt",5}}}}
	 condition := bson.D{{"age",12}}
	 resulted_id  , err:= collection.UpdateMany(ctx ,new_filter , condition)
	 if err!= nil {
		log.Fatal(err)
		return
	 }
	 fmt.Println("The updated rows",resulted_id.UpsertedCount)
	 
}