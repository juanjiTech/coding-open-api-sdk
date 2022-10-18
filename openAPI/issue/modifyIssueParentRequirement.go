package issue

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyIssueParentRequirementReq struct {
	IssueCode       int64  `json:"IssueCode"`       // 10
	ParentIssueCode int64  `json:"ParentIssueCode"` // 1
	ProjectName     string `json:"ProjectName"`     // project-demo
}

type ModifyIssueParentRequirementResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ModifyIssueParentRequirement 修改事项父需求
func (i *Issue) ModifyIssueParentRequirement(req *ModifyIssueParentRequirementReq) (resp ModifyIssueParentRequirementResp, err error) {
	err = poster.Post(i.c, "ModifyIssueParentRequirement", req, &resp)
	return
}
