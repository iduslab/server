package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gangjun06/iduslab/commands"
	"github.com/gangjun06/iduslab/db"
	"github.com/gangjun06/iduslab/utils"
)

var prefix string

func main() {

	config, loadConfigErr := utils.LoadConfig()
	if loadConfigErr != nil {
		log.Fatalln("Error While Loading Config File.\nMake sure config.json is located in the project root and written correctly." + loadConfigErr.Error())
		return
	}

	db.Init(config.Mongodb)

	prefix = config.Prefix

	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalln("Error Creating Discord Session, " + err.Error())
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// 봇이면 무시하기
	if m.Author.Bot == true {
		return
	}

	// 프리픽스 안맞으면 무시하기
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	// 채팅 파싱
	list := strings.Split(m.Content, " ")
	command := string([]rune(list[0])[1:])
	var args []string
	if len(list) > 1 {
		args = list[1:]
	}

	switch command {
	case "상자추가":
		commands.AddBox(s, m, args)
	case "상자목록":
		commands.ShowBoxList(s, m, args)
	case "쪽지추가":
		commands.AddMemo(s, m, args)
	case "상자열기":
		commands.PickMemo(s, m, args)
	case "도움":
		commands.Help(s, m, args)
	}
}
