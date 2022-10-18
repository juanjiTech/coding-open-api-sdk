package issue

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Issue struct {
	c *config.Basic
}

// New 新建事项相关函数
func New(c *config.Basic) Issue {
	return Issue{c}
}
