package issue

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateIssueBlockReq struct {
	BlockIssueCode int64  `json:"BlockIssueCode"` // 4
	IssueCode      int64  `json:"IssueCode"`      // 1
	ProjectName    string `json:"ProjectName"`    // project-demo
}

type CreateIssueBlockResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// CreateIssueBlock 添加前置事项
func (i *Issue) CreateIssueBlock(req *CreateIssueBlockReq) (resp CreateIssueBlockResp, err error) {
	err = poster.Post(i.c, "CreateIssueBlock", req, &resp)
	return
}
