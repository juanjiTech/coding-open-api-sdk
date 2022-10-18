package cd

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Cd struct {
	c *config.Basic
}

// New 新建持续部署相关函数
func New(c *config.Basic) Cd {
	return Cd{c}
}
