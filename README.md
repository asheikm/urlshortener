[![Go Report Card](https://goreportcard.com/badge/github.com/asheikm/urlshortener)](https://goreportcard.com/report/github.com/asheikm/urlshortener)

## URLShortener

URL Shortener App - This application shortens a given URL and stores the data in a flat file.
**Note:** This application is not production-ready and was developed for learning purposes.

## Getting Started

To run this app locally:

1. Clone this repository.
2. Install the Go programming language ([instructions](https://golang.org/doc/install)).
3. Install Docker ([instructions](https://docs.docker.com/engine/install/)).

## Basic Configuration

This app uses environment variables for configuration. To run the app as a standalone application, set the following environment variables:

```bash
   export LISTEN_PORT=8080`
   export LOGFILE_PATH="./log/log.out"
   export SHORT_DOMAIN="localhost:8080/"
```

## Build and run

Compile and generate binary:

- How to compile and generate binary
   - Put the code in GOPATH which is usually ($HOME/go/src)
   - Optional: Run `go install ` will get the dependency packages (Go mod take care of this anyway)
   - Run `go build -o bin/urlshortener main.go `(This will generate binary file name urlshortener in bin directory this will later copied to docker image)
   - Makefile has be available with the codebase ( For windows users may not use this directly, need tools like cygwin or similar cross compiler toolsets)
   - If you want to build using make file , just run the command `make build`

Run Urlshortener as standalone app:

   - Run the binary file urlshortener ./urlshortener from cli which will listen of default port 8080
   - To run the app in the background `./urlshortener &` (This only works on linux/unix flavor default, make sure you have necessary tools in windows )

Building and runnning app from docker image:

   - Please make sure you built the binary either by runnign `make build`  cmd or running `go build -o bin/urlshortener main.go` from source path
   - To build into docker image `docker build . -t urlshortener:1.0 `
   - To list images from the local docker hub repo `docker images`
   - To run docker image ` docker run -it -d -p 8080:8080 urlshortener:1.0` (This will run urlshortener app and listen on port 8080)
   - Note : The image has been build using centos, not using golang image, since we can directly copy the binary to docker image than building it in the docker image
            which will expose the source if forget to delete.

## Unit Test

Unit Testing
    
   - This app uses go programming languages default test framework 
   - Run `go test ./...` from project directory, this command will get the result something similar below 
     
     `[user@dev utils]$ go test -v ./...`

 
## Application Usage
   
Curl command to generate shortened url
   
  - Run `curl -X POST -d '{"url": "google.com"}'  http://localhost:8080/shrink --header "Content-Type:application/json"` which will get you the shortend url 
  - Run `curl -X GET -H "Content-type: application/json" \-H "Accept: application/json" -d '{"url":"google.com"}' "http://localhost:8080/shrink"` (This will fectch if the shortened url is present inmemory or flatfile
  - Run ` curl -v -X GET -H "Content-type: application/json" \-H "Accept: application/json" -d '{"url":"http://localhost/<shortcode>"}' "http://localhost:8080/redirect"` to redirect given url 
  - We can also use postman or similar tools to check GET/POST methods described above.

## Known issues
  - Please check github repo issues

## Key Updates:
1. **Clarity:** Improved language for better understanding.
2. **Consistency:** Maintained a consistent format across sections.
3. **Readability:** Simplified instructions and organized content for easier navigation.


