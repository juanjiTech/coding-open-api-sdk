package issue

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteIssueBlockReq struct {
	BlockIssueCode int64  `json:"BlockIssueCode"` // 4
	IssueCode      int64  `json:"IssueCode"`      // 1
	ProjectName    string `json:"ProjectName"`    // project-demo
}

type DeleteIssueBlockResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DeleteIssueBlock 删除前置事项
func (i *Issue) DeleteIssueBlock(req *DeleteIssueBlockReq) (resp DeleteIssueBlockResp, err error) {
	err = poster.Post(i.c, "DeleteIssueBlock", req, &resp)
	return
}
