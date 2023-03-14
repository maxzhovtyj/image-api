FROM golang:1.20-alpine3.17 AS builder

RUN go version
ENV GOPATH=/

COPY . /github.com/maxzhovtyj/image-api/
WORKDIR /github.com/maxzhovtyj/image-api/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/maxzhovtyj/image-api/.bin/main .
COPY --from=0 /github.com/maxzhovtyj/image-api/.env .

EXPOSE 3000

CMD ["./main"]