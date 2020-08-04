package commands

import (
	"strconv"
	"strings"

	"github.com/gangjun06/bot01/db"

	"github.com/bwmarrin/discordgo"
)

// AddMemo 쪽지추가 <상자id> [익명(0=익명, 1=공개) 기본값=공개] <쪽지에 넣을내용>: 쪽지를 상자에 넣습니다
func AddMemo(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "쪽지추가 <상자id> [익명(0=익명, 1=공개) 기본값=공개] <쪽지에 넣을내용>\n다음과 같이 입력해주세요")
		return
	}

	boxID, err := strconv.Atoi(args[0])

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "상자 id는 숫자로 넣어주세요.")
		return
	}

	var anon bool

	switch args[1] {
	case "1":
		anon = false
	case "0":
		anon = true
	default:
		s.ChannelMessageSend(m.ChannelID, "익명 여부는 0, 1로 적어주세요")
		return
	}

	args = args[2:]

	text := strings.Join(args, " ")

	if len(text) > 30 {
		s.ChannelMessageSend(m.ChannelID, "쪽지 내용은 10글자 이하로 적어주세요")
		return
	}

	if _, err := db.GetBox(boxID); err != nil {
		s.ChannelMessageSend(m.ChannelID, "id와 일치하는 상자가 존재하지 않습니다")
		return
	}

	if err := db.AddMemo(boxID, m.Author.ID, anon, text); err != nil {
		s.ChannelMessageSend(m.ChannelID, "추가도중 에러가 발생하였습니다")
		return
	}
	s.ChannelMessageSend(m.ChannelID, "성공적으로 추가되었습니다")
}
