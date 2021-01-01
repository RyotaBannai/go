package main

import "testing"

/*
	・カバレッジプロファイルをファイルへ出力
	go test -cover -coverprofile=cover.out
	・HTML へ変換
	go tool cover -html=cover.out -o cover.html

	also refer to:
	https://qiita.com/kkohtaka/items/965fe08821cda8c9da8a
*/

type Case struct {
	in, out string
}

var cases = []Case{
	{"today is sunny", "a few words"},
	{"Go go go go go get success!", "many words"},
}

func TestWords(t *testing.T) {
	for i, c := range cases {
		w := Words(c.in)
		if w != c.out {
			t.Errorf("#%d: Words(%s) got %s; want %s", i, c.in, w, c.out)
		}
	}
}
