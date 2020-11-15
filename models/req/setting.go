package req

type SettingAddValue struct {
	Name        string `form:"name"`
	Description string `form:"name"`
	Value       string `form:"value"`
}

type SettingUpdateValue struct {
	Value string `json:"value"`
}
