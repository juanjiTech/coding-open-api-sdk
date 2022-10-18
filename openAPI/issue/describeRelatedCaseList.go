package issue

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeRelatedCaseListReq struct {
	Action      string `json:"Action"`      // DescribeRelatedCaseList
	IssueCode   int64  `json:"IssueCode"`   // 1
	ProjectName string `json:"ProjectName"` // my-demo
}

type DescribeRelatedCaseListResp struct {
	Response struct {
		Cases []struct {
			ID       int64  `json:"Id"`       // 1
			Name     string `json:"Name"`     // 注册商城验证码验证通过
			Priority int64  `json:"Priority"` // 1
		} `json:"Cases"`
		RequestID string `json:"RequestId"` // b1410375-b707-6415-85cc-f0c5ec78a5ba
	} `json:"Response"`
}

// DescribeRelatedCaseList 事项关联的测试用例
func (i *Issue) DescribeRelatedCaseList(req *DescribeRelatedCaseListReq) (resp DescribeRelatedCaseListResp, err error) {
	err = poster.Post(i.c, "DescribeRelatedCaseList", req, &resp)
	return
}
