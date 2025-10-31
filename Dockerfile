FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
ENV DB_PATH = $PWD

COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /root/
ENV DB_PATH = $PWD
COPY --from=builder /app/main .
COPY --from=builder /app/Vault ./Vault
CMD ["./main"]
