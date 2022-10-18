package issue

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteIssueReq struct {
	Action      string `json:"Action"`      // DeleteIssue
	IssueCode   int64  `json:"IssueCode"`   // 3
	ProjectName string `json:"ProjectName"` // demo-project
}

type DeleteIssueResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 135f80f0-0577-6421-293e-ec231ff3b337
	} `json:"Response"`
}

// DeleteIssue 删除事项
func (i *Issue) DeleteIssue(req *DeleteIssueReq) (resp DeleteIssueResp, err error) {
	err = poster.Post(i.c, "DeleteIssue", req, &resp)
	return
}
