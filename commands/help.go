package commands

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/iduslab/structure"
	embedUtil "github.com/gangjun06/iduslab/utils/embed"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	embed := embedUtil.New(s, m.ChannelID, "도움말")

	jsonFile, err := ioutil.ReadFile("help.json")
	if err != nil {
		embed.SendEmbed(embedUtil.ERR_REQUEST, "도움말을 불러오는도중 에러가 발생하였습니다")
	}
	var help []structure.Help
	err = json.Unmarshal(jsonFile, &help)
	if err != nil {
		embed.SendEmbed(embedUtil.ERR_REQUEST, "도움말을 불러오는도중 에러가 발생하였습니다")
	}

	for _, item := range help {
		embed.AddListField(item.Command+" "+item.Usage, item.Description, false)
	}

	embed.SendEmbed(embedUtil.WITH_LIST, "봇 도움말을 보여줍니다")
}
