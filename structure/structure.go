package structure

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"
)

// StructTemplate defines the template used to generate the Go struct code.
const StructTemplate = `
type {{ .StructName }} struct {
{{ range $fieldName, $fieldType := .Fields }}
	{{ $fieldName }} {{ $fieldType }}
{{ end }}
}
`

// StructData represents the data for the struct template.
type StructData struct {
	StructName string
	Fields     map[string]string
}

func Schema(collections []*mongo.Collection, client *mongo.Client, collectionNames []string) {
	// MongoDB connection configuration
	for i, collection := range collections {
		// Fetch a sample document from the collection
		sampleDocument := bson.M{}
		fmt.Println(collection)
		err := collection.FindOne(context.Background(), bson.D{}).Decode(&sampleDocument)
		if err != nil {
			log.Fatal(err)
		}

		// Generate the struct name
		structName := strings.Title(collectionNames[i])

		// Generate the field names and types
		fields := make(map[string]string)
		for fieldName, fieldValue := range sampleDocument {
			fieldName = replaceFieldID(fieldName)
			fieldType := reflect.TypeOf(fieldValue).String()
			fields[fieldName] = fieldType
		}
		delete(fields, "_id")

		// Add the "ID" field
		fields["ID"] = "primitive.ObjectID `bson:\"_id\"`"

		// Generate the Go struct code
		structData := StructData{
			StructName: structName,
			Fields:     fields,
		}
		structCode, err := GenerateStructCode(structData)
		if err != nil {
			log.Fatal(err)
		}
		extratext := "package structure\n import (\"go.mongodb.org/mongo-driver/bson/primitive\")\n"
		// Save the struct code to a file
		fileName := fmt.Sprintf("./structure/%s_struct.go", strings.ToLower(structName))
		err = SaveToFile(fileName, extratext+structCode)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Generated struct code saved to %s\n", fileName)
	}
}

// replaceFieldID replaces the field name "_id" with "ID".
func replaceFieldID(FieldName string) string {
	if FieldName == "_id" {
		return "ID"
	}
	return FieldName
}

// generateStructCode generates the Go struct code using the provided struct data.
func GenerateStructCode(data StructData) (string, error) {
	tmpl, err := template.New("struct").Parse(StructTemplate)
	if err != nil {
		return "", err
	}

	var structCodeBuilder strings.Builder
	err = tmpl.Execute(&structCodeBuilder, data)
	if err != nil {
		return "", err
	}

	return structCodeBuilder.String(), nil
}

// saveToFile saves the provided content to a file with the specified name.
func SaveToFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

