# Protocol Buffer 運用

本プロジェクトは **Protocol Buffer** を使用して UE/Unity クライアントと通信します。

## 📦 ファイル構成

```
proto/                              # .proto定義ファイル
├── player.proto
├── game.proto
└── ...

internal/presentation/
├── pb/                             # 生成されたProtoBufコード
│   ├── player.pb.go
│   ├── game.pb.go
│   └── ...
└── marshaler/                      # 変換ロジック
    └── proto_converter.go
```

## 🔄 データフロー

```
HTTP Handler
    ↓
ドメインモデル（Entity）
    ↓
marsha ler（変換）
    ↓
Protobuf（バイナリ化）
    ↓
HTTP Response
```

## 🛠 ワークフロー

### 1. Proto定義の作成

```protobuf
// proto/player.proto
syntax = "proto3";
package game;

message Player {
    int64 id = 1;
    string name = 2;
    int32 level = 3;
}
```

### 2. Goコード生成

```bash
protoc --go_out=internal/presentation/pb proto/player.proto
```

### 3. Marshaler実装

```go
// internal/presentation/marshaler/proto_converter.go
package marshaler

import (
    "your/module/internal/domain/entity"
    "your/module/internal/presentation/pb"
)

func PlayerToProto(e *entity.Player) *pb.Player {
    return &pb.Player{
        Id:    e.ID,
        Name:  e.Name,
        Level: int32(e.Level),
    }
}
```

### 4. ハンドラーで使用

```go
// internal/presentation/handler/player_handler.go
func (h *PlayerHandler) GetPlayer(c *gin.Context) {
    player, _ := h.usecase.GetMyPlayer(ctx)
    protoPlayer := marshaler.PlayerToProto(player)
    c.ProtoBuf(200, protoPlayer)
}
```

## 📌 重要なポイント

- **Protobuf は Presentation層に限定** → ハンドラーでのみ使用
- **ドメイン層は proto を知らない** → エンティティと独立
- **Marshaler で変換を集約** → 一箇所で管理

## 🔐 セキュリティ

- 機密情報（トークンなど）は Protobuf レスポンスに含めない
- 検証は Handler で実施してから Marshal する

## 📖 参考

- [Google Protocol Buffers](https://developers.google.com/protocol-buffers)
- [protoc-gen-go](https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go)
