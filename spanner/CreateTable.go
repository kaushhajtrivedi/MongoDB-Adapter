package spanner
import(
"fmt"
"Hackathon/structure"
adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"cloud.google.com/go/spanner"
	"context"
)
func CreateTable(adminClient *database.DatabaseAdminClient, spannerClient *spanner.Client) error {
	
var op *database.UpdateDatabaseDdlOperation
var ddlString string
var fields []structure.SpannerField
var err error

book0:= structure.Book1{}
fields = extractFieldsFromStruct(book0)
ddlString = generateddl(fields,"book1")
	
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
			
		}
book1:= structure.Book{}
fields = extractFieldsFromStruct(book1)
ddlString = generateddl(fields,"book")
	
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
			
		}
return nil
}