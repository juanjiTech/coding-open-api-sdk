package test

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Test struct {
	c *config.Basic
}

// New 新建测试管理相关函数
func New(c *config.Basic) Test {
	return Test{c}
}
