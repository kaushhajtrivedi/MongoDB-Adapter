package mongodbtospanner

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTypeName(t reflect.Type) string {
	// var spannerType string
	switch t.Kind() {
	case reflect.String:
		return "STRING(1024)"
	case reflect.Int, reflect.Int32, reflect.Int64:
		return "INT64"
	case reflect.Float32, reflect.Float64:
		return "FLOAT64"
	case reflect.Bool:
		return "BOOL"
	case reflect.Struct:
		if t == reflect.TypeOf(time.Time{}) {
			return "TIMESTAMP"
		}
	case reflect.Slice:
		sliceElemType := t.Elem().Kind()
		if sliceElemType == reflect.String {
			return "ARRAY<STRING(1024)>"
		}
		// Add support for other slice element types as needed
	case reflect.Ptr:
		pointerElemType := t.Elem().Kind()
		if pointerElemType == reflect.String {
			return "STRING(1024)"
		}
		// Add support for other nullable types as needed
	case reflect.Interface:
		return "JSON"
	// Handle other data types as per your specific requirements
	default:
		if t == reflect.TypeOf(primitive.ObjectID{}) {
			return "STRING(1024)"
		} else if t == reflect.TypeOf([]byte{}) {
			return "BYTES"
		} else {
			return "STRING(1024)" // Default to STRING if type not recognized
		}
		//

	}
	return "STRING(1024)"
}
