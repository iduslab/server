package embed

import "github.com/bwmarrin/discordgo"

type SendEmbedType int

const (
	ERR_REQUEST SendEmbedType = 1 + iota
	ERR_BOT
	WITH_LIST
	ONLY_TEXT
)

type SendEmbed struct {
	S           *discordgo.Session
	ChannelID   string
	Kind        SendEmbedType
	Title       string
	Description string
	List        []*SendEmbedField
	Footer      bool
}

type SendEmbedField struct {
	Title  string
	Text   string
	Inline bool
}
