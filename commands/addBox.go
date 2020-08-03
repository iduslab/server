package commands

import (
	"strings"

	"github.com/gangjun06/bot01/utils"

	"github.com/gangjun06/bot01/db"

	"github.com/bwmarrin/discordgo"
)

// AddBox 쪽지를 담을 상자를 추가합니다
func AddBox(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if m.Author.ID != utils.Config().OwnerID {
		s.ChannelMessageSend(m.ChannelID, "상자를 추가할 권한이 없습니다..")
		return
	}

	if len(args) < 1 {
		s.ChannelMessageSend(m.ChannelID, "상자의 이름을 적어주세요")
		return
	}

	text := strings.Join(args, " ")
	if len(text) > 20 {
		s.ChannelMessageSend(m.ChannelID, "이름을 20글자 아래로 적어주세요.")
		return
	}

	err := db.AddBox(text)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "추가도중 에러가 발생하였습니다")
		return
	}
	s.ChannelMessageSend(m.ChannelID, "성공적으로 추가되었습니다")
}
