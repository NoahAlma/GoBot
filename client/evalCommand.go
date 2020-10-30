package wh

import (
	"github.com/bwmarrin/discordgo"
	"github.com/robertkrimen/otto"
	"strings"
)

type evalCommand struct{}

var jsVm = otto.New()

func (evalCommand) Aliases() []string {
	return []string{
		"eval",
	}
}

func (e *evalCommand) Run(session *discordgo.Session, message *discordgo.Message) error {
	content := strings.TrimPrefix(message.Content, "  ")
	if content[0] == ' ' {
		content = content[1:]
	}
	argsV := strings.Split(content, " ")
	args := argsV[1:]
	input := strings.Join(args, " ")
	jsVm.Set("args", args)
	jsVm.Set("session", session)
	jsVm.Set("message", message)
	val, err := jsVm.Run(input)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, err.Error())
	}
	if !val.IsNull() {
		str, err := val.ToString()
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, err.Error())
		}
		session.ChannelMessageSend(message.ChannelID, str)
	}
	return err
}
