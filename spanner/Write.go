package spanner
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

func Write(SpannerClient *spanner.Client, book []interface{}){
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	var mutations []*spanner.Mutation
	var fields []structure.SpannerField
	var f []string
	var k int
	var field structure.SpannerField
	var bookType reflect.Type
	var bookValue reflect.Value

var err error
 var fieldValues []interface{}
	mongoschema0 := structure.Book1{}
	fields = extractFieldsFromStruct(mongoschema0)
	f = make([]string, len(fields))
	for k, field = range fields {
		f[k] = field.Name
	}
	fmt.Println(f, fields, mongoschema0)
	book0:=book[0].([]structure.Book1)

	mutations = make([]*spanner.Mutation, len(book0))
	for i,data0 := range book0{
		bookType = reflect.TypeOf(data0)
		bookValue = reflect.ValueOf(data0)
		fmt.Println(bookValue,bookType,bookType)
		fieldValues = make([]interface{},  bookType.NumField())
	
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


 mutations[i] = spanner.Insert("book1", f, fieldValues)
	}
fmt.Println(mutations)
	_, err = SpannerClient.Apply(ctx, mutations)
	if err != nil {
		log.Fatal(err)
	}
  
	
	
	mongoschema1 := structure.Book{}
	fields = extractFieldsFromStruct(mongoschema1)
	f = make([]string, len(fields))
	for k, field = range fields {
		f[k] = field.Name
	}
	fmt.Println(f, fields, mongoschema1)
	book1:=book[1].([]structure.Book)

	mutations = make([]*spanner.Mutation, len(book1))
	for i,data1 := range book1{
		bookType = reflect.TypeOf(data1)
		bookValue = reflect.ValueOf(data1)
		fmt.Println(bookValue,bookType,bookType)
		fieldValues = make([]interface{},  bookType.NumField())
	
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


 mutations[i] = spanner.Insert("book", f, fieldValues)
	}
fmt.Println(mutations)
	_, err = SpannerClient.Apply(ctx, mutations)
	if err != nil {
		log.Fatal(err)
	}
  
	
	

}