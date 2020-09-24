package main

import (
	"./client"
	"./config"
	"fmt"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client.Start()

	<-make(chan struct{})
	return
}