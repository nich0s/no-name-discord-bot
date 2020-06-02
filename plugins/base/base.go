package base

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Handle - the word Handle
func Handle(s *discordgo.Session, e *discordgo.Event) {
	if e.Type == "MESSAGE_CREATE" {
		event := e.Struct.(*discordgo.MessageCreate)
		fmt.Println("it works")
		switch event.Content {
		case ".echo":
			echoMessage(event, s)
		}
	}
}

func echoMessage(m *discordgo.MessageCreate, s *discordgo.Session) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	s.ChannelMessageSend(m.ChannelID, m.Content)
}
