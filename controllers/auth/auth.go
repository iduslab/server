package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iduslab/backend/models/req"
	"github.com/iduslab/backend/utils"
	"github.com/iduslab/backend/utils/res"
)

func Auth(c *gin.Context) {
	query := c.MustGet("query").(*req.Auth)
	r := res.New(c)

	result, err := utils.DiscordOAuth2(query.RedirectUri, query.Code, "")
	if err != nil {
		r.SendError(res.ERR_SERVER, err.Error())
	}

	r.Response(result)
}

func AuthURL(c *gin.Context) {
	r := res.New(c)
	query := c.MustGet("query").(*req.AuthURL)

	url := fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=identify%%20email", utils.Config().Discord.ClientID, query.RedirectUri)

	r.Response(res.R{
		"link": url,
	})
}

func SignInToken(c *gin.Context) {
	r := res.New(c)
	body := c.MustGet("body").(*req.AuthToken)

	if user, err := utils.DiscordUsersMe(body.AccessToken); err == nil {
		fmt.Println(user)
		fmt.Println(err)
		if hasPermission, err := utils.HasPermission(user.ID); err == nil {
			r.Response(res.R{
				"access_token":  body.AccessToken,
				"refresh_token": body.RefreshToken,
				"is_admin":      hasPermission,
			})
		} else {
			r.SendError(res.ERR_SERVER, err.Error())
		}
		return
	}

	result, err := utils.DiscordOAuth2(body.RedirectUri, "", body.RefreshToken)
	if err != nil {
		r.SendError(res.ERR_AUTH, err.Error())
		return
	}
	user, err := utils.DiscordUsersMe(body.AccessToken)
	if err != nil {
		r.SendError(res.ERR_SERVER, err.Error())
		return
	}

	hasPermission, err := utils.HasPermission(user.ID)
	if err != nil {
		r.SendError(res.ERR_SERVER, "")
		return
	}
	r.Response(res.R{
		"access_token":  result.AccessToken,
		"refresh_token": result.RefreshToken,
		"is_admin":      hasPermission,
	})
}
