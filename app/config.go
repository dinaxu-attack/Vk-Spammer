package app

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Message     string
	Proxies     []string
	Anticaptcha string
}

var (
	Message     string
	Proxies     []string
	Anticaptcha string
)

func Read() {
	var Configg Config
	file, _ := ioutil.ReadFile("./config.json")
	json.Unmarshal(file, &Configg)

	Anticaptcha = Configg.Anticaptcha
	Proxies = Configg.Proxies
	Message = Configg.Message
}
