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

- postgreSQL がインストール済でサービスが起動していること
- .env ファイルを各々の環境に応じた設定にすること
  - ← Todo: Docker ファイルを用意する

## .env

```yml
# .env

PORT=8080
POSTGRES_USER=root      # ユーザー名
POSTGRES_PW=root        # パスワード
POSTGRES_DB=golang      # DB名
POSTGRES_PORT=5432      # ポート
POSTGRES_HOST=localhost # ホスト名
SECRET=uu5pveql
GO_ENV=dev
API_DOMAIN=localhost
FE_URL=http://localhost:5173
```

## 始め方

windows の場合

```bash
set GO_ENV=dev
go run migrate/migrate.go # 初回のみ
go run ./main.go
```
