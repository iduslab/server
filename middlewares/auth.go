package middlewares

import (
	"strings"

	"github.com/iduslab/backend/models"

	"github.com/iduslab/backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/iduslab/backend/utils/res"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := res.New(c)
		authHeader := c.GetHeader("Authorization")
		spilitHeader := strings.Split(authHeader, "Bearer ")
		if len(spilitHeader) != 2 {
			r.SendError(res.ERR_AUTH, "Error while parsing header")
			return
		}
		token := spilitHeader[1]
		user, err := utils.DiscordUsersMe(token)
		if err != nil {
			r.SendError(res.ERR_SERVER, "Error while loading user info")
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := res.New(c)
		user := c.MustGet("user").(*models.DiscordUser)
		isAdmin, _ := utils.HasPermission(user.ID)
		if !isAdmin {
			r.SendError(res.ERR_PERMISSION, "You are not admin!")
			return
		}
		c.Next()
	}
}
