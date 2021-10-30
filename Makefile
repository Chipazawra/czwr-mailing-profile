SVC_NAME=profile
BINARY_NAME=czwr-mailing-${SVC_NAME}
PORT=8884
SIGNING_KEY=SIGNING_KEY
DB_USER=admin
DB_PASS=admin
DB_CLST=czwrmongo.yrzjn.mongodb.net

test:
	go test ./... -v

swag:
	swag init -d ./cmd/${SVC_NAME}/ -o ./doc -g main.go --parseDependency

build:
	go build -o bin/${BINARY_NAME}.exe ./cmd/${SVC_NAME}/main.go

run:
	bin/${BINARY_NAME}.exe -host 0.0.0.0 -port ${PORT} -dbuser ${DB_USER} -dbpass ${DB_PASS} -dbclst ${DB_CLST}

build_and_run: swag build run
	echo "build_and_run"

build_docker:
	docker build --tag czwr-mailing-${SVC_NAME} .