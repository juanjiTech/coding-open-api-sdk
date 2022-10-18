package test

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteTestCaseSectionReq struct {
	Action      string `json:"Action"`             // DeleteTestCaseSection
	ProjectName int64  `json:"ProjectName,string"` // 2
	SectionID   int64  `json:"SectionId"`          // 1
}

type DeleteTestCaseSectionResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 7a671ad6-42f6-1d9d-5c0c-f0723933e9a1
	} `json:"Response"`
}

// DeleteTestCaseSection 删除测试用例分组
func (t *Test) DeleteTestCaseSection(req *DeleteTestCaseSectionReq) (resp DeleteTestCaseSectionResp, err error) {
	err = poster.Post(t.c, "DeleteTestCaseSection", req, &resp)
	return
}
