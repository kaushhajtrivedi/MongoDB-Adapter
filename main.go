package main

import (
	"Hackathon/mongodb"
	"Hackathon/spanner"
	"Hackathon/structure"
	"fmt"
)

func main() {
	fmt.Println("Enter MongoDB URL:")
	var MongodbURL string
	var DatabaseName string
	fmt.Scanln(&MongodbURL)
	fmt.Println("Enter Database Name:")
	fmt.Scanln(&DatabaseName)

	// var MongodbURL string = "mongodb://localhost:27017"
	fmt.Println("Connecting to Mongodb Server ...")
	Client := mongodb.Connection(MongodbURL, DatabaseName)

	fmt.Println("Connecting to Mongodb Database...")
	Coll, collectionNames := mongodb.GetDatabase(Client, DatabaseName)

	fmt.Println("Reading to Mongodb Database...")
	structure.Schema(Coll, Client, collectionNames)
	book := mongodb.Read(Coll)

	spanner.SpannerEnv()
	fmt.Println("Creating Spanner Client ...")
	SpannerClient, adminClient := spanner.CreateClient()
	fmt.Println("Spanner Client Succesfully Created")

	fmt.Println("Creating table in Spanner.... ")
	spanner.Autocreate(collectionNames)
	spanner.CreateTable(adminClient, SpannerClient)
	fmt.Println("Table Created Succesfully")
	spanner.Autowrite(collectionNames)
	spanner.Write(SpannerClient, book)
	fmt.Println("Data Succefully Written to Spanner,Total Entries:")
	fmt.Println("Reading Data from Spanner")
	spanner.Read(SpannerClient)

}
