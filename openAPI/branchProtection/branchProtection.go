package branchProtection

import "github.com/juanjiTech/coding-open-api-sdk/config"

type BranchProtection struct {
	c *config.Basic
}

// New 新建保护分支规则相关函数
func New(c *config.Basic) BranchProtection {
	return BranchProtection{c}
}
