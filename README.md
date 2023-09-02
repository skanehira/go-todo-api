# Todo API
ent + OpenAPIを使ったサンプルTodo API

## 起動
```sh
go run .
```

## コード生成
`ent/schema`からコードを生成する
スキーマに変更があるたびに実行する必要がある

```
$ make generate
```

## Swagger Editor
APIの仕様を確認用

```sh
# コンテナ立ち上がったらブラウザで http://localhost:8888 を開く
$ make doc
```
