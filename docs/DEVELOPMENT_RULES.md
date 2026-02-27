# 開発ルール・命名規則

## 🏷 関数命名規則

### リポジトリ層（Infrastructure）

DB クエリ操作に基づく命名。`Get` + `[条件]` で統一：

```go
// ✅ 良い例
GetUserGameInfo()        // ユーザー固有データ
GetAllPlayers()          // 全レコード
GetPlayersByRank(rank)   // 条件付き取得
GetUserByID(id)          // IDで検索

// ❌ 避ける例
FetchUserData()          // "Get" を使う
FindGameInfo()           // 層内で統一した動詞を使う
```

### アプリケーション層（Application）

**ユースケース名を優先**。`GetUser*` の命名は避けます：

```go
// ✅ 良い例
GetMyGameInfo()          // 自分のゲーム情報取得（ユースケース）
GetLeaderboard()         // ランキング取得
GetPlayerProfile()       // プロフィール取得

// ❌ 避ける例
GetUserGameInfo()        // Repository層の命名を引きずらない
GetUserLeaderboard()     // "User" をつけすぎ
```

### ハンドラー層（Presentation）

`Handle` + ユースケース名で統一：

```go
// ✅ 良い例
HandleGetMyGameInfo()
HandleGetLeaderboard()
HandleUpdatePlayerScore()
```

## 🎯 重要ルール：層の違いを命名で表現

| 層 | 例 | ポイント |
|---|---|---------|
| **Repository** | `GetUserGameInfo()` | DB固有の操作。user, by, all など |
| **UseCase** | `GetMyGameInfo()` | ビジネスフロー。ユースケース名で表現 |
| **Handler** | `HandleGetMyGameInfo()` | HTTP処理。`Handle` プレフィックス |

## ⚠️ エラーハンドリング

### 必須事項
すべてのエラーは `fmt.Errorf` で文脈を含めてラップしてください：

```go
// ✅ 良い例
func (r *userRepo) GetGameInfo(ctx context.Context, userID int64) (*GameInfo, error) {
    var info GameInfo
    err := r.db.GetContext(ctx, &info, query, userID)
    if err != nil {
        return nil, fmt.Errorf("GetGameInfo: user_id=%d: %w", userID, err)
    }
    return &info, nil
}

// ❌ 避ける例
if err != nil {
    return nil, err  // 文脈なし
}
```

### Context は必須

外部通信やDB操作を伴う関数には、第一引数に `context.Context` を含めてください：

```go
// ✅ 良い例
func (r *userRepo) GetGameInfo(ctx context.Context, userID int64) (*GameInfo, error) {
    return r.db.GetContext(ctx, ...)
}

// ❌ 避ける例
func (r *userRepo) GetGameInfo(userID int64) (*GameInfo, error) {
    return r.db.Get(...)
}
```

## 🔄 トランザクション・整合性

### DBとキャッシュの同時更新

順序を決める場合、説明にその根拠を含めてください：

```go
// 例：DB先行
// 1. DBに書き込む（真実のソース）
// 2. キャッシュを更新
// 失敗時：DBコミット後のキャッシュ失敗は許容（遅延無効化で対応）
```

## ⚙️ パフォーマンス原則

### N+1 問題の禁止

```go
// ❌ 避ける例：ループ内でクエリ
for _, userID := range userIDs {
    info, _ := repo.GetGameInfo(ctx, userID)
    results = append(results, info)
}

// ✅ 良い例：バルク取得
infos, _ := repo.GetGameInfoBatch(ctx, userIDs)
```

## 🔐 並行処理

goroutine で共有リソースにアクセスする場合、`sync.Mutex` など排他制御を必須とし、デッドロックリスクを説明してください：

```go
// ✅ 良い例
type Cache struct {
    mu    sync.RWMutex
    items map[string]interface{}
}

func (c *Cache) Get(key string) interface{} {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.items[key]
}
```

## 📖 関連ドキュメント

- [アーキテクチャ](ARCHITECTURE.md)
- [データベース設計](DATABASE.md)
