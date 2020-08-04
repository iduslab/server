package commands

import (
	"strings"

	"github.com/gangjun06/bot01/utils"
	embedUtil "github.com/gangjun06/bot01/utils/embed"

	"github.com/gangjun06/bot01/db"

	"github.com/bwmarrin/discordgo"
)

// AddBox 쪽지를 담을 상자를 추가합니다
func AddBox(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	embed := embedUtil.New(s, m.ChannelID, "상자추가")

	if m.Author.ID != utils.Config().OwnerID {
		embed.SendEmbed(embedUtil.ERR_REQUEST, "상자를 추가할 권한이 없습니다..")
		return
	}

	if len(args) < 1 {
		embed.SendEmbed(embedUtil.ERR_REQUEST, "상자의 이름을 적어주세요")
		return
	}

	text := strings.Join(args, " ")
	if len(text) > 25 {
		embed.SendEmbed(embedUtil.ERR_REQUEST, "이름을 20글자 아래로 적어주세요.")
		return
	}

	err := db.AddBox(text)
	if err != nil {
		embed.SendEmbed(embedUtil.ERR_BOT, "추가도중 에러가 발생하였습니다")
		return
	}
	embed.SendEmbed(embedUtil.ONLY_TEXT, "성공적으로 추가되었습니다")
}
