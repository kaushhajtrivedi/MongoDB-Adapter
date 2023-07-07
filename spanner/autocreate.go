package spanner

import (
	"Hackathon/structure"
	"fmt"
	"log"
	"strings"
)

func Autocreate(collectionNames []string) {
	fileName := fmt.Sprintf("./spanner/CreateTable.go")
	str1 := `package spanner
import(
"fmt"
"Hackathon/structure"
adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"cloud.google.com/go/spanner"
	"context"
)
func CreateTable(adminClient *database.DatabaseAdminClient, spannerClient *spanner.Client) error {
	`
	str2 := "\nvar op *database.UpdateDatabaseDdlOperation\nvar ddlString string\nvar fields []structure.SpannerField\nvar err error\n"

	for i, j := range collectionNames {
		str2 = str2 + fmt.Sprintf("\nbook%d:= structure.%s{}\nfields = extractFieldsFromStruct(book%d)\nddlString = generateddl(fields,", i, strings.Title(j), i)
		str2 = str2 + `"` + j + `")`
		str2 = str2 + "\n"
		str2 = str2 + `	
		fmt.Println(ddlString)
		fmt.Println("ddl string:", ddlString, "database name:", spannerClient.DatabaseName())

		// Execute the DDL statement.
		op, err = adminClient.UpdateDatabaseDdl(context.Background(), &adminpb.UpdateDatabaseDdlRequest{
			Database:   spannerClient.DatabaseName(),
			Statements: []string{ddlString},
		})

		fmt.Println("hii")

		if err != nil {
			return fmt.Errorf("failed to execute DDL statement: %v", err)
		}

		if err := op.Wait(context.Background()); err != nil {
			return fmt.Errorf("failed to wait for DDL operation completion: %v", err)
			
		}`

	}
	str2 = str2 + "\nreturn nil\n}"

	err := structure.SaveToFile(fileName, str1+str2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated struct code saved to %s\n", fileName)

}
