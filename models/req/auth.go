package req

type Auth struct {
	Code        string `form:"code"`
	RedirectUri string `form:"redirect_uri"`
}

type AuthURL struct {
	RedirectUri string `form:"redirect_uri"`
}
