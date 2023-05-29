BINARY_NAME=main

build: 
	go build -o ${BINARY_NAME} main.go

build-production:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} main.go
 
run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}
