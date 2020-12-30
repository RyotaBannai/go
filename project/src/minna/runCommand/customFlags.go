package main

import (
	"flag"
	"fmt"
	"os"
)

/*
	カスタムの flag を作りたい
		・オプションのエラー処理を独自に実装
		・エラーの出力先を stdout 以外にしたい
*/
/*
	// https://github.com/golang/go/blob/846dce9d05f19a1f53465e62a304dea21b99f910/src/flag/flag.go
	・details:

	type ErrorHandling int

	const (
		ContinueOnError ErrorHandling = iota // Return a descriptive error.
		ExitOnError                          // Call os.Exit(2) or for -h/-help Exit(0).
		PanicOnError                         // Call panic with a descriptive error.
	)

	// IntVar は var を int でキャストしたもの
	func IntVar(p *int, name string, value int, usage string){
		CommandLine.Var(newIntValue(value, p), name, usage)
	}

	type intValue int

	// just assigning a value..
	func newIntValue(val int, p *int) *intValue {
		*p = val
		return (*intValue)(p)
	}

	var CommandLine = NewFlagSet(os.Args[0], ExitOnError)

	func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
		f := &FlagSet{
			name:          name,
			errorHandling: errorHandling,
		}
		f.Usage = f.defaultUsage
		return f
	}

	// CommandLine.Var を呼び出すたびに、formal にオプションが追加されていく..
	func (f *FlagSet) Var(value Value, name string, usage string) {
		// Remember the default value as a string; it won't change.
		flag := &Flag{name, usage, value, value.String()} // ここで flag を新規作成
		_, alreadythere := f.formal[name]
		if alreadythere {
			var msg string
			if f.name == "" {
				msg = fmt.Sprintf("flag redefined: %s", name)
			} else {
				msg = fmt.Sprintf("%s flag redefined: %s", f.name, name)
			}
			fmt.Fprintln(f.Output(), msg)
			panic(msg) // Happens only if flags are declared with identical names
		}
		if f.formal == nil {
			f.formal = make(map[string]*Flag)
		}
		f.formal[name] = flag
	}

	type FlagSet struct {
		Usage func()

		name          string
		parsed        bool
		actual        map[string]*Flag
		formal        map[string]*Flag
		args          []string // arguments after flags
		errorHandling ErrorHandling
		output        io.Writer // nil means stderr; use Output() accessor
	}

	type Flag struct {
		Name     string // name as it appears on command line
		Usage    string // help message
		Value    Value  // value as set
		DefValue string // default value (as text); for usage message
	}

*/

func handlerErr(err error) {
	fmt.Println(err)
}

func main() {
	flags := flag.NewFlagSet("awesomeCmd", flag.ContinueOnError)
	if err := flags.Parse(os.Args[1:]); err != nil {
		// 独自のエラー処理
		handlerErr(err)
	}
	// show usage
	// same as package definition
	flags.Usage = func() { flags.PrintDefaults() }

}
