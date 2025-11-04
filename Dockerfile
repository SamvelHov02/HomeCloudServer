FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
ENV DB_PATH=/app

COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
ENV DB_PATH=/app
COPY --from=builder /app/main .
COPY --from=builder /app/Vault ./Vault
CMD ["./main"]
