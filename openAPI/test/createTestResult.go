package test

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateTestResultReq struct {
	Action           string   `json:"Action"`           // CreateTestResult
	CustomStepStatus []string `json:"CustomStepStatus"` // PASSED
	ProjectName      string   `json:"ProjectName"`      // project-demo
	Status           string   `json:"Status"`           // PASSED
	TestID           int64    `json:"TestId"`           // 110
}

type CreateTestResultResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 7a671ad6-42f6-1d9d-5c0c-f0723933e9a1
	} `json:"Response"`
}

// CreateTestResult 测试任务添加测试结果
func (t *Test) CreateTestResult(req *CreateTestResultReq) (resp CreateTestResultResp, err error) {
	err = poster.Post(t.c, "CreateTestResult", req, &resp)
	return
}
