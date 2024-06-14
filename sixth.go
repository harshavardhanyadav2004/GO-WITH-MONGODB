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
func DeleteFromDatabase(){
	fmt.Println("Connecting into the database")
	clientoptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx , cancel := context.WithTimeout(context.Background(),10*time.Second)
	
	client , err := mongo.Connect(ctx , clientoptions)
	if err!= nil {
		log.Fatal(err)
		return 
	}
	fmt.Println("The data base id is ",cancel)
	collection := client.Database("test.db").Collection("Students")
	deleted_id , err:= collection.DeleteOne(ctx ,bson.D{{Key :"name",Value : "Harsha"}})
	if err!= nil {
	 log.Fatal(err)
	 return 
	}
	fmt.Println("The deleted rows are ",deleted_id.DeletedCount)
	fmt.Println("Deleting multiple rows from the database")
	new_collection , err := collection.DeleteMany(ctx , bson.D{{Key :"age",Value : bson.D{{Key : "$lt",Value : 5}}}})
	if err!= nil {
		log.Fatal(err)
		return 
	}
	fmt.Println("The collections updated are ",new_collection.DeletedCount)

}