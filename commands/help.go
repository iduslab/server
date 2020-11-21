package commands

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
	"github.com/iduslab/backend/models"
	"github.com/iduslab/backend/utils/embed"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	e := embed.New(s, m.ChannelID, "도움말")

	jsonFile, err := ioutil.ReadFile("help.json")
	if err != nil {
		e.SendEmbed(embed.ERR_REQUEST, "도움말을 불러오는도중 에러가 발생하였습니다")
		return
	}
	var help []models.Help
	err = json.Unmarshal(jsonFile, &help)
	if err != nil {
		e.SendEmbed(embed.ERR_REQUEST, "도움말을 불러오는도중 에러가 발생하였습니다")
		return
	}

	for _, item := range help {
		e.AddListField(item.Command+" "+item.Usage, item.Description, false)
	}

	e.SendEmbed(embed.WITH_LIST, "봇 도움말을 보여줍니다")
}
