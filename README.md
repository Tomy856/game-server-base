# Game Server Base

このリポジトリは、**ドメイン駆動設計（DDD）** と **クリーンアーキテクチャ** に基づいて構築されたゲームサーバープロジェクトの雛形です。

UE/Unity クライアントとの通信に **Protocol Buffer** を使用し、複数の DB（センターDB + ユーザーDB群）を活用した堅牢なスケーラビリティを実現します。

## 🚀 クイックスタート

### 前提条件
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL
- Redis

### セットアップ

```bash
# リポジトリをクローン
git clone <repository-url>
cd game-server-base

# 依存関係をインストール
go mod download

# DB とキャッシュを起動
docker-compose up -d

# マイグレーション実行
# (マイグレーション方法は docs/DATABASE.md を参照)

# サーバーを起動
go run main.go
```

## 📚 ドキュメント

プロジェクトの詳細は以下の各ドキュメントを参照してください：

| ドキュメント | 説明 |
|-----------|------|
| **[アーキテクチャ](docs/ARCHITECTURE.md)** | 層構成、フォルダ構成、依存関係 |
| **[データベース設計](docs/DATABASE.md)** | DB 接続パターン、マイグレーション管理 |
| **[開発ルール](docs/DEVELOPMENT_RULES.md)** | 命名規則、エラーハンドリング、並行処理 |
| **[Protocol Buffer 運用](docs/PROTOBUF.md)** | .proto ファイル、marshal 実装 |
| **[テスト戦略](docs/TESTING.md)** | ユニットテスト、統合テスト |
| **[デプロイメント](docs/DEPLOYMENT.md)** | Docker、CI/CD、本番環境設定 |

## 🎯 設計の特徴

### ドメイン中心設計
ビジネスロジックを中心に据え、フレームワークやインフラに依存しない設計を実現。

### 多層 DB アーキテクチャ
- **センターDB**: ユーザー基本情報と DB 接続情報
- **ユーザーDB群**: スケーラブルな個別ゲーム情報管理

### クリーンなデータフロー
Protocol Buffer による型安全な UE/Unity クライアント通信。

## 📋 ルール概要

新しく参加するメンバーは必ず以下を確認してください：

- **Repository 層**: DB クエリ中心の命名（`GetUser*`, `GetBy*` など） ✅
- **UseCase 層**: ユースケース本来の名前（`GetMyGameInfo()` など） ❌ `GetUserGameInfo()`
- **すべての Error**: `fmt.Errorf` で文脈をラップ
- **DB 操作**: 第一引数に `context.Context` を含める
- **N+1 禁止**: ループ内のクエリ発行は避け、バルク取得を使用

詳細は [開発ルール](docs/DEVELOPMENT_RULES.md) を参照。

## 📁 フォルダ構成（概要）

```
game-server-base/
├── migrations/               # DB マイグレーション SQL
├── proto/                    # Protocol Buffer 定義
├── internal/                 # 非公開ビジネスロジック
│   ├── domain/              # ドメイン層
│   ├── application/          # アプリケーション層
│   ├── infrastructure/       # インフラ層
│   └── presentation/         # プレゼンテーション層
├── pkg/                      # 公開ライブラリ
├── docker/                   # Docker 関連
├── docs/                     # ドキュメント
├── docker-compose.yml        # 開発環境設定
├── go.mod / go.sum          # 依存管理
└── main.go                  # エントリーポイント
```

詳細は [アーキテクチャ](docs/ARCHITECTURE.md) を参照。

## � 動作環境

### 非サーバープログラマー向け
1. Rancher Desktop をインストール
2. Windows にこの Git リポジトリをクローン
3. `tools/start.bat` でサーバーを起動
4. ブラウザで `http://127.0.0.1:8080` にアクセスして動作確認
5. 終了時は `tools/stop.bat` でサーバーを停止

### サーバープログラマー向け
1. VSCode をインストール
2. Rancher Desktop をインストール
3. 管理者権限で PowerShell を開く
4. `wsl --install -d Ubuntu` で Ubuntu をインストール
5. Ubuntu を起動し、ユーザー名とパスワードを適当に決める
6. Ubuntu 内で適当なフォルダを作成し、この Git リポジトリをクローン
7. Rancher Desktop の Preferences > **WSL** を開く
8. 先ほどインストールした「Ubuntu」がリストに出現するのでチェックを入れる
9. [Apply] ボタンをクリック。これにより Ubuntu 内で `docker` コマンドが利用可能になる
10. Ubuntu にクローンしたリポジトリ上で `code .` を実行
11. VSCode 上で Docker が立ち上がり、仮想環境内でサーバーが起動する
12. ブラウザで `http://127.0.0.1:8080` にアクセスして動作確認

## �🔧 主な技術スタック

- **言語**: Go 1.21+
- **フレームワーク**: Gin
- **DB**: PostgreSQL
- **キャッシュ**: Redis / Memcached
- **通信**: Protocol Buffer
- **テスト**: Go testing（標準）、Table-Driven Tests

## 📞 トラブルシューティング

各ドキュメント内に FAQ とトラブルシューティングが記載されています。

---

**新しく参加するチームメンバーへ**: 上記のドキュメントを読めば、プロジェクト構造と開発ルールが完全に理解できるように設計されています。何か質問や改善提案があればお気軽に PR を作成してください！

