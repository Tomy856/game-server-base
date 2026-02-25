FROM golang:1.26-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy && \
    go build -o main .

CMD ["./main"]