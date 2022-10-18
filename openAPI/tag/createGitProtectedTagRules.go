package tag

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateGitProtectedTagRulesReq struct {
	Action  string   `json:"Action"`  // CreateGitProtectedTagRules
	DepotID int64    `json:"DepotId"` // 2
	Rule    []string `json:"Rule"`    // 2022.01*
}

type CreateGitProtectedTagRulesResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // e7251f89-b65c-258b-b96e-62e25a2ef74c
	} `json:"Response"`
}

// CreateGitProtectedTagRules 批量创建标签保护规则
func (t *Tag) CreateGitProtectedTagRules(req *CreateGitProtectedTagRulesReq) (resp CreateGitProtectedTagRulesResp, err error) {
	err = poster.Post(t.c, "CreateGitProtectedTagRules", req, &resp)
	return
}
