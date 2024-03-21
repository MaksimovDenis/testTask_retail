FROM golang:1.21.6-alpine AS builder

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apk update && apk add --no-cache postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o testtask_retail ./cmd/main.go


CMD ["./testtask_retail"]