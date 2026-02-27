# アーキテクチャ概要

このプロジェクトは、**ドメイン駆動設計（DDD）** と **クリーンアーキテクチャ** に基づいて構築されています。

## 🎯 基本理念

- **ドメイン中心**: ビジネスロジックを中心に据え、他の層から独立させます。
- **層の分離**: 各層は明確な責務を持ち、依存関係は内側から外側へ向かうようにします。
- **技術非依存**: ドメイン層は具体的なフレームワークやインフラに依存せず、テストや移植性が高い設計を目指します。

## 📁 フォルダ構成

```
internal/
├── domain/              # ドメイン層（ビジネスルール）
│   ├── entity/          # エンティティ（ビジネスオブジェクト）
│   ├── value/           # 値オブジェクト（不変なドメイン概念）
│   ├── repository/      # リポジトリのインターフェース（契約）
│   └── service/         # ドメインサービス（エンティティに属さない操作）
├── application/         # アプリケーション層（ユースケース）
│   ├── usecase/         # ユースケース実装（ビジネスフロー）
│   ├── service/         # アプリケーションサービス（レイヤー間の調整）
│   └── dto/             # DTO（DataTransferObject：層間のデータ受け渡し）
├── infrastructure/      # インフラ層（実装詳細）
│   ├── persistence/     # DBリポジトリ実装
│   ├── cache/           # キャッシュ実装
│   └── external/        # 外部APIクライアント等
└── presentation/        # プレゼンテーション層（APIなどの入出力）
    ├── handler/         # HTTPハンドラー（コントローラー）
    ├── router/          # ルーティング設定
    └── middleware/      # ミドルウェア

pkg/                     # 外部共有パッケージやユーティリティ
```

## 📚 各層の責務

### Domain層（最も変わりにくい）
ビジネスルールを実装します。フレームワークやDBに依存しません。

```go
// domain/entity/user.go
type User struct {
    ID   int64
    UUID string
}
```

### Application層
ユースケースを実装します。ドメイン層の規則に従い、外部への呼び出しを調整します。

```go
// application/usecase/get_game_info.go
type GetGameInfoUseCase struct {
    centerRepo repository.CenterRepository
    userRepo   repository.UserRepository
}
```

### Infrastructure層
DB、キャッシュ、外部API などの具体的な技術を実装します。

```go
// infrastructure/persistence/user_repo.go
type UserRepository struct {
    db *sqlx.DB
}
```

### Presentation層（最も頻繁に変わる）
HTTPハンドラー、ルーティング、ミドルウェアなど、APIの入出力を担当します。

```go
// presentation/handler/game_handler.go
func (h *GameHandler) GetMyGameInfo(c *gin.Context) {
    // リクエスト取得 → ユースケース呼び出し → レスポンス返却
}
```

## 🔄 依存関係の方向

```
presentation/handler
         ↓
application/usecase
         ↓
domain/{entity, repository, service}
         ↓
infrastructure/{persistence, cache, external}
```

**重要**: 下層が上層に依存してはいけません。  
例：`domain/`が`infrastructure/`をインポートしてはいけない。

## 📖 詳細ドキュメント

- [データベース設計・接続パターン](DATABASE.md)
- [開発ルール・命名規則](DEVELOPMENT_RULES.md)
- [Protocol Buffer 運用](PROTOBUF.md)
