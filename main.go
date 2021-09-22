package main

import (
	"fmt"

	"testbot.com/TestBot/bot"
	"testbot.com/TestBot/config"
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
