BINARY_NAME=main

build: 
	go build -o ${BINARY_NAME} main.go

run:
	make build
	./${BINARY_NAME}

build-production:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} main.go

run-production:
	./${BINARY_NAME} 2>&1 > output.log &

scp-to-server:
	scp .env.prd root@161.35.247.132:/root/languagequiz/backend/.env
	scp -r migrations root@161.35.247.132:/root/languagequiz/backend
	scp ${BINARY_NAME} root@161.35.247.132:/root/languagequiz/backend
