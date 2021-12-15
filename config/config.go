package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Token string
var BotPrefix string
var config *configStruct

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Println(e.Error())
		return e
	}

	fmt.Println(string(file))

	e = json.Unmarshal(file, &config)

	if e != nil {
		fmt.Println(e.Error())
		return e
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	return nil
}
