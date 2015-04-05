package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/k0kubun/pp"
)

type Config struct {
	Consumer struct {
		Key    string `toml:"key"`
		Secret string `toml:"secret"`
	} `toml:"consumer_key"`
	Access []struct {
		Token  string `toml:"token"`
		Secret string `toml:"secret"`
	} `toml:"access_token"`
}

var config Config

func loadConfig() {
	_, err := toml.DecodeFile(*conf, &config)
	if err != nil {
		log.Fatal(err)
	}
	initTwitter()
	pp.Print(apis)
}
