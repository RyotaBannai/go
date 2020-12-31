package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

type Command interface {
	Synopsis() string      // コマンドの説明 50 文字程度で簡潔に
	Help() string          // 詳細なヘルプメッセージ
	Run(args []string) int // コマンドの機能を実装　引数を受け取り終了ステータスを返す
}

type AddCommand struct {
	name string
}

func (c *AddCommand) Synopsis() string {
	// TODO: get options
	/*
		taskOp := flag.NewFlagSet("-n", flag.ContinueOnError)
		taskOp.StringVar(&taskName, "name", "[no name]", "set task name")
		taskDesc := flag.NewFlagSet("-d", flag.ContinueOnError)
		...
		parse()..
	*/
	return "Add todo task to list"
}

func (c *AddCommand) Help() string {
	return "Usage: todo add [option]"
}

func (c *AddCommand) Run(args []string) int {
	// TODO: todo を追加
	fmt.Printf("added task to list [%s]\n", c.name)
	fmt.Println(c)
	return 0
}

func main() {
	c := cli.NewCLI("todo", "0.1.0")
	c.Args = os.Args[1:]

	fmt.Println(c.Args)

	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &AddCommand{name: "high priority"}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)

	// go run . add qq
}
