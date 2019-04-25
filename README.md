# 構成

```
zoo
|-- main.go
|-- app.go
|-- app_test.go
|-- go.mod
|
|-- animals/
      |-- elephant.go
      |-- monkey.go
      |-- rabbit.go
      |-- animals_test.go
```

## main.go
- エントリポイント mainパッケージ、func main()
- animalsパッケージをimportする
  - インポートパスは、モジュール名/animals
  - 相対パス（./animals）はダメ

## app.go
- mainパッケージ
- 関数を切り出した

## app_test.go
- app.goのテストコード
- 命名は、_test.go
- import "testing"
- テスト関数の名前は、Test関数名(t *testing T)

## go.mod
- `go mod init モジュール名`で作成されたモジュールファイル

## elephant.go, monkey.go, rabbit.go
- animalsパッケージ
- mainパッケージから呼び出される

## animals_test.go
- animalsパッケージ内のテストコード
- 1ファイルづつ分けてもいいし、このようにまとめてもいい

# メモ
- `go run main.go`
  - エラーになる。main.go内で同じmainパッケージapp.go内の関数を呼んでいる
  - runコマンドの引数は、引数のファイルしかrunしない
  - 正しくは、`go run main.go app.go`

- `go build main.go`
  - 上の`go run main.go`同様にエラーになる
  - この場合は、`go build`
  - runとは違い、引数なしですべての.goが対象になる

- `go test ./animals`
  - animalsパッケージのテストを行う
  - animals-test.goが動く
  - 結果のみ出力
  - 詳細を出力したい場合は、-vをつける
  - animalsディレクトリに移動して実行する場合は、'go test`で可
