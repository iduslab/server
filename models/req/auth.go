package req

type Auth struct {
	Code        string `form:"code"`
	RedirectUri string `form:"redirect_uri"`
}

type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	RedirectUri  string `form:"redirect_uri"`
}

type AuthURL struct {
	RedirectUri string `form:"redirect_uri"`
}

type AuthSetup struct {
	NickName string   `json:"nickname"`
	JoinWith string   `json:"join_with"`
	WantToDo string   `json:"want_to_do"`
	Message  string   `json:"message"`
	Roles    []string `json:"roles"`
}
