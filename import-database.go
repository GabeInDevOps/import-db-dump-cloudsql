package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/sqladmin/v1"
)

func main() {
	ctx := context.Background()

	// Parse command-line flags.
	projectID := flag.String("project", "", "Google Cloud project ID")
	instanceID := flag.String("instance", "", "CloudSQL instance ID")
	keyPath := flag.String("key", "", "Path to service account key file")
	clientName := flag.String("client", "", "Name of the client database")
	dumpURI := flag.String("dump", "", "URI of the database dump file in Google Cloud Storage")
	flag.Parse()

	// Validate command-line arguments.
	if *projectID == "" || *instanceID == "" || *keyPath == "" || *clientName == "" || *dumpURI == "" {
		fmt.Println("Usage: import-database -project <projectID> -instance <instanceID> -key <keyPath> -client <clientName> -dump <dumpURI>")
		return
	}

	// Create a new CloudSQL API client using the service account key file.
	client, err := sqladmin.NewService(ctx, option.WithCredentialsFile(*keyPath))
	if err != nil {
		fmt.Printf("Failed to create CloudSQL client: %v\n", err)
		return
	}

	// Configure the import request.
	importReq := &sqladmin.InstancesImportRequest{
		ImportContext: &sqladmin.ImportContext{
			FileType: "SQL",
			Uri:      *dumpURI,
			Database: *clientName,
		},
	}

	// Call the Import method to start the import operation.
	op, err := client.Instances.Import(*projectID, *instanceID, importReq).Do()
	if err != nil {
		fmt.Printf("Failed to start import operation: %v\n", err)
		return
	}

	// Wait for the operation to complete.
	op, err = waitForOperation(ctx, client, *projectID, op.Name)
	if err != nil {
		fmt.Printf("Import operation failed: %v\n", err)
		return
	}

	fmt.Printf("Import operation completed successfully: %v\n", op)
}

func waitForOperation(ctx context.Context, client *sqladmin.Service, projectID, opName string) (*sqladmin.Operation, error) {
	for {
		op, err := client.Operations.Get(projectID, opName).Do()
		if err != nil {
			return nil, fmt.Errorf("Failed to get operation status: %v", err)
		}
		if op.Status == "DONE" {
			return op, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Second * 5):
		}
	}
}
