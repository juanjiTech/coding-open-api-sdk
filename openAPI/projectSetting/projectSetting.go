package projectSetting

import "github.com/juanjiTech/coding-open-api-sdk/config"

type ProjectSetting struct {
	c *config.Basic
}

// New 新建项目设置相关函数
func New(c *config.Basic) ProjectSetting {
	return ProjectSetting{c}
}
