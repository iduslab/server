package commands

import (
	"strconv"
	"strings"

	"github.com/gangjun06/bot01/db"
	embedUtil "github.com/gangjun06/bot01/utils/embed"

	"github.com/bwmarrin/discordgo"
)

// AddMemo 쪽지추가 <상자id> <닉네임 공개여부(익명/공개 로 적어주세요)> <쪽지에 넣을내용>: 쪽지를 상자에 넣습니다
func AddMemo(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	embed := embedUtil.New(s, m.ChannelID, "쪽지추가")

	if len(args) < 2 {
		embed.SendEmbed(embedUtil.ERR_REQUEST, "쪽지추가 <상자id> <닉네임 공개여부(익명/공개 로 적어주세요)> <쪽지에 넣을내용>\n다음과 같이 입력해주세요")
		return
	}

	boxID, err := strconv.Atoi(args[0])

	if err != nil {
		embed.SendEmbed(embedUtil.ERR_REQUEST, "상자 id는 숫자로 넣어주세요.")
		return
	}

	var anon bool

	switch args[1] {
	case "익명":
		anon = true
	case "공개":
		anon = false
	default:
		embed.SendEmbed(embedUtil.ERR_REQUEST, "익명 여부는 익명/공개 둘중 하나로 적어주세요.")
		return
	}

	args = args[2:]

	text := strings.Join(args, " ")

	if len(text) > 30 {
		embed.SendEmbed(embedUtil.ERR_REQUEST, "쪽지 내용은 10글자 이하로 적어주세요")
		return
	}

	s.ChannelMessageDelete(m.ChannelID, m.ID)

	if _, err := db.GetBox(boxID); err != nil {
		embed.SendEmbed(embedUtil.ERR_REQUEST, "id와 일치하는 상자가 존재하지 않습니다")
		return
	}

	if err := db.AddMemo(boxID, m.Author.ID, anon, text); err != nil {
		embed.SendEmbed(embedUtil.ERR_BOT, "추가도중 에러가 발생하였습니다")
		return
	}

	embed.SendEmbed(embedUtil.ONLY_TEXT, m.Author.Username+"님의 쪽지가 성공적으로 추가되었습니다.")
}
