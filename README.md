# todo-backend
todo管理アプリのバックエンドAPI

## Description
Go言語（[echo](https://echo.labstack.com/)フレームワーク）を用いたtodo管理アプリのAPIサーバの実装  
[Clean Architecture](https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047)を参考に「model」「repository」「usecase」「delivery」の4層のアーキテクチャを採用

## todo管理アプリ
以下参照  
[todo管理アプリについて](https://www.notion.so/prmcy/ToDo-14f83b283c4b4bd088ee9f11ebe5be13)

## 機能
* ID、パスワードによる認証機能
* Redisをバックエンドとするセッション管理（[gorilla/sessions](https://github.com/gorilla/sessions)）
* MySQLバックエンドとしたデータベースによるtodo情報の作成、編集、削除
* カーソルを使用したtodo情報のFetch機能
* データベースのmigration機能（[rubenv/sql-migrate](https://github.com/rubenv/sql-migrate)）
* Clean Architectureとmockを使用した各層のテスト
* Go modulesを使用したpackage管理

## Usage

### ディレクトリ構成
```
[project root]
├── client/ (yinhr/todo-frontend）
├── db/
│   ├── Dockerfile (下記Gist参照)
│   └── my.cnf (下記Gist参照)
├── nginx/ (yinhr/todo-nginx）
├── server/ (yinhr/todo-backend *このリポジトリ)
└── docker-compose.yml (下記Gist参照)
```
* [Dockerfile](https://gist.github.com/yinhr/3ff5456bc9859af9de7bde2923b84f94)
* [my.cnf](https://gist.github.com/yinhr/ee5fe7dc88831de8f5994447c89cff93)
* [docker-compose.yml](https://gist.github.com/yinhr/bfe1c20f700df5fca2a44ad18f7f3102)

### コンテナ起動
```
docker-compose up --build
```

### ブラウザで下記にアクセス
```
http://localhost:3000
```

