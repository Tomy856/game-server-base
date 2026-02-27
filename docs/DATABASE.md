# データベース設計・接続パターン

## 🏗 DB構成

本プロジェクトは以下の DB 構成を採用しています：

- **センターDB**: ユーザー基本情報（ID、UUID）と、ユーザー用個別 DB の接続情報を保持
- **ユーザーDB群**: 複数存在。各ユーザーの個別ゲーム情報を格納

このデザインにより、スケーラビリティを実現します。

## 🔗 DB接続の基本フロー

```
API リクエスト
  │
  ├─① センターDB にアクセス
  │   ユーザー UUID から ユーザーID＋接続情報を取得
  │
  ├─② 接続情報をもとに該当ユーザーDB に接続
  │
  └─③ ユーザーDB からユーザーの個別ゲーム情報を取得
```

## 💾 マイグレーション管理

### ファイル配置
SQL ファイルはプロジェクトルート直下に配置します：

```
migrations/
├── 001_create_center_users.sql
├── 002_create_game_info.sql
└── ...
```

### 設計方針
- マイグレーション（SQL実装）はインフラ層の責務です。
- スキーマはドメイン層のエンティティ定義に従って設計します。

```
domain/entity/*.go    → ビジネス要件の定義
migrations/*.sql      → それに対応する SQL 実装
```

### マイグレーションツール
プロジェクトは `migrate`, `goose` などのマイグレーションツール使用を想定しています。

## 📊 リポジトリパターン

### ドメイン層でインターフェースを定義

```go
// internal/domain/repository/center_repository.go
package repository

type CenterRecord struct {
    UserID int64
    DBHost string
    DBName string
}

type CenterRepository interface {
    FindByUUID(ctx context.Context, uuid string) (*CenterRecord, error)
}
```

### インフラ層で実装

```go
// internal/infrastructure/persistence/center_repo.go
package persistence

type centerRepo struct {
    db *sqlx.DB
}

func (r *centerRepo) FindByUUID(ctx context.Context, uuid string) (*CenterRecord, error) {
    var rec CenterRecord
    err := r.db.GetContext(ctx, &rec, `
        SELECT user_id, db_host, db_name FROM center_users WHERE uuid = $1
    `, uuid)
    if err != nil {
        return nil, fmt.Errorf("FindByUUID: %w", err)
    }
    return &rec, nil
}
```

## 🎯 DB アクセスパターン

よくあるパターンと、対応する関数名：

| パターン | 関数名例 | 層 |
|---------|---------|---|
| ユーザー固有レコード取得 | `GetUserGameInfo()` | Repository |
| 全レコード取得 | `GetAllPlayers()` | Repository |
| 特定条件で取得 | `GetPlayersByRank()` | Repository |

**層別の命名ルール詳細は [DEVELOPMENT_RULES.md](DEVELOPMENT_RULES.md) を参照してください。**

## 🔐 DB セキュリティ原則

- **ドメイン層**はDB技術に依存してはいけません。
- テーブルスキーマの詳細はInfrastructure層に閉じ込めます。
- リポジトリインターフェースを通じてのみドメイン層がデータアクセスできます。

## 📝 参考

- `go.mod` で使用する DB ドライバー（`github.com/lib/pq` など）を確認してください
- SQLクエリは `sqlx` の Context 対応メソッド (`GetContext`, `ExecContext`) を使用してください
