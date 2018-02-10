

## Prerequisite
This project requires the following dependencies:

- Go - version 1.8 and above
- [dep](https://github.com/golang/dep) - Dependency package manager


## Running the application

To the run the application execute the following commands:

1. `dep ensure` - This will download all necessary dependencies
2. `go run main` - Run the application (runs on port 8080)


## Running tests
To run the test, execute the following command to run all tests including
the ones in the sub folders: `go test ./...`

## Endpoints
This service has the following endpoints:
1. /api/v1/createInspiration (POST)
2. /api/v1/healthcheck (GET)



## Sample payload:
The following bash command below will make a valid request and save the response into a file called image.png.
```bash
curl  -H "Content-Type: application/json" -d @sample-payload.json localhost:8080/api/v1/createInspiration > image.png
```


## Limitations/Technical Debt
The following limitations and technical debt remains due to time constraints:
- Supports only the following MIME types: image/png and image/jpeg due to time restrictions
- Integration tests hit external services rather than having a full contained test using Docker containers