package commit

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Commit struct {
	c *config.Basic
}

// New 新建提交相关函数
func New(c *config.Basic) Commit {
	return Commit{c}
}
