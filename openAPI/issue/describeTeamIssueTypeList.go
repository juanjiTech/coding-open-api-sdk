package issue

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeTeamIssueTypeListReq struct {
	Action string `json:"Action"` // DescribeTeamIssueTypeList
}

type DescribeTeamIssueTypeListResp struct {
	IssueTypes []struct {
		Description string `json:"Description"` //
		ID          int64  `json:"Id"`          // 1
		IsSystem    bool   `json:"IsSystem"`    // false
		IssueType   string `json:"IssueType"`   // REQUIREMENT
		Name        string `json:"Name"`        // 用户故事
	} `json:"IssueTypes"`
}

// DescribeTeamIssueTypeList 查询企业所有事项类型列表
func (i *Issue) DescribeTeamIssueTypeList(req *DescribeTeamIssueTypeListReq) (resp DescribeTeamIssueTypeListResp, err error) {
	err = poster.Post(i.c, "DescribeTeamIssueTypeList", req, &resp)
	return
}
