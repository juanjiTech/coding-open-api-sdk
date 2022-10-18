package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateDepotFilePushRulePrivilegeReq struct {
	Action         string `json:"Action"`         // CreateDepotFilePushRulePrivilege
	DepotPath      string `json:"DepotPath"`      // codingcorp/project-d/depot-1
	FilePushRuleID int64  `json:"FilePushRuleId"` // 5
	IsDeny         bool   `json:"IsDeny"`         // true
	IsRole         bool   `json:"IsRole"`         // false
	IsUser         bool   `json:"IsUser"`         // true
	UserGlobalKey  string `json:"UserGlobalKey"`  // coding-coding
}

type CreateDepotFilePushRulePrivilegeResp struct {
	Action         string `json:"Action"`         // CreateDepotFilePushRulePrivilege
	DepotPath      string `json:"DepotPath"`      // codingcorp/project-d/depot-1
	FilePushRuleID int64  `json:"FilePushRuleId"` // 5
	IsDeny         bool   `json:"IsDeny"`         // true
	IsRole         bool   `json:"IsRole"`         // true
	IsUser         bool   `json:"IsUser"`         // false
	RoleID         int64  `json:"RoleId"`         // 4
}

// CreateDepotFilePushRulePrivilege 新增 git 仓库文件推送规则特权者
func (g *Git) CreateDepotFilePushRulePrivilege(req *CreateDepotFilePushRulePrivilegeReq) (resp CreateDepotFilePushRulePrivilegeResp, err error) {
	err = poster.Post(g.c, "CreateDepotFilePushRulePrivilege", req, &resp)
	return
}
