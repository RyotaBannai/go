### Reference

- http://www.tohoho-web.com/ex/golang.html#hello-world
- [Tour for Go](https://tour.golang.org/moretypes/15)
- [Gorourtine でシェルっぽい書き方をしてみた](https://qiita.com/fufu44/items/e768a5ac0187b4986783)
- [Go でオブジェクト思考](https://qiita.com/kitoko552/items/a6698c68379a8cd8b999#embed%E5%9F%8B%E3%82%81%E8%BE%BC%E3%81%BF)
- [バッチプログラムの Go 言語移行で気をつけるべきこと](https://www.xdata.jp/blogs/bigdata/go_lang.html)

### Memo

- 構造体やそのフィールド、関数、メソッドのスコープは、名前の先頭が大文字か小文字かで決まる
  - 大文字なら public、小文字なら同じパッケージ内に閉じたスコープ

### & and *
`The & Operator`: `&` goes in front of a variable when you want to get that `variable's memory address`.
`The * Operator`: `*` goes in front of a variable that holds a memory address and `resolves it` (it is therefore the counterpart to the & operator).
  - It goes and gets the thing that the pointer was pointing at, e.g. *myString
- `メソッドが変数レシーバである場合、呼び出し時に、変数、または、ポインタのいずれかのレシーバとして取ることができる`
```go
// メソッドが変数レシーバ である場合 (*Vertex じゃない.)
func (v Vertex) Abs() float64 {
  return math.Sqrt(...)
}
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
// この場合、 p.Abs() は (*p).Abs() として解釈される
```

### pointer receiver を使う二つの理由
1. メソッドがレシーバが指す先の変数を変更するため
2. メソッド呼び出しごとに変数のコピーを避けるため（レシーバが大きな構造体である場合に特に効果的である）
- 一般的に値レシーバとポインタレシーバのどちらかで全てのメソッドを与え、`混在させるべきではない.`

