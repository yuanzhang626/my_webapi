package model

var Identify *IdentifyS

// IdentifyS 结构体用于存储身份验证相关信息
type IdentifyS struct {
	ServerUrl string
	DCID      string
	LCID      int
	UserName  string
	Pwd       string
	AppId     string
	AppSecret string
	OrgNum    int
}

// NewIdentify 函数用于创建一个新的 IdentifyS 实例
func NewIdentify(serverUrl, dcid, userName, appId, appSecret string, orgNum int, lcid int, pwd string) *IdentifyS {
	if lcid == 0 {
		lcid = 2052 //todo 默认值
	}
	i := &IdentifyS{
		ServerUrl: serverUrl,
		DCID:      dcid,
		LCID:      lcid,
		UserName:  userName,
		Pwd:       pwd,
		AppId:     appId,
		AppSecret: appSecret,
		OrgNum:    orgNum,
	}
	Identify = i
	
	return i
}
