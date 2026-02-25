# 1. 標準的なOS（Debian）を使用。これだけでVSCodeは動きます
FROM golang:1.21-bookworm

# 2. 作業ディレクトリの設定
WORKDIR /app

# 3. ファイルのコピー
COPY . .

# 4. ビルド処理
RUN go mod tidy && \
    go build -o main .

# 5. 実行
CMD ["./main"]