package wh

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	
	if m.GuildID != "" && !strings.HasPrefix(m.Content, "wh?") {
		return
	}

	content := strings.TrimLeft(strings.TrimSpace(m.Content), "wh?")
	for _, c := range wh.commands {

		for _, alias := range c.Aliases() {
			if strings.HasPrefix(content, alias) {
				err := c.Run(s, m.Message)
				if err != nil {
					log.Println(err)
				}
				return
			}
		}
	}

	return

}


func onReady(s *discordgo.Session, ready *discordgo.Ready) {
	fmt.Println("WH-Bot is running.")
	fmt.Println(ready.User.String())
	fmt.Println("ID: " + ready.User.ID)
	fmt.Println("\t" + strings.Repeat("-", 30) + "\t")
	return
}
