package main

import (
	"fmt"
	"os"

	toml "github.com/BurntSushi/toml"
)

var (
	conf          config
	configfile    = "skaiacraft.toml"
	dataDir       = "skaiacraft/"
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

		// set default config
		conf.Login.Type = "invalid"
		conf.Login.Username = ""
		conf.Login.Password = ""
		conf.Minecraft.Version = "1.18.2"
		conf.Minecraft.Gamedir = ".minecraft"
		conf.Minecraft.Assets = "assets"

		toml.NewEncoder(f).Encode(conf)
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
