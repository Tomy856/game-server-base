# デプロイメント

本プロジェクトの本番環境への展開方法を説明します。

## 🐳 Docker

### ビルド

```bash
docker build -f docker/Dockerfile -t game-server:latest .
```

### ローカル実行

```bash
docker-compose up
```

## 📦 バイナリビルド

```bash
# 開発環境
go build -o bin/game-server main.go

# クロスコンパイル（Linux）
GOOS=linux GOARCH=amd64 go build -o bin/game-server-linux main.go
```

## 🚀 CI/CD

### GitHub Actions（例）

```yaml
name: Build and Deploy

on:
  push:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: 1.21
      - run: go test ./...
      - run: docker build -t game-server:${{ github.sha }} .
```

## 📋 デプロイメントチェックリスト

- [ ] 全テストが通っているか確認
- [ ] マイグレーション実行済みか確認
- [ ] 環境変数（DB設定など）が定義されているか
- [ ] ログレベルは本番適切か
- [ ] セキュリティキー・トークンが正しいか

## 🔐 本番環境設定

### 環境変数の例
```bash
DB_HOST=production-db.example.com
DB_PORT=5432
DB_NAME=game_center
DB_USER=${SECRET_DB_USER}
DB_PASSWORD=${SECRET_DB_PASSWORD}
REDIS_URL=redis://production-redis:6379
LOG_LEVEL=info
```

### ヘルスチェック

```bash
curl http://localhost:8080/health
```

## 📖 関連ドキュメント

- [アーキテクチャ](ARCHITECTURE.md)
- [開発ルール](DEVELOPMENT_RULES.md)
