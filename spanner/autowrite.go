package spanner

import (
	"Hackathon/structure"

	"fmt"
	"log"
	"strings"
)

func Autowrite(collectionNames []string) {
	fileName := fmt.Sprintf("./spanner/Write.go")
	str1 := `package spanner
import(
	"context"
	"time"
	"Hackathon/structure"
	"cloud.google.com/go/spanner"
	"fmt"
	"log"
	"reflect"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
`

	str2 := "\nfunc Write(SpannerClient *spanner.Client, book []interface{})"


	str2 = str2 + "{\n\tctx, _ := context.WithTimeout(context.Background(), 15*time.Second)\n"
	str3 := "\tvar mutations []*spanner.Mutation\n\tvar fields []structure.SpannerField\n\tvar f []string\n\tvar k int\n\tvar field structure.SpannerField\n\tvar bookType reflect.Type\n\tvar bookValue reflect.Value\n\nvar err error\n var fieldValues []interface{}\n"
	for i, name := range collectionNames {

		str3 = str3 + fmt.Sprintf("")
		str3 = str3 + fmt.Sprintf("\tmongoschema%d := structure.%s{}\n", i, strings.Title(name))
		str3 = str3 + fmt.Sprintf("\tfields = extractFieldsFromStruct(mongoschema%d)\n", i)
		str3 = str3 + fmt.Sprintf("\tf = make([]string, len(fields))\n")
		str3 = str3 + fmt.Sprintf("\tfor k, field = range fields {\n")
		str3 = str3 + fmt.Sprintf("\t\tf[k] = field.Name\n\t}\n")
		str3 = str3 + fmt.Sprintf("\tfmt.Println(f, fields, mongoschema%d)\n", i)


		str3 = str3 + fmt.Sprintf("\tbook%d:=book[%d].([]structure.%s)\n", i, i, strings.Title(name))
		str3 = str3 + fmt.Sprintf("\n\tmutations = make([]*spanner.Mutation, len(book%d))\n",i)
		str3 = str3 + fmt.Sprintf("\tfor i,data%d := range book%d{\n", i, i)
		str3 = str3 + fmt.Sprintf("\t\tbookType = reflect.TypeOf(data%d)\n", i)
		str3 = str3 + fmt.Sprintf("\t\tbookValue = reflect.ValueOf(data%d)\n", i)
		str3 = str3 + fmt.Sprintf("\t\tfmt.Println(bookValue,bookType,bookType)\n")
		
		str3 = str3 + "\t\tfieldValues = make([]interface{},  bookType.NumField())\n"
		str3 = str3 + `	
		for j := 0; j < bookType.NumField(); j++ {
			t := bookValue.Field(j).Type()
			switch bookValue.Field(j).Kind() {
			case reflect.Int32:
				fieldValue := int64(bookValue.Field(j).Interface().(int32))
				fieldValues[j] = fieldValue
			case reflect.Float32:
				fieldValue := float64(bookValue.Field(j).Interface().(float32))
				fieldValues[j] = fieldValue
			case reflect.Struct:
				if t == reflect.TypeOf(time.Time{}) {
					fieldValue := bookValue.Field(j).Interface().(time.Time).Format(time.RFC3339)
					fieldValues[j] = fieldValue
				}
			case reflect.String:
				strValue := bookValue.Field(j).Interface().(string)
				// var fieldValues string
				if len(strValue) > 1024 {
					truncatedValue := strValue[:1024] // Truncate the string to 1024 characters
					fieldValues[j] = spanner.NullString{StringVal: truncatedValue, Valid: true}
				} else {
					fieldValues[j] = spanner.NullString{StringVal: strValue, Valid: true}
				}
				// fieldValues[j] = fieldValue
			case reflect.Interface:
				// Handle interface fields
				if t == reflect.TypeOf(primitive.ObjectID{}) {
					fieldValue := bookValue.Field(j).Interface().(primitive.ObjectID).Hex()
					fieldValues[j] = fieldValue
				}

			case reflect.Slice:
				// Handle slice fields
				slice := bookValue.Field(j)
				length := bookValue.Field(j).Len()
				array := make([]interface{}, length)
				for k := 0; k < length; k++ {
					array[k] = slice.Index(k).Interface()
				}
				fieldValues[j] = array
			default:
				// Convert unknown field type to string
				fieldValue := fmt.Sprintf("%v", bookValue.Field(j).Interface())
				fieldValues[j] = fieldValue

			}

		}
		fmt.Println(fieldValues)

`

		str3 = str3 + "\n"
		str3 = str3 + fmt.Sprintf(" mutations[i] = spanner.Insert(")
		str3 = str3 + `"`
		str3 = str3 + name
		str3 = str3 + `"`
		str3 = str3 + fmt.Sprintf(", f, fieldValues)\n\t}\n")
		str3 = str3 + "fmt.Println(mutations)\n"

		// Apply the mutations to Spanner
		str3 = str3 + `	_, err = SpannerClient.Apply(ctx, mutations)
	if err != nil {
		log.Fatal(err)
	}
  
	
	`

		str3 = str3 + "\n"
	}
	str3 = str3 + "\n}"
	err := structure.SaveToFile(fileName, str1+str2+str3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated struct code saved to %s\n", fileName)

}
