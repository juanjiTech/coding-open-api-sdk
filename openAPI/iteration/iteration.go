package iteration

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Iteration struct {
	c *config.Basic
}

// New 新建迭代相关函数
func New(c *config.Basic) Iteration {
	return Iteration{c}
}
