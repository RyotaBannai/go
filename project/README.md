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

### メソッド
- コンストラクタを提供する 命名規則は `New[Struct Name]`
```go
package person
import (
	"fmt"
)
type Person struct{
	Id int
	Name string
}

func (p *Person) Greet() {
	fmt.Printf("Hello! I'm %s", p.Name)
}

func NewUser(id int, name string) *Person {
	u := new(Person)
	u.Id = id
	u.Name = name
	return u
}

func main()  {
 f := (*Person).Greet // 関数としてのメソッド
 f(&Person{Id: 0, Name: "Ryota"})
}
```
- メソッドを関数使用する

### nil interface
- `nil interface` がメソッドを呼び出すとランタイムエラー
- `interface を実装した具体的な値が nil の場合`は nil をレシーバーとして呼び出されるため、ランタイムエラーにはならない
```go
// つまりこのような nil interface はエラー

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
}

func main() {
	var j I
	var t *T
	j = t
	j.M() // nil as value is ok.

	var i I // nil interface
	i.M() // Panic!
}

```
- `the empty interface`: 空の interface は未知の型を扱うコードで使用される（入ってくる型が未知なので、empty interface を定義し、なんでも受け入れるようにする.
）

### interface に対しての考え方
- `interface` はメソッドの集まりである。それらのメソッドは値に対して適用されるわけであり、`interface は値を保持するという概念`が重要である。interface を通して、その値の型をチェックすることができるし、その値の型によって条件分することもできる。（type assertions, type switches）

### goroutine
- 通常、片方準備ができるまで送受信はブロックされるため、明確なロックや条件変数が無くても goroutine の同期が可能である.

### pointer 
- ポインタ型は、型情報とメモリ上のデータがあるアドレスを保持したもの（データ型は、型情報とそのデータを保持したもの）
- string 型は immutable で、文字列のインデックスのポインタを取り出して、それ自体を変更することはできない → 通常の文字列の操作は別のデータを作成している
- string 型は、内部的に「文字列の実態へのポインタ」と「文字列のバイト長」によって構成される. つまり string 型は、その型の仕組みそのものにポインタを内包してるため、関数の引数として string 型を値渡ししたとしても、文字列の実態へのポインタと文字列のバイト長という２つの値がコピーされるだけでことが足りるため、文字列の実態がコピーされることはない
