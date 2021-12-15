package main

import (
	"fmt"

	"github.com/Serenity0204/Xuan-Bot/bot"
	"github.com/Serenity0204/Xuan-Bot/config"
)

func main() {
	e := config.ReadConfig()

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	bot.InitMsg()

	<-make(chan struct{})
	return
}
