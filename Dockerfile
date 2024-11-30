FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app ./cmd

FROM golang:1.23-alpine

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

ENTRYPOINT ["./app"]

EXPOSE 8080
