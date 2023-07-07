package mongodb

import (
	"context"
	"fmt"
	// "encoding/json"
	// "Hackathon/structure"
	"reflect"
	"log"
	"time"
	"Hackathon/structure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
func Read(Collection []*mongo.Collection) []interface{}{
	var cursors =make([]*mongo.Cursor,2)
	var data []interface{}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	
	for i, Coll := range Collection {
		cursor, err := Coll.Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err)
		
		}
	cursors[i]=cursor
	}
		var book1 []structure.Book1
	if err := cursors[0].All(context.TODO(), &book1); err != nil {
	fmt.Println(err)
	}
	fmt.Println(book1)
	fmt.Println(reflect.TypeOf(book1))
	data=append(data,book1)
	fmt.Println(data)
	fmt.Println(reflect.TypeOf(data[0]))
	var book []structure.Book
	if err := cursors[1].All(context.TODO(), &book); err != nil {
	fmt.Println(err)
	}
	fmt.Println(book)
	fmt.Println(reflect.TypeOf(book))
	data=append(data,book)
	fmt.Println(data)
	fmt.Println(reflect.TypeOf(data[0]))
	return data
}