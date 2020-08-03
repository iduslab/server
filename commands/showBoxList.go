package commands

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/bot01/db"
)

// ShowBoxList 현제 생성된 상자들의 목록을 보여줍니다
func ShowBoxList(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	embed := &discordgo.MessageEmbed{
		Title:       "상자목록",
		Description: "추가된 상자들의 목록을 보여줍니다",
	}

	var fields []*discordgo.MessageEmbedField

	data, err := db.GetAllBox()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "데이터를 가져오는도중에 오류가 발생하였습니다")
		return
	}
	for _, item := range data {
		id := strconv.Itoa(item.ID)
		field := new(discordgo.MessageEmbedField)
		field.Name = id
		field.Inline = true
		field.Value = item.Text
		fields = append(fields, field)
	}

	embed.Fields = fields
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
