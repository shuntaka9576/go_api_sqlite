# go_api_sqlite with Litestream on AppRunner

## はじめに
本リポジトリは、GoのREST APIアプリケーションをAppRunner上にホスティングしつつ、LitestreamでSQLite3のデータをS3にレプリケートするサンプルコードです。

GoのREST APIアプリは、[詳解Go言語Webアプリケーション開発](https://www.c-r.com/book/detail/1462)で実装する[budougumi0617/go_todo_app](https://github.com/budougumi0617/go_todo_app)をベースにSQLiteに差し替えています。

## AWS環境作成手順

### 構成
|サービス名|物理名|用途|スタック名
|---|---|---|---|
|ECR|go-api-sqlite|AppRunnerに載せるコンテナイメージレジストリ|dev-go-api-sqlite-ecr|
|S3|(dev)-go-api-sqlite-replica-(アカウントID)|Litestreamのレプリカ先S3|dev-go-api-sqlite-bucket|
|App Runner|dev-go-api-sqlite-app-runner|TODO|dev-go-api-sqlite-app-runner|

### ECR作成

CDKディレクトリへcd
```bash
cd ./_cdk
```

```bash
yarn cdk deploy -c stageName=dev dev-go-api-sqlite-ecr
```

### コンテナイメージ作成とECR登録

```bash
make build

# ログイン(要assume-role)
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin <アカウントID>.dkr.ecr.ap-northeast-1.amazonaws.com
# latestタグうち
docker tag go-api-sqlite:latest <アカウントID>.dkr.ecr.ap-northeast-1.amazonaws.com/go-api-sqlite:latest
# コンテナをECRへpush
docker push <アカウントID>.dkr.ecr.ap-northeast-1.amazonaws.com/go-api-sqlite:latest
```


### AppRunner with S3を作成

```bash
# 依存でS3もデプロイされる
yarn cdk deploy -c stageName=dev dev-go-api-sqlite-app-runner
```


## Goアプリデバッグ

### API一覧

|API PATH|用途|
|---|---|
|POST /tasks|タスク作成|
|GET /tasks|タスク一覧取得|

### 環境変数

|変数名|デフォルト値|説明|
|---|---|---|
|PORT|8080|REST APIの待ち受けポート|
|DB_PATH|todo.db|SQLite3のDBファイルパス|

### コマンド

```bash
go run .
```

```bash
curl -XPOST localhost:8080/tasks -d '{"title": "test1"}'
curl -XGET localhost:8080/tasks
```

