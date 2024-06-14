package main 
import (
	"fmt"
	"mongo_package"
)
func main(){
	 fmt.Println("Connecting into the database")
	 fmt.Println("Creating  a collection into the database")
	 mongo_package.ConnectToMongoDB()
	 fmt.Println("Connecting into  the collection ")
	 mongo_package.InsertedIntoCollection()
	 fmt.Println("Retreiving from the database using the MongoDriver")
	 mongo_package.SelectIntoDataBase()
	 fmt.Println("Updating in the database using the MongoDriver")
	 mongo_package.UpdateIntoTheDatabase()
	 fmt.Println("Deleting from the database")
	 mongo_package.DeleteFromDatabase()
	 
}



