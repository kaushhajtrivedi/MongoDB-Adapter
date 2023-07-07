package spanner

import (
	"Hackathon/mongodbtospanner"
	"Hackathon/structure"
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
	"cloud.google.com/go/spanner"
	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"google.golang.org/api/iterator"
)



func SpannerEnv() error {
	if err := os.Setenv("SPANNER_EMULATOR_HOST", "localhost:9010"); err != nil {
		fmt.Println("Error while setting environment variable: ", err)
		return err
	}
	return nil
}



func CreateClient() (*spanner.Client, *database.DatabaseAdminClient) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	Client, err := spanner.NewClient(ctx, "projects/spanner-project/instances/spanner-instance/databases/spanner-database")
	if err != nil {
		log.Fatal(err)
	}
	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Spanner admin client: %v", err)
	}
	return Client, adminClient
}



func extractFieldsFromStruct(data interface{}) []structure.SpannerField {
	var fields []structure.SpannerField
	t := reflect.TypeOf(data)
	fmt.Println(t)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		fieldType := mongodbtospanner.GetTypeName(field.Type)

		fields = append(fields, structure.SpannerField{Name: fieldName, Type: fieldType})
	}
	return fields
}



func generateddl(fields []structure.SpannerField, collectionName string) string {
	sqlSchema := "CREATE TABLE " + collectionName + " (\n"
	for _, field := range fields {
		sqlSchema += fmt.Sprintf("\t\t\t%s %s,\n", field.Name, field.Type)
	}
	sqlSchema = sqlSchema[:len(sqlSchema)-2] + "\n\t) PRIMARY KEY (ID)"
	return sqlSchema
}



func Read(Client *spanner.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	tableName := "book"
	columns := []string{"Name", "Age"}
	stmt := spanner.Statement{SQL: fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), tableName)}
	iter := Client.Single().Query(ctx, stmt)
	defer iter.Stop()
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var (
			col1 string
			col2 int64
		)
		if err := row.Columns(&col1, &col2 /* add more variables for other columns */); err != nil {
			log.Fatal(err)
		}
		fmt.Print("Name:", col1)
		fmt.Println("	Age:", col2)
	}
}
