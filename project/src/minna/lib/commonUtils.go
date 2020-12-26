package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

type config struct {
	ApiKey   string `json:"api_key"`
	MaxCount int    `json:"max_count"`
}

func LoadConfig() (*config, error) {
	u, err := user.Current()
	if err != nil {
		fmt.Println("error occurred while reading user.")
		return nil, err
	}

	// go はコンパイルされた後、設定ファイルを変更できて内容を反映できるようにしたいため、package の中に含めない方が良い.
	// windows: %APPDATA%\go-app\.config.json
	// linux: $HOME/.config

	fname := filepath.Join(u.HomeDir, ".config", "go-app", "config.json") // /Users/{username}/.config/go-app/config.json
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("error occurred while opening config file.")
		return nil, err
	}
	defer f.Close()
	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}
