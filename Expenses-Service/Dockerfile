FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
CMD ["./main"]
