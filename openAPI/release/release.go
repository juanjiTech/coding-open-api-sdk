package release

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Release struct {
	c *config.Basic
}

// New 新建版本相关函数
func New(c *config.Basic) Release {
	return Release{c}
}
