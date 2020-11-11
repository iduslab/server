package commands

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/iduslab/db"
	"github.com/gangjun06/iduslab/utils/embed"
)

// ShowBoxList 현제 생성된 상자들의 목록을 보여줍니다
func ShowBoxList(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	e := embed.New(s, m.ChannelID, "상자목록")

	data, err := db.GetAllBox()
	if err != nil {
		e.SendEmbed(embed.ERR_BOT, "데이터를 가져오는도중에 오류가 발생하였습니다")
		return
	}
	for i, item := range *data {
		e.AddListField(strconv.Itoa(i+1)+". "+item.Title, item.Description+" - 생성일: "+item.Timestamp.Format("2006년 01월 02일"), false)
	}

	e.SendEmbed(embed.WITH_LIST, "등록된 상자 목록을 보여줍니다")
}
