package main

import (
	"fmt"
	"taxi-bot/database"
	"taxi-bot/tgbot"
)

func main() {
	fmt.Println("start")
	err := database.CheckTable()

	if err == nil {
		tgbot.Start()
	}

	fmt.Println("stop")
}
