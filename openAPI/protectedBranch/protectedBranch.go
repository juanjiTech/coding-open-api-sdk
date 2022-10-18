package protectedBranch

import "github.com/juanjiTech/coding-open-api-sdk/config"

type ProtectedBranch struct {
	c *config.Basic
}

// New 新建保护分支相关函数
func New(c *config.Basic) ProtectedBranch {
	return ProtectedBranch{c}
}
