package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/iduslab/backend/models"
)

func DiscordUsersMe(accessToken string) (*models.DiscordUser, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "https://discord.com/api/v6/users/@me", nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result models.DiscordUser
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		return &result, errors.New("request error")
	}

	return &result, err
}

func DiscordOAuth2(redirectUri, code, refreshToken string) (*models.DiscordOAuth2, error) {
	config := Config().Discord

	data := url.Values{}
	data.Set("redirect_uri", redirectUri)
	data.Set("scope", "identify guilds")
	data.Set("client_id", config.ClientID)
	data.Set("client_secret", config.ClientSecret)

	if refreshToken == "" {
		data.Set("grant_type", "authorization_code")
		data.Set("code", code)
	} else {
		data.Set("grant_type", "refresh_token")
		data.Set("refresh_token", refreshToken)
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, "https://discord.com/api/v6/oauth2/token", strings.NewReader(data.Encode())) // URL-encoded payload
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result models.DiscordOAuth2
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return &result, err
	}

	return &result, err
}
