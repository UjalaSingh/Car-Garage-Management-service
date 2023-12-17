## Basic Installation  - 
- clone from this github repo 
- make sure you have go language installed in your system (https://go.dev/doc/install)
- make sure you have docker installed (https://www.docker.com/products/docker-desktop/)
- go into the cloned folder and run this command
   > go mod download

## Start the Project - 
- start docker desktop or docker
- then run this command - 
  > docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30

- if you are having trouble creating an image because of a port try changing the port number in the docker command and .env file in the config folder.
- run this command `go run main.go`
- you can test all the APIs using postman [download and test API](/Car%20API%20collection.postman_collection.json) 

# to run unit tests, do -
- make sure to close the previous program that is being run using `go run main.go`
- run `go test ./...` or alternatively run `go test ./... -v`


 
