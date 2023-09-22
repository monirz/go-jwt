FROM golang:1.21-alpine as build

RUN apk add --no-cache curl openssh-keygen  openssl

ENV CGO_ENABLED=0
ENV GO111MODULE=on
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN ssh-keygen -t rsa -b 4096 -m PEM -f /app/keys/private.pem -q -N ""
RUN openssl rsa -in /app/keys/private.pem -pubout -outform PEM -out /app/keys/public.pem
RUN ls keys 

RUN go build -o server cmd/gojwt/main.go

CMD ["./server"]