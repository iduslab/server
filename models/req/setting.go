package req

type SettingAddValue struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       string `json:"value"`
}

type SettingUpdateValue struct {
	Value string `json:"value"`
}
