BINARY_NAME=main

build: 
	go build -o ${BINARY_NAME} main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

run-production:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME} 2>&1 > output.log &
