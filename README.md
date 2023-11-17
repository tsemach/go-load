### Go Workspaces
https://bysabbir.medium.com/go-workspaces-simplifying-multi-modular-projects-dc1a489302a

## Running the restapi server
````bash
go run main.go
````

## Calling health api
````bash
curl -k --cert certs/client.crt --key certs/client.key --cacert certs/ca.crt https://localhost:8080/health
````