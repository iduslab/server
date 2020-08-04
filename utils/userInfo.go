package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gangjun06/bot01/structure"
)

func GetUserInfoByUserId(id string) (*structure.DiscordUser, error) {
	req, err := http.NewRequest("GET", "https://discord.com/api/v6/users/"+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bot "+Config().Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := &structure.DiscordUser{}
	json.NewDecoder(resp.Body).Decode(result)

	return result, nil
}
