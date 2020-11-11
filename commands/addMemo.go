package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/iduslab/db"
	"github.com/gangjun06/iduslab/utils/embed"
)

func AddMemo(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	e := embed.New(s, m.ChannelID, "쪽지추가")

	if len(args) < 2 {
		e.SendEmbed(embed.ERR_REQUEST, "쪽지추가 <상자이름> <쪽지에 넣을내용>\n다음과 같이 입력해주세요")
		return
	}

	box, err := db.GetBox(args[0])
	if err != nil {
		e.SendEmbed(embed.ERR_REQUEST, "이름과 일치하는 상자가 존재하지 않습니다")
		return
	}

	text := strings.Join(args[1:], " ")

	if len(text) > 30 {
		e.SendEmbed(embed.ERR_REQUEST, "쪽지 내용은 10글자 이하로 적어주세요")
		return
	}

	if err := db.AddMemo(box.ID, m.Author.ID, text); err != nil {
		e.SendEmbed(embed.ERR_BOT, "추가도중 에러가 발생하였습니다")
		return
	}

	s.ChannelMessageDelete(m.ChannelID, m.ID)
	e.SendEmbed(embed.ONLY_TEXT, m.Author.Username+"님의 쪽지가 성공적으로 추가되었습니다.")
}
