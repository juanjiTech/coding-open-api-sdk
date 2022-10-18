package ci

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteCodingCIJobReq struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DeleteCodingCIJob 删除构建计划
func (c *Ci) DeleteCodingCIJob(req *DeleteCodingCIJobReq) error {
	return poster.Post(c.c, "DeleteCodingCIJob", req, nil)
}
