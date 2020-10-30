package main

import (
	"log"
	"./client"
)

func main() {

	bot := wh.NewWH()
	err := bot.Start("token here")

	if err != nil {
		log.Fatal(err)
	}
}
