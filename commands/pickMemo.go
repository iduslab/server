package commands

import (
	"strconv"

	"github.com/gangjun06/iduslab/utils"
	embedUtil "github.com/gangjun06/iduslab/utils/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/iduslab/db"
)

// PickMemo 상자열기 <상자id> [개수 (기본값=2)]>: 쪽지를 상자에서 몇개 꺼냅니다
func PickMemo(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	embed := embedUtil.New(s, m.ChannelID, "상자열기")

	if len(args) < 1 {
		embed.SendEmbed(embedUtil.ERR_REQUEST, "상자 id를 적어주세요")
		return
	}

	boxID, err := strconv.Atoi(args[0])

	if err != nil {
		embed.SendEmbed(embedUtil.ERR_REQUEST, "상자 id는 숫자로 적어주세요.")
		return
	}

	count := 2
	if len(args) > 1 {
		num, err := strconv.Atoi(args[1])
		if err != nil {
			embed.SendEmbed(embedUtil.ERR_REQUEST, "개수는 숫자로 적어주셔야 합니다")
			return
		}
		count = num
	}

	boxData, err := db.GetBox(boxID)

	if err != nil {
		embed.SendEmbed(embedUtil.ERR_REQUEST, "id와 일치하는 상자가 존재하지 않습니다")
		return
	}

	notes, err := db.PickMemo(boxID, count)
	if err != nil {
		embed.SendEmbed(embedUtil.ERR_BOT, "뽑는도중 에러가 발생하였습니다.")
		return
	}

	for i, item := range notes {
		var author string
		if item.Anon {
			author = "익명"
		} else {
			user, err := utils.GetUserInfoByUserId(item.Author)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "불러오는도중 에러가 발생하였습니다")
				return
			}
			author = user.Username + "#" + user.Discriminator
		}
		embed.AddListField(strconv.Itoa(i+1)+". 작성자: "+author, item.Text, false)
	}

	embed.SendEmbed(embedUtil.WITH_LIST, args[0]+"번 상자("+boxData.Text+")에 든 쪽지를 "+strconv.Itoa(len(notes))+"개 가져왔습니다")

}
