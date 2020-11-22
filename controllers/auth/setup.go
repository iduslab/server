package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iduslab/backend/db"
	"github.com/iduslab/backend/models"
	"github.com/iduslab/backend/models/req"
	"github.com/iduslab/backend/utils"
	"github.com/iduslab/backend/utils/res"
)

func Roles(c *gin.Context) {
	r := res.New(c)
	roles, _ := utils.GetRoles()
	r.Response(res.R{
		"data": roles,
	})
}

func IsMember(c *gin.Context) {
	r := res.New(c)
	user := c.MustGet("user").(*models.DiscordUser)
	member, _ := utils.GetMemberInfo(user.ID)
	memberRole, _ := utils.GetMemberRole()
	for _, d := range member.Roles {
		if d == memberRole {
			r.Response(res.R{
				"is_member": true,
			})
			return
		}
	}
	r.Response(res.R{
		"is_member": false,
	})
}

func Setup(c *gin.Context) {
	r := res.New(c)
	body := c.MustGet("body").(*req.AuthSetup)
	user := c.MustGet("user").(*models.DiscordUser)

	introduceChannelID, _ := db.GetSetting("introduceChannelID")
	guildID := utils.Config().Discord.ServerID
	memberRoleID, _ := utils.GetMemberRole()
	utils.Session().ChannelMessageSend(introduceChannelID, fmt.Sprintf("별명: %s\n가입경로: %s\n서버에서 하고싶은것: %s\n하고싶은말: %s\n", body.NickName, body.JoinWith, body.WantToDo, body.Message))
	utils.Session().GuildMemberRoleAdd(guildID, user.ID, memberRoleID)

	roles, _ := utils.GetRoles()

	for _, d := range body.Roles {
		for _, d2 := range *roles {
			if d2.Name == d {
				utils.Session().GuildMemberRoleAdd(guildID, user.ID, d2.RoleID)
				break
			}
		}
	}

	r.Response(res.R{})
}
