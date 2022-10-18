package test

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteTestCaseReq struct {
	Action      string `json:"Action"`             // DeleteTestCase
	CaseID      int64  `json:"CaseId"`             // 1
	ProjectName int64  `json:"ProjectName,string"` // 2
}

type DeleteTestCaseResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 7a671ad6-42f6-1d9d-5c0c-f0723933e9a0
	} `json:"Response"`
}

// DeleteTestCase 删除测试用例
func (t *Test) DeleteTestCase(req *DeleteTestCaseReq) (resp DeleteTestCaseResp, err error) {
	err = poster.Post(t.c, "DeleteTestCase", req, &resp)
	return
}
