package wh

import "github.com/bwmarrin/discordgo"

type Command interface {
	Aliases() []string
	Run(session *discordgo.Session, message *discordgo.Message) error
}

var commands = []Command{
	&evalCommand{},
}
