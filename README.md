# go api template

go の API のテンプレートリポジトリです

## ディレクトリ構成

DIは`pkg/adapter/http/handler`でまとめて行います


```
.
├── Makefile
├── README.md
├── cmd
│   └── api // api起動
├── docker
│   ├── dev
│   │   ├── api
│   │   │   └── Dockerfile
│   │   └── db
│   │       └── mysql
│   │           ├── Dockerfile
│   │           └── initdb.d // mysql初期化sql
│   └── prod
│       └── api
│           └── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── pkg
│   ├── adapter
│   │   └── http // http層
│   │       ├── handler // 各APIのDIはここでする
│   │       ├── middleware
│   │       ├── response
│   │       └── router
│   ├── domain
│   │   ├── model // ビジネスロジック
│   │   └── repository // pkg/infra/persistanceの抽象
│   ├── env // 環境変数初期化
│   ├── infra
│   │   ├── dto // dbとやりとりするための構造体
│   │   ├── mysql // mysql接続ロジック
│   │   ├── persistence // domain/repositoryの具体
│   │   └── registry // cmd/api/server.goで呼び出されるrepositoryの初期化
│   ├── interactor // usecaseの具体
│   │   └── mock // usecaseのmockをまとめる
│   └── usecase // usecaseの抽象
├── scripts // 環境構築時に必要なスクリプトなど
```
