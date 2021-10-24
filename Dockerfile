# syntax=docker/dockerfile:1
FROM golang:1.17.1-alpine

ENV SVC_NAME=profile

WORKDIR $GOPATH/src/czwr-mailing-${SVC_NAME}/

COPY . .

RUN go mod download

RUN go build -o ./bin/czwr-mailing-${SVC_NAME} ./cmd/${SVC_NAME}/main.go

EXPOSE 8884

ENTRYPOINT ["./bin/czwr-mailing-profile"]