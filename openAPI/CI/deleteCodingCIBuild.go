package ci

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteCodingCIBuildReq struct {
	Action  string `json:"Action"`  // DeleteCodingCIBuild
	BuildID int64  `json:"BuildId"` // 55
}

type DeleteCodingCIBuildResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DeleteCodingCIBuild 删除构建
func (c *Ci) DeleteCodingCIBuild(req *DeleteCodingCIBuildReq) (resp DeleteCodingCIBuildResp, err error) {
	err = poster.Post(c.c, "DeleteCodingCIBuild", req, &resp)
	return
}
