package model

import (
	"strings"
)

// Cookie 结构体表示一个Cookie对象
type Cookie struct {
	Name   string
	Value  string
	Path   string
	Domain string
	Secure bool
}

// NewCookie 用于创建一个新的Cookie对象
func NewCookie(cookie string) *Cookie {
	c := &Cookie{
		Name:   "",
		Value:  "",
		Path:   "",
		Domain: "",
		Secure: false,
	}
	if cookie != "" {
		arr := strings.Split(cookie, ";")
		for i, itemStr := range arr {
			itemStr = strings.TrimSpace(itemStr)
			item := strings.SplitN(itemStr, "=", 2)
			if len(item) == 2 {
				key := strings.ToLower(item[0])
				if key == "expires" {
					continue
				} else if key == "path" {
					c.Path = item[1]
				} else if key == "domain" {
					c.Domain = item[1]
				} else if i == 0 {
					c.Name = item[0]
					c.Value = item[1]
				}
			} else if itemStr == "SECURE" {
				c.Secure = true
			}
		}
	}
	return c
}

// ToString 将Cookie对象转换为字符串
func (c *Cookie) ToString() string {
	return c.Name + "=" + c.Value
}

// Parse 解析Cookie字符串并返回Cookie对象
func Parse(ck string) *Cookie {
	c := NewCookie(ck)
	if c.Name == "" {
		return nil
	}
	return c
}
