package utils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/iduslab/backend/db"
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

func HasPermission(user_id string) (bool, error) {
	user, err := GetMemberInfo(user_id)
	if err != nil {
		return false, err
	}
	value, err := db.GetSetting("permissionID")
	if err != nil {
		return false, err
	}

	hasPermission := false
	for _, d := range user.Roles {
		if d == value {
			hasPermission = true
		}
	}
	return hasPermission, nil
}
