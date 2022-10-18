package branch

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Branch struct {
	c *config.Basic
}

// New 新建分支相关函数
func New(c *config.Basic) Branch {
	return Branch{c}
}
