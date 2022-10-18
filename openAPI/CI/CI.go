package ci

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Ci struct {
	c *config.Basic
}

// New 新建持续集成相关函数
func New(c *config.Basic) Ci {
	return Ci{c}
}
