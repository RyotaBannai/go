package animals

/*
	・*_test.go はテストファイルとして特別扱いされる.
	・"go test [directory name]"
	・"go test -v [directory name]" // test の詳細まで表示
	・test はそのディレクトリのみで行われ、再起的には実行されない
	・再起的に実行した場合は "go test ./..." とする -> カレントディレクトリ以下を再起的にテスト
*/
import "testing"

func TestElephant(t *testing.T) {
	expect := "Grass"
	actual := Elephant()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	} else {
		t.Log("Test was success!")
	}
}
