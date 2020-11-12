package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/iduslab/backend/db"
	"github.com/iduslab/backend/utils"
	"github.com/iduslab/backend/utils/embed"
)

// AddBox 쪽지를 담을 상자를 추가합니다
func AddBox(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	e := embed.New(s, m.ChannelID, "상자추가")

	hasPermission := false
	for _, d := range m.Member.Roles {
		if d == utils.Config().Discord.PermissionRole {
			hasPermission = true
		}
	}
	if !hasPermission {
		e.SendEmbed(embed.ERR_REQUEST, "상자를 추가할 권한이 없습니다")
		return
	}

	if len(args) < 2 {
		e.SendEmbed(embed.ERR_REQUEST, "상자의 이름/설명을 적어주세요")
		return
	}

	text := args[0]
	if len(text) > 25 {
		e.SendEmbed(embed.ERR_REQUEST, "이름을 20글자 아래로 적어주세요.")
		return
	}

	description := strings.Join(args[1:], " ")

	err := db.AddBox(text, description)
	if err != nil {
		e.SendEmbed(embed.ERR_BOT, "추가도중 에러가 발생하였습니다")
		return
	}
	e.SendEmbed(embed.ONLY_TEXT, "성공적으로 추가되었습니다")
}
