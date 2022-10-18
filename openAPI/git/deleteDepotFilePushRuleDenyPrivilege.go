package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteDepotFilePushRuleDenyPrivilegeReq struct {
	Action         string `json:"Action"`         // DeleteDepotFilePushRuleDenyPrivilege
	DepotPath      string `json:"DepotPath"`      // codingcorp/project-d/depot-1
	FilePushRuleID int64  `json:"FilePushRuleId"` // 5
	IsRole         bool   `json:"IsRole"`         // false
	IsUser         bool   `json:"IsUser"`         // true
	UserGlobalKey  string `json:"UserGlobalKey"`  // coding-coding
}

type DeleteDepotFilePushRuleDenyPrivilegeResp struct {
	Action         string `json:"Action"`         // DeleteDepotFilePushRuleDenyPrivilege
	DepotPath      string `json:"DepotPath"`      // codingcorp/project-d/depot-1
	FilePushRuleID int64  `json:"FilePushRuleId"` // 5
	IsRole         bool   `json:"IsRole"`         // true
	IsUser         bool   `json:"IsUser"`         // false
	RoleID         int64  `json:"RoleId"`         // 4
}

// DeleteDepotFilePushRuleDenyPrivilege 删除 git 仓库特权者文件推送权限
func (g *Git) DeleteDepotFilePushRuleDenyPrivilege(req *DeleteDepotFilePushRuleDenyPrivilegeReq) (resp DeleteDepotFilePushRuleDenyPrivilegeResp, err error) {
	err = poster.Post(g.c, "DeleteDepotFilePushRuleDenyPrivilege", req, &resp)
	return
}
