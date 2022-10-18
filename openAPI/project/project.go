package project

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Project struct {
	c *config.Basic
}

// New 新建项目相关函数
func New(c *config.Basic) Project {
	return Project{c}
}
