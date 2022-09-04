gorpを試す　　

# sqliteをセットアップする。
- SQLite用のドライバーのインストール
```go
go get github.com/mattn/go-sqlite3
```

SQLite3 導入
```go
brew install sqlite
```

ビルドと実行
```go
go build sqlite.go
./sqlite
```