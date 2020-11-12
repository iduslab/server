package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/iduslab/backend/models/req"
	"github.com/iduslab/backend/utils"
	"github.com/iduslab/backend/utils/res"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	query := c.MustGet("query").(*req.Auth)
	r := res.New(c)

	config := utils.Config().Discord

	data := url.Values{}
	data.Set("code", query.Code)
	data.Set("grant_type", "authorization_code")
	data.Set("scope", "identify guilds")
	data.Set("redirect_uri", query.RedirectUri)
	data.Set("client_id", config.ClientID)
	data.Set("client_secret", config.ClientSecret)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, "https://discord.com/api/v6/oauth2/token", strings.NewReader(data.Encode())) // URL-encoded payload
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		r.SendError(res.ERR_SERVER, err.Error())
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.SendError(res.ERR_SERVER, err.Error())
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		r.SendError(res.ERR_SERVER, err.Error())
		return
	}

	if resp.StatusCode != 200 {
		errmsg := result["error_description"].(string)
		r.SendError(res.ERR_SERVER, errmsg)
		return
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
