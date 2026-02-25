FROM golang:1.23-alpine

WORKDIR /app

# ルートにあるファイルをコンテナにコピー
COPY . .

RUN go mod init game-server && \
    go mod tidy && \
    go build -o main .

CMD ["./main"]