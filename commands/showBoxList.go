package commands

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/bot01/db"
	embedUtil "github.com/gangjun06/bot01/utils/embed"
)

// ShowBoxList 현제 생성된 상자들의 목록을 보여줍니다
func ShowBoxList(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	embed := embedUtil.New(s, m.ChannelID, "상자목록")

	data, err := db.GetAllBox()
	if err != nil {
		embed.SendEmbed(embedUtil.ERR_BOT, "데이터를 가져오는도중에 오류가 발생하였습니다")
		return
	}
	for _, item := range data {
		embed.AddListField(strconv.Itoa(item.ID), item.Text, false)
	}

	embed.SendEmbed(embedUtil.WITH_LIST, "등록된 상자 목록을 보여줍니다")
}
