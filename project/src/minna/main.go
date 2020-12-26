package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/structs"
	"github.com/k0kubun/pp"
	"log"
	"minna/lib"
	_ "startingGo/buildins/useful"
	"sync"
)

type KeyValue struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewKeyValue() *KeyValue {
	return &KeyValue{store: make(map[string]string)}
}

func (kv *KeyValue) Set(key, val string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.store[key] = val
}

func (kv *KeyValue) Get(key string) (string, bool) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	val, ok := kv.store[key]
	return val, ok
}

func readMap() {
	kv := NewKeyValue()
	kv.Set("Key", "Value")
	val, ok := kv.Get("Key")
	if ok {
		fmt.Println(val)
	}
}

func main() {
	cfg, err := lib.LoadConfigEnv()
	if err != nil {
		log.Fatal("error occurred..")
	}

	// 綺麗に表示させる
	b, _ := json.MarshalIndent(&cfg, "", "	")
	fmt.Println(string(b))

	// convert Struct to Map
	mapped := structs.Map(&cfg)
	fmt.Println(mapped)
	fmt.Printf("%T\n", mapped)

	// use spew
	spew.Dump(mapped)

	// pretty print
	pp.Println(mapped)

}
