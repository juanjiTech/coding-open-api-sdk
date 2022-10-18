package artifact

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Artifact struct {
	c *config.Basic
}

// New 新建制品仓库相关函数
func New(c *config.Basic) Artifact {
	return Artifact{c}
}
