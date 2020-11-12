package utils

import (
	"github.com/bwmarrin/discordgo"
)

var session_local *discordgo.Session

func SetBotSession(s *discordgo.Session) {
	session_local = s
}

func Session() *discordgo.Session {
	return session_local
}

func GetMemberInfo(user_id string) (*discordgo.Member, error) {
	serverID := Config().Discord.ServerID
	member, err := Session().GuildMember(serverID, user_id)
	return member, err
}
