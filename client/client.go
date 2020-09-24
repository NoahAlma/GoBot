package client

import (
	"../config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var prefix string

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	prefix = config.BotPrefix

	if strings.ToLower(m.Content) == prefix + "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}