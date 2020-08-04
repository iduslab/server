package embed

import (
	"github.com/bwmarrin/discordgo"
)

func New(s *discordgo.Session, channelID string, title string) *SendEmbed {
	return &SendEmbed{S: s, ChannelID: channelID, Title: title, Footer: true}
}

func (s *SendEmbed) AddListField(title, text string, inline bool) {
	s.List = append(s.List, &SendEmbedField{Title: title, Text: text, Inline: inline})
}

func (s *SendEmbed) SendEmbed(kind SendEmbedType, description string) {
	embed := &discordgo.MessageEmbed{
		Title:       s.Title,
		Description: description,
	}

	if kind == WITH_LIST {
		var fields []*discordgo.MessageEmbedField
		for _, item := range s.List {
			field := new(discordgo.MessageEmbedField)
			field.Name = item.Title
			field.Value = item.Text
			field.Inline = item.Inline
			fields = append(fields, field)
		}
		embed.Fields = fields
	}

	switch kind {
	case ERR_BOT:
		embed.Color = 0xffd5cd
	case ERR_REQUEST:
		embed.Color = 0xefbbcf
	case WITH_LIST:
		embed.Color = 0xc3aed6
	case ONLY_TEXT:
		embed.Color = 0xc3aed6
	}

	s.S.ChannelMessageSendEmbed(s.ChannelID, embed)
}
