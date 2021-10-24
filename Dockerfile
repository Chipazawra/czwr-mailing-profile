# syntax=docker/dockerfile:1
FROM golang:1.17.1-alpine

WORKDIR $GOPATH/src/czwr-mailing-prorile/

COPY . .

RUN go mod download

RUN go build -o ./bin/czwr-mailing-prorile ./cmd/prorile/main.go

EXPOSE 8885

ENTRYPOINT ["./bin/czwr-mailing-prorile"]