FROM golang:1.21-alpine as build

RUN apk add --no-cache curl

ENV CGO_ENABLED=0
ENV GO111MODULE=on
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server cmd/gojwt/main.go

FROM alpine:latest

COPY --from=build /app /app

CMD ["./server"]