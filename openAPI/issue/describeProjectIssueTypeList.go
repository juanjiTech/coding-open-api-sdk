package issue

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeProjectIssueTypeListReq struct {
	Action      string `json:"Action"`      // DescribeProjectIssueTypeList
	ProjectName string `json:"ProjectName"` // TestDemoProject
}

type DescribeProjectIssueTypeListResp struct {
	IssueTypes []struct {
		Description            string        `json:"Description"`            //
		ID                     int64         `json:"Id"`                     // 1
		IsSystem               bool          `json:"IsSystem"`               // false
		IssueType              string        `json:"IssueType"`              // REQUIREMENT
		Name                   string        `json:"Name"`                   // 用户故事
		SplitTargetIssueTypeID []interface{} `json:"SplitTargetIssueTypeId"` // <nil>
		SplitType              string        `json:"SplitType"`              // ALL_REQUIREMENT
	} `json:"IssueTypes"`
}

// DescribeProjectIssueTypeList 查询项目的事项类型列表
func (i *Issue) DescribeProjectIssueTypeList(req *DescribeProjectIssueTypeListReq) (resp DescribeProjectIssueTypeListResp, err error) {
	err = poster.Post(i.c, "DescribeProjectIssueTypeList", req, &resp)
	return
}
