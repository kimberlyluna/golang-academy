# Wizeline Academy Golang Workshop

## Explanation

The idea was to create a proxy that will redirect and prioritize requests based on the domain

## Setup

### Run the project

Run the project by using the command `go run main.go` on the root directory. This will compile and run the executable of the program.

Once the server is up and running, start making requests by using the following command `curl -H 'domain: delta' http://localhost:8080/ping`


### Run the tests

Run the tests of the project by using the command `go test ./`on the root directory

### Generate requests

Run the following curl command on any terminal `curl -H 'domain: delta' http://localhost:8080/ping`

