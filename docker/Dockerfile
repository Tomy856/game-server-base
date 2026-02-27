# Build stage
FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy && \
    go build -o main .

# Runtime stage
FROM alpine:3.23

RUN apk add --no-cache ca-certificates git

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]