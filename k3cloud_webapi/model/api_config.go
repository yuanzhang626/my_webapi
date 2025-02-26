package model

var ApiConfig *ApiConfigS

// ApiConfig 结构体定义
type ApiConfigS struct {
	XorCode        string
	ServerURL      string
	Dcid           string
	UserName       string
	AppID          string
	AppSecret      string
	Lcid           int
	OrgNum         int
	ConnectTimeout int
	RequestTimeout int
	Proxy          string
}

func NewApiConfig() *ApiConfigS {
	c := &ApiConfigS{
		ServerURL:      "https://api.kingdee.com/galaxyapi/",
		Dcid:           "",
		UserName:       "",
		AppID:          "",
		AppSecret:      "",
		Lcid:           2052,
		OrgNum:         0,
		ConnectTimeout: 120,
		RequestTimeout: 120,
		Proxy:          "",
	}
	ApiConfig = c

	return c
}
