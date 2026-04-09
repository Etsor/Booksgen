FROM golang:1.25-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./...

RUN go build -ldflags="-s -w" -o /booksgen ./cmd/

FROM alpine:latest

COPY --from=builder /booksgen /usr/local/bin/booksgen

EXPOSE 8080

ENTRYPOINT ["booksgen"]
