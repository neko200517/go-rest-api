## 概要

- クリーンアーキテクチャー
- Go 言語
- REST API
  - Echo Framework
  - ユーザー認証
  - CORS 対応
  - CSRF 対策
  - バリデーション
  - ORM (gorm)

## 前提条件

- Docker が実行できる環境

## 始め方

```bash
docker-compose up -d --build
```

初回のみ（マイグレーション）

```bash
docker-compose exec app sh
go run migrate/migrate.go
```

## 動作確認

http://localhost:8080 にアクセスして {"message":"Not Found"} が返ってきたら OK
