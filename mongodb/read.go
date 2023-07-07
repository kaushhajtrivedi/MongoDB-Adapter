package mongodb

// import (
// 	"context"
// 	"fmt"

// 	// "Hackathon/structure"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// // read

// func Read(Collection []*mongo.Collection) [][]interface{} {
// 	// book := make([][]interface{}, len(Collection))
// 	var book = make([][]interface{}, len(Collection))
// 	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
// 	for i, Coll := range Collection {

// 		cursor, err := Coll.Find(ctx, bson.D{})
// 		// fmt.Print(cursor)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if err = cursor.All(context.TODO(), &book[i]); err != nil {
// 			fmt.Println(err)
// 		}
// 		// var emp []Student
// 		//defer cursor.Close(ctx)
// 		// for cursor.Next(ctx) {
// 		// 	//
// 		// 	var episode bson.M
// 		// 	if err = cursor.Decode(&episode); err != nil {
// 		// 		log.Fatal(err)
// 		// 	}
// 		// 	fmt.Println("Id:", episode["_id"], "\tName:", episode["Name"], "\t Age:", episode["Age"])

// 		// }
// 		// for _, i := range book[i] {
// 		// 	fmt.Println("ID:", i.ID)
// 		// 	fmt.Println("Name:", i.Name, "\nAge:", i.Age)
// 		// 	fmt.Println("\n------------------------------------------------\n")
// 		// }
// 	}
// 	//fmt.Println(book)
// 	return book

// }
