package models

type DiscordOAuth2 struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"identify"`
}

type DiscordUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// {
//   "content": "a",
//   "embeds": [
//     {
//       "title": "공지사항",
//       "description": "설명",
//       "color": 16426522,
//       "fields": [
//         {
//           "name": "a",
//           "value": "asdfasdfasddfa"
//         },
//         {
//           "name": "b",
//           "value": "asdfasdfasdfasf"
//         }
//       ]
//     }
//   ]
// }
type DiscordWebhook struct {
	Content string `json:"content"`
	Embeds  []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Color       int    `json:"color"`
		fields      []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		}
	} `json:"embeds"`
}
