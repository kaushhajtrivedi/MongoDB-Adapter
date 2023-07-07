package mongodb

import (
	"Hackathon/structure"
	"fmt"
	"log"
	"strings"
)

func Autogenerate(collectionName []string) {
	fileName := fmt.Sprintf("./mongodb/reda.go")
	str1 := `package mongodb

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
)`
	str2 := fmt.Sprintf("\nfunc Read(Collection []*mongo.Collection) []interface{}")
	// if len(collectionName) == 1 {
	// 	str2 = str2 + fmt.Sprintf("[]structure.%s", strings.Title(collectionName[0]))
	// } else {
	// 	str2 = str2 + "( "
	// 	for i, name := range collectionName {
	// 		str2 = str2 + fmt.Sprintf("[]structure.%s", strings.Title(name))
	// 		if (i + 1) != len(collectionName) {
	// 			str2 = str2 + ","
	// 		}
	// 	}
	// 	str2 = str2 + ")"
	// }
	str2 = str2 + "{\n"
	str2 = str2 + fmt.Sprintf("\tvar cursors =make([]*mongo.Cursor,%d)\n", len(collectionName))

	str3 := `	var data []interface{}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	
	for i, Coll := range Collection {
		cursor, err := Coll.Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err)
		
		}
	cursors[i]=cursor
	}`
	str4 := "\n\t"
	for i, name := range collectionName {
		str4 = str4 + fmt.Sprintf("\tvar %s []structure.%s\n", name, strings.Title(name))
		str4 = str4 + fmt.Sprintf("\tif err := cursors[%d].All(context.TODO(), &%s); err != nil {\n\tfmt.Println(err)\n\t}\n\tfmt.Println(%s)\n\tfmt.Println(reflect.TypeOf(%s))\n\tdata=append(data,%s)\n\tfmt.Println(data)\n\tfmt.Println(reflect.TypeOf(data[0]))\n", i, name, name, name, name)
	}
	str5 := "\treturn data"
	// for i, name := range collectionName {
	// 	str5 = str5 + fmt.Sprintf(name)
	// 	if (i + 1) != len(collectionName) {
	// 		str5 = str5 + ","
	// 	}
	// }
	str5 = str5 + "\n}"
	err := structure.SaveToFile(fileName, str1+str2+str3+str4+str5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated struct code saved to %s\n", fileName)

}

// func Read(Collection []*mongo.Collection)[] []interface {}{
// 	// book := make([][]interface{},
// 	book := [][]interface{}{
//
// 	str3 := "\n"
// 	for _, name := range collectionName {
// 		str3 = str3 + fmt.Sprintf("\t\t{make([]structure.%s,0)},\n", strings.Title(name))
// 	}
// 	str3 = str3 + "}\n"

// 	str2 := `ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
// 	for i, Coll := range Collection {

// 		cursor, err := Coll.Find(ctx, bson.D{})
// 		// fmt.Print(cursor)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if err = cursor.All(context.TODO(), &book[i]); err != nil {
// 			fmt.Println(err)

// 			// var emp []Student
// 			//defer cursor.Close(ctx)
// 			// for cursor.Next(ctx) {
// 			// 	//
// 			// 	var episode bson.M
// 			// 	if err = cursor.Decode(&episode); err != nil {
// 			// 		log.Fatal(err)
// 			// 	}
// 			// 	fmt.Println("Id:", episode["_id"], "\tName:", episode["Name"], "\t Age:", episode["Age"])

// 			// }
// 			// for _, i := range book[i] {
// 			// 	fmt.Println("ID:", i.ID)
// 			// 	fmt.Println("Name:", i.Name, "\nAge:", i.Age)
// 			// 	fmt.Println("\n------------------------------------------------\n")
// 			// }
// 			fmt.Println(reflect.TypeOf(book[i][0]))
// 		}
// 		}`
// 	str4 := "\n"
// 	for k, name := range collectionName {
// 		str4 = str4 + fmt.Sprintf("jsonData%d, err := json.Marshal(book[%d])\n if err != nil {\n\tlog.Fatal(err)\n}\n", k, k)
// 		str4 = str4 + fmt.Sprintf("book%d := make([]structure.%s,len(book[%d]))\n", k, strings.Title(name), k)
// 		str4 = str4 + fmt.Sprintf("err = bson.Unmarshal(jsonData%d, &book%d)\nif err != nil {\n\tlog.Fatal(err)}\n", k, k)

// 	}

// 	str5 := `

// 	fmt.Println(book)
// 	return book

// }
// `
// 	err := structure.SaveToFile(fileName, str1+str3+str2+str4+str5)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Generated struct code saved to %s\n", fileName)

// }
