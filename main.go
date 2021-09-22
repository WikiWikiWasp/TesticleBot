package main

import (
	"fmt"

	"testbot.com/m/bot"
	"testbot.com/m/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
