package test

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateCaseResultReq struct {
	Action           string   `json:"Action"`           // CreateCaseResult
	CaseID           int64    `json:"CaseId"`           // 110
	CustomStepStatus []string `json:"CustomStepStatus"` // PASSED
	ProjectName      string   `json:"ProjectName"`      // project-demo
	RunID            int64    `json:"RunId"`            // 110
	Status           string   `json:"Status"`           // PASSED
}

type CreateCaseResultResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 7a671ad6-42f6-1d9d-5c0c-f0723933e9a1
	} `json:"Response"`
}

// CreateCaseResult 测试用例添加测试结果
func (t *Test) CreateCaseResult(req *CreateCaseResultReq) (resp CreateCaseResultResp, err error) {
	err = poster.Post(t.c, "CreateCaseResult", req, &resp)
	return
}
