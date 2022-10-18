package wiki

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Wiki struct {
	c *config.Basic
}

// New 新建文档管理相关函数
func New(c *config.Basic) Wiki {
	return Wiki{c}
}
