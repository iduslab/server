package commands

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/bot01/db"
)

// PickMemo 상자열기 <상자id> [개수 (기본값=2)]>: 쪽지를 상자에서 몇개 꺼냅니다
func PickMemo(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 1 {
		s.ChannelMessageSend(m.ChannelID, "상자 id를 적어주세요")
		return
	}

	boxID, err := strconv.Atoi(args[0])

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "상자 id는 숫자로 넣어주세요.")
		return
	}

	count = 2
	if len(args) > 1 {
		num, err := strconv.Atoi(args[1])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "개수는 숫자로 적어주셔야 합니다")
			return
		}
		count = num
	}

	if _, err := db.GetBox(boxID); err != nil {
		s.ChannelMessageSend(m.ChannelID, "id와 일치하는 상자가 존재하지 않습니다")
		return
	}

	notes, err := db.PickMemo(boxID, count)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "뽑는도중 에러가 발생하였습니다.")
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "상자열기",
		Description: args[0] + "번 상자에 든 쪽지를 " + strconv.Itoa(len(notes)) + "개 가져왔습니다",
	}

	var fields []*discordgo.MessageEmbedField

	for i, item := range notes {
		field := new(discordgo.MessageEmbedField)
		field.Name = strconv.Itoa(i + 1)
		field.Inline = true
		field.Value = item.Text
		fields = append(fields, field)
	}

	embed.Fields = fields
	s.ChannelMessageSendEmbed(m.ChannelID, embed)

}
