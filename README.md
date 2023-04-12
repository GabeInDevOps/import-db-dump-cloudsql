# Import a database dump into CloudSQL using Golang and the Google Cloud SDK

This is a command-line tool written in Golang that imports a database dump into a CloudSQL instance in Google Cloud Platform. The tool uses the Google Cloud SDK for Go and requires a service account key file with the necessary CloudSQL API scopes.

## Prerequisites

Before you can use this tool, you must have the following:

* A Google Cloud project with billing enabled.
* A CloudSQL instance in your project. For more information on creating a CloudSQL instance, see Creating instances.
* A service account with the necessary permissions to import databases into CloudSQL instances. For more information on creating a service account and assigning the appropriate permissions, see Creating and managing service accounts.

## Usage

To use this tool, follow these steps:

* Clone this repository to your local machine

* Navigate to the cloned repository directory:
```bash
cd import-db-dump-cloudsql
```

* Build the tool using the following command:
```go
go build -o import-database import-database.go
```

* Run the tool with the following command, replacing the placeholders with the appropriate values:
```bash
./import-database -project <projectID> -instance <instanceID> -key <keyPath> -client <clientName> -dump <dumpURI>
<projectID>: The ID of your Google Cloud project.
<instanceID>: The ID of your CloudSQL instance.
<keyPath>: The path to your service account key file.
<clientName>: The name of the client database to import the dump into.
<dumpURI>: The URI of the dump file in Google Cloud Storage.
```
For example:

```bash
./import-database -project my-project -instance my-instance -key /path/to/key.json -client my-client-db -dump gs://my-bucket/my-dump.sql
```

Wait for the import operation to complete. The tool will print the operation details once the import completes.

## License

This tool is licensed under the MIT License.