package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/iduslab/backend/middlewares"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"

	"github.com/iduslab/backend/commands"
	"github.com/iduslab/backend/db"
	"github.com/iduslab/backend/routes"
	"github.com/iduslab/backend/utils"
)

var prefix string
var dg *discordgo.Session

func main() {
	rand.Seed(time.Now().UnixNano())
	utils.LoadConfig()
	InitDB()
	InitBot()
	defer dg.Close()

	InitServer()
}

func InitDB() {
	db.Init(utils.Config().DB.Mongodb)
}

func InitBot() {
	var err error
	config := utils.Config().Discord

	prefix = config.Prefix
	dg, err = discordgo.New("Bot " + config.Token)
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

	utils.SetBotSession(dg)
}

func InitServer() {
	config := utils.Config().Server
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(middlewares.Cors())
	// r.Use(cors.Default())
	version1 := r.Group("/v1")
	routes.InitRoutes(version1)
	r.Run(":" + strconv.Itoa(config.Port))
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
	case "워터마크":
		commands.WaterMark(s, m, args)
	case "도움":
		commands.Help(s, m, args)
	}
}
