FROM golang:alpine AS builder

RUN apk add --no-cache gcc musl-dev

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o idea-shop-api main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/idea-shop-api .

COPY .env .env

EXPOSE 3000

CMD ["./idea-shop-api"]