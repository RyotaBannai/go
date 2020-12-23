package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

/*
	・リクエストヘッダー　→　func DumpRequestOut
	・レスポンスヘッダー　→ func DumpResponse
	で内容を確認
*/

var (
	endpoint    = "https://pokeapi.co/api/v2/"
	pokemonList = "pokemon?limit=10&offset=10"
)

func Mapper(r io.Reader) (result map[string]interface{}) {
	byteArray, _ := ioutil.ReadAll(r)
	if unmarshalErr := json.Unmarshal([]byte(byteArray), &result); unmarshalErr != nil {
		fmt.Println("unmarshal Err occurred.")
		return nil
	}
	return
}

// シンプルな方法
func getPokemon() {
	resp, err := http.Get(endpoint + pokemonList)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	prettyPrinter(Mapper(resp.Body))
}

// http.Client を使う方法
func getPokemon2() {
	req, _ := http.NewRequest("GET", endpoint+pokemonList, nil) // req.Header.Set("Authorization", "Bearer access-token")
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	prettyPrinter(Mapper(resp.Body))
}
