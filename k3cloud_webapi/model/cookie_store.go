package model

// CookieStore 结构体用于存储会话ID和cookies
type CookieStore struct {
	SID     string
	Cookies map[string]*Cookie
}

// NewCookieStore 是CookieStore的构造函数，用于创建一个新的CookieStore实例
func NewCookieStore(sid string, cookies map[string]*Cookie) *CookieStore {
	if cookies == nil {
		cookies = make(map[string]*Cookie)
	}
	return &CookieStore{
		SID:     sid,
		Cookies: cookies,
	}
}

// SetSID 方法用于设置CookieStore的会话ID
func (cs *CookieStore) SetSID(sid string) {
	if sid != "" {
		cs.SID = sid
	}
}
