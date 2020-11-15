package models

type DiscordOAuth2 struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"identify"`
}

type DiscordUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
