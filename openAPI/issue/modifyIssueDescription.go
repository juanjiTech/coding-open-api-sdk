package issue

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyIssueDescriptionReq struct {
	Action      string `json:"Action"`      // ModifyIssueDescription
	Description string `json:"Description"` // **description**
	IssueCode   int64  `json:"IssueCode"`   // 17
	ProjectName string `json:"ProjectName"` // shashajingdianxiangmu2
}

type ModifyIssueDescriptionResp struct {
	Response struct {
		IssueDescription struct {
			Description string `json:"Description"` // <p><strong>this description</strong></p>
		} `json:"IssueDescription"`
		RequestID string `json:"RequestId"` // 15c9d55a-79b3-4a72-1f31-f101ae9c0ddf
	} `json:"Response"`
}

// ModifyIssueDescription 修改事项描述
func (i *Issue) ModifyIssueDescription(req *ModifyIssueDescriptionReq) (resp ModifyIssueDescriptionResp, err error) {
	err = poster.Post(i.c, "ModifyIssueDescription", req, &resp)
	return
}
