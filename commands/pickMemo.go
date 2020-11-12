package commands

import (
	"fmt"
	"strconv"

	"github.com/iduslab/backend/utils/embed"

	"github.com/bwmarrin/discordgo"
	"github.com/iduslab/backend/db"
)

// PickMemo 상자열기 <상자id> [개수 (기본값=2)]>: 쪽지를 상자에서 몇개 꺼냅니다
func PickMemo(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	e := embed.New(s, m.ChannelID, "상자열기")

	if len(args) < 2 {
		e.SendEmbed(embed.ERR_REQUEST, "상자뽑기 <상자이름> <개수> 와 같이 입력해주세요")
		return
	}

	box, err := db.GetBox(args[0])
	if err != nil {
		e.SendEmbed(embed.ERR_REQUEST, "이름과 일치하는 상자가 존재하지 않습니다")
		return
	}

	count, err := strconv.Atoi(args[1])
	if err != nil {
		e.SendEmbed(embed.ERR_REQUEST, "개수는 숫자로 적어주셔야 합니다")
		return
	}

	if count > 5 {
		e.SendEmbed(embed.ERR_REQUEST, "개수는 5개 이하로 적어주세요")
		return
	}

	notes, err := db.PickMemo(box.ID, count)
	if err != nil {
		fmt.Println(err)
		e.SendEmbed(embed.ERR_BOT, "뽑는도중 에러가 발생하였습니다.")
		return
	}

	for i, item := range *notes {
		e.AddListField(strconv.Itoa(i+1)+".", item.Text, false)
	}

	e.SendEmbed(embed.WITH_LIST, args[0]+" 상자에 든 쪽지를 "+strconv.Itoa(len(*notes))+"개 가져왔습니다")

}
