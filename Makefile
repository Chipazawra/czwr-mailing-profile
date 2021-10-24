SVC_NAME=profile
BINARY_NAME=czwr-mailing-${SVC_NAME}
PORT=8884

swag:
	swag init -d ./cmd/${SVC_NAME}/ -o ./doc -g main.go --parseDependency

build:
	go build -o bin/${BINARY_NAME}.exe ./cmd/${SVC_NAME}/main.go

run:
	bin/${BINARY_NAME}.exe --host 0.0.0.0 --port ${PORT}

build_and_run: swag build run
	echo "build_and_run"

build_docker:
	docker build --tag czwr-mailing-${SVC_NAME} .