package test

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateTestCaseReq struct {
	Action      string `json:"Action"` // CreateTestCase
	CustomSteps []struct {
		Content  string `json:"Content"`  // expected
		Expected string `json:"Expected"` // steps
	} `json:"CustomSteps"`
	Expected     string `json:"Expected"`           // expected
	Preconds     string `json:"Preconds"`           // preconds
	Priority     int64  `json:"Priority"`           // 1
	ProjectName  int64  `json:"ProjectName,string"` // 2
	SectionID    int64  `json:"SectionId"`          // 1
	Steps        string `json:"Steps"`              // Steps
	TemplateType string `json:"TemplateType"`       // STEPS
	Title        string `json:"Title"`              // helel
}

type CreateTestCaseResp struct {
	Response struct {
		Data struct {
			Case struct {
				CreatedAt   string `json:"CreatedAt"`        // 2020-09-15 16:34:21
				CreatedBy   int64  `json:"CreatedBy,string"` // 1
				CustomSteps []struct {
					Content  string `json:"Content"`  // expected
					Expected string `json:"Expected"` // steps
				} `json:"CustomSteps"`
				Expected     string `json:"Expected"`     // expected
				ID           int64  `json:"Id"`           // 65
				Preconds     string `json:"Preconds"`     // preconds
				Priority     int64  `json:"Priority"`     // 2
				SectionID    int64  `json:"SectionId"`    // 1
				Sort         int64  `json:"Sort"`         // 15
				Steps        string `json:"Steps"`        // setps
				TemplateType string `json:"TemplateType"` // STEPS
				Title        string `json:"Title"`        // helel
				UpdatedAt    string `json:"UpdatedAt"`    // 2020-09-15 16:34:21
			} `json:"Case"`
		} `json:"Data"`
		RequestID string `json:"RequestId"` // 7a671ad6-42f6-1d9d-5c0c-f0723933e9a0
	} `json:"Response"`
}

// CreateTestCase 创建测试用例
func (t *Test) CreateTestCase(req *CreateTestCaseReq) (resp CreateTestCaseResp, err error) {
	err = poster.Post(t.c, "CreateTestCase", req, &resp)
	return
}
