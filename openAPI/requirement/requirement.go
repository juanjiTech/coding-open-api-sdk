package requirement

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Requirement struct {
	c *config.Basic
}

// New 新建需求相关函数
func New(c *config.Basic) Requirement {
	return Requirement{c}
}
