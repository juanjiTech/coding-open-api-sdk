package ci

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ClearCodingCIJobCacheReq struct {
	ID int64 `json:"Id"` // 0
}

type ClearCodingCIJobCacheResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ClearCodingCIJobCache 清理构建计划缓存
func (c *Ci) ClearCodingCIJobCache(req *ClearCodingCIJobCacheReq) (resp ClearCodingCIJobCacheResp, err error) {
	err = poster.Post(c.c, "ClearCodingCIJobCache", req, &resp)
	return
}
