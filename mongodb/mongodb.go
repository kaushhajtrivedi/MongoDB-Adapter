package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection(MongodbURL string, DatabaseName string) *mongo.Client {

	opts := options.Client().ApplyURI(MongodbURL)
	Client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	if err := Client.Database(DatabaseName).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return Client
}

func GetDatabase(Client *mongo.Client, DatabaseName string) ([]*mongo.Collection, []string) {
	filter := bson.M{"name": bson.M{"$type": "string"}}        // Example filter criteria
	listOptions := options.ListCollections().SetNameOnly(true) // Example list options

	collectionNames, err := Client.Database(DatabaseName).ListCollectionNames(context.Background(), filter, listOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Collections in this database:")
	for _, CollectionName := range collectionNames {
		fmt.Println(CollectionName)
	}
	//inserting data
	coll := make([]*mongo.Collection, len(collectionNames))
	for i, CollectionName := range collectionNames {
		coll[i] = Client.Database(DatabaseName).Collection(CollectionName)
		fmt.Println(coll[i])
	}
	Autogenerate(collectionNames)

	return coll, collectionNames
}
