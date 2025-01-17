package types

type OfficialAccount struct {
	AppID     string `json:"appId"`     // appid
	AppSecret string `json:"appSecret"` // appsecret
}

type OfficialTemplateFiled struct {
	Keyword string `json:"keyword"`
	Name    string `json:"name"`
	Value   string `json:"value"`
	Color   string `json:"color"`
}

type OfficialTemplate struct {
	TemplateID  string                   `json:"templateId"`
	Title       string                   `json:"title"`
	Fields      []*OfficialTemplateFiled `json:"fields"`
	MiniProgram struct {
		AppID    string `json:"appId"`
		PagePath string `json:"pagePath"`
	} `json:"miniProgram,omitempty"`
	URL string `json:"url"`
}
