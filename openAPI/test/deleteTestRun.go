package test

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteTestRunReq struct {
	Action      string `json:"Action"`      // DeleteTestRun
	ProjectName string `json:"ProjectName"` // project name
	RunID       int64  `json:"RunId"`       // 1
}

type DeleteTestRunResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // daba6d8a-0a62-e6fd-f502-a4594888561a
	} `json:"Response"`
}

// DeleteTestRun 删除测试计划
func (t *Test) DeleteTestRun(req *DeleteTestRunReq) (resp DeleteTestRunResp, err error) {
	err = poster.Post(t.c, "DeleteTestRun", req, &resp)
	return
}
