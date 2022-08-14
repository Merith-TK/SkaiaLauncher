package main

import (
	"fmt"
	"os"
	"strings"

	toml "github.com/BurntSushi/toml"
)

var (
	conf          config
	configfile    = strings.TrimSuffix(os.Args[0], ".exe") + ".toml"
	dataDir       = "data/"
	defaultConfig = `
[login]
type="invalid"
username=""
password=""
	
[minecraft]
version="1.18.2"
gamedir=".minecraft"
assets="assets"
`

// this is being ignored
)

type config struct {
	Login struct {
		Type     string `toml:"type"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	} `toml:"login"`
	Minecraft struct {
		Version string `toml:"version"`
		Gamedir string `toml:"gamedir"`
		Assets  string `toml:"assets"`
	} `toml:"minecraft"`
}

func setupConfig() error {
	if _, err := os.Stat(configfile); os.IsNotExist(err) {
		fmt.Println("[SkaiaCraft] No config found, creating default config")
		f, err := os.Create(configfile)
		if err != nil {
			return err
		}
		toml.NewEncoder(f).Encode(conf)
		// write config file
		f.Close()
	} else {
		fmt.Println("[SkaiaCraft] Found config, loading config")
		_, err := toml.DecodeFile(configfile, &conf)
		if err != nil {
			return err
		}
	}

	return nil
}
