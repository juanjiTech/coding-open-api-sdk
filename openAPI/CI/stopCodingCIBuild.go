package ci

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type StopCodingCIBuildReq struct {
	Action  string `json:"Action"`  // StopCodingCIBuild
	BuildID int64  `json:"BuildId"` // 1
}

type StopCodingCIBuildResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// StopCodingCIBuild 停止构建
func (c *Ci) StopCodingCIBuild(req *StopCodingCIBuildReq) (resp StopCodingCIBuildResp, err error) {
	err = poster.Post(c.c, "StopCodingCIBuild", req, &resp)
	return
}
