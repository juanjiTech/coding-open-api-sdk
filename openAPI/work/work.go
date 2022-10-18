package work

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Work struct {
	c *config.Basic
}

// New 新建工时相关函数
func New(c *config.Basic) Work {
	return Work{c}
}
