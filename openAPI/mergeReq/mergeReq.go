package mergeReq

import "github.com/juanjiTech/coding-open-api-sdk/config"

type MergeReq struct {
	c *config.Basic
}

// New 新建合并请求相关函数
func New(c *config.Basic) MergeReq {
	return MergeReq{c}
}
