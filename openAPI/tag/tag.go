package tag

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Tag struct {
	c *config.Basic
}

// New 新建标签相关函数
func New(c *config.Basic) Tag {
	return Tag{c}
}
