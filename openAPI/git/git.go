package git

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Git struct {
	c *config.Basic
}

// New 新建代码托管相关函数
func New(c *config.Basic) Git {
	return Git{c}
}
