package tag

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateGitProtectedTagRuleReq struct {
	Action  string `json:"Action"`  // CreateGitProtectedTagRule
	DepotID int64  `json:"DepotId"` // 2
	Rule    string `json:"Rule"`    // 2022*
}

type CreateGitProtectedTagRuleResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 3d9b3771-9b56-7da8-a31d-9e5bb402d659
	} `json:"Response"`
}

// CreateGitProtectedTagRule 创建标签保护规则
func (t *Tag) CreateGitProtectedTagRule(req *CreateGitProtectedTagRuleReq) (resp CreateGitProtectedTagRuleResp, err error) {
	err = poster.Post(t.c, "CreateGitProtectedTagRule", req, &resp)
	return
}
