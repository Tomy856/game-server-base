# テスト戦略

本プロジェクトではテストを多層的に実施することで、品質を確保します。

## 📚 テストの種類

### 1. ユニットテスト

**対象：** ドメイン層、アプリケーション層のロジック  
**実装：** 依存性を Mock/Stub で置き換える

```go
// internal/domain/service/game_service_test.go
func TestCalculateScore(t *testing.T) {
    // Arrange
    svc := NewGameService()
    
    // Act
    score := svc.CalculateScore(100, 2)
    
    // Assert
    if score != 200 {
        t.Errorf("expected 200, got %d", score)
    }
}
```

### 2. ユースケーステスト

**対象：** Application層のユースケース  
**実装：** Repository インターフェースをモック

```go
// internal/application/usecase/get_game_info_test.go
type mockCenterRepo struct{}

func (m *mockCenterRepo) FindByUUID(ctx context.Context, uuid string) (*CenterRecord, error) {
    return &CenterRecord{UserID: 123}, nil
}

func TestGetGameInfo(t *testing.T) {
    uc := NewGetGameInfoUseCase(&mockCenterRepo{}, &mockUserRepo{})
    info, _ := uc.Execute(context.TODO(), "test-uuid")
    // assertions...
}
```

### 3. 統合テスト

**対象：** Repository + 実DB、複数層の連携  
**実装：** Docker Compose で test DB を起動

```bash
# テスト前に test DB 起動
docker-compose -f docker-compose.test.yml up -d

# テスト実行
go test ./...

# テスト後に cleanup
docker-compose -f docker-compose.test.yml down
```

## ✅ テスト実装ガイドライン

### テストの命名規則
```go
// ✅ 良い例
TestGetGameInfo_Successful()
TestGetGameInfo_UserNotFound()
TestGetGameInfo_InvalidUUID()

// ❌ 避ける例
Test1()
TestGetGameInfo()  // 結果を含めない
```

### 依存性のモック化
```go
// ✅ 良い例：インターフェース経由
type mockRepo struct{}

func (m *mockRepo) FindByUUID(ctx context.Context, uuid string) (*CenterRecord, error) {
    return &CenterRecord{UserID: 123}, nil
}

// ❌ 避ける例：実装を直接モック
```

## 🎯 テストカバレッジ目標

- **ドメイン層**: 80%以上
- **アプリケーション層**: 70%以上
- **肯定系・否定系**両方を実装

## 🔍 テスト実行

```bash
# 全テスト実行
go test ./...

# カバレッジ確認
go test -cover ./...

# プロファイル付き
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 📖 参考

- [Golang Testing](https://golang.org/doc/effective_go#testing)
- [Table-Driven Tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
