package commands

import "github.com/bwmarrin/discordgo"

func AddBox(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSend(m.ChannelID, "args[0]: "+args[0]+" args[1]: "+args[1])
}
