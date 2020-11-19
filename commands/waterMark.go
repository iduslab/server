package commands

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/h2non/bimg"
	"github.com/iduslab/backend/utils"

	"github.com/bwmarrin/discordgo"
	"github.com/iduslab/backend/utils/embed"
)

func WaterMark(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	e := embed.New(s, m.ChannelID, "워터마크")
	if len(m.Attachments) < 1 {
		e.SendEmbed(embed.ERR_REQUEST, "사진을 첨부하여 주세요")
		return
	}

	if len(args) != 2 {
		e.SendEmbed(embed.ERR_REQUEST, "워터마크 <글자> <색상 (예: #ffffff)>과 같이 입력해주세요.")
		return
	}

	if !strings.HasPrefix(args[1], "#") {
		e.SendEmbed(embed.ERR_REQUEST, "워터마크 <글자> <색상 (예: #ffffff)>과 같이 입력해주세요.")
		return
	}

	colorR, colorG, colorB, err := utils.Hex2RGB(strings.TrimLeft(args[1], "#"))
	if err != nil {
		e.SendEmbed(embed.ERR_REQUEST, "색상이 올바르지 않습니다")
		return
	}

	fileNameSpilit := strings.Split(m.Attachments[0].Filename, ".")
	fileNameSpilitLength := len(fileNameSpilit)
	if fileNameSpilitLength < 2 {
		e.SendEmbed(embed.ERR_REQUEST, "확장자가 올바르지 않습니다.")
		return
	}
	fileExtension := fileNameSpilit[fileNameSpilitLength-1]

	if fileExtension == "png" || fileExtension == "svg" {
		e.SendEmbed(embed.ERR_REQUEST, "해당 포멧의 이미지는 지원하지 않습니다")
		return
	}

	filePath := fmt.Sprintf("temp/%s.%s", m.Attachments[0].ID, fileExtension)

	if err := utils.DownloadImageViaURL(m.Attachments[0].URL, filePath); err != nil {
		fmt.Println(err.Error())
		e.SendEmbed(embed.ERR_BOT, "사진처리중 문제가 발생하였습니다.")
		return
	}

	buffer, err := bimg.Read(filePath)
	if err != nil {
		fmt.Println(err.Error())
		e.SendEmbed(embed.ERR_BOT, "사진처리중 문제가 발생하였습니다.")
		return
	}

	watermark := bimg.Watermark{
		Text:       args[0],
		Opacity:    0.25,
		Width:      200,
		DPI:        100,
		Margin:     150,
		Font:       "sans bold 12",
		Background: bimg.Color{colorR, colorG, colorB},
	}

	newImage, err := bimg.NewImage(buffer).Watermark(watermark)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		e.SendEmbed(embed.ERR_BOT, "사진처리중 문제가 발생하였습니다.")
		return
	}

	r := bytes.NewReader(newImage)

	s.ChannelFileSend(m.ChannelID, m.Attachments[0].ID+"."+fileExtension, r)

}
