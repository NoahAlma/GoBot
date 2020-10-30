package wh

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var wh *WHBot

type WHBot struct {
	session *discordgo.Session
	exitc   chan os.Signal
	commands []Command
}


func (n *WHBot) Start(token string) (err error) {

	n.session, err = discordgo.New("Bot " + token)

	if err != nil {
		return errors.New("error creating session: " + err.Error())
	}

	wh = n

	n.session.AddHandler(messageCreateHandler)
	n.session.AddHandler(onReady)

	err = n.session.Open()

	if err != nil {
		return errors.New("error opening connection: " + err.Error())
	}

	defer n.session.Close()

	signal.Notify(n.exitc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-n.exitc

	return nil
}

func NewWH() *WHBot {
	return &WHBot{
		exitc:    make(chan os.Signal, 1),
		commands: commands,
	}
}
