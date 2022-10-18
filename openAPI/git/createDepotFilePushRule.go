package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateDepotFilePushRuleReq struct {
	Action           string `json:"Action"`           // CreateDepotFilePushRule
	DepotPath        string `json:"DepotPath"`        // codingcorp/project-d/depot-1
	IsDenyForAllUser bool   `json:"IsDenyForAllUser"` // true
	Pattern          string `json:"Pattern"`          // *.js
}

type CreateDepotFilePushRuleResp struct {
	Response struct {
		GitFilePushRule []struct {
			FilePushRuleID   int64  `json:"FilePushRuleId"`   // 5
			IsDenyForAllUser bool   `json:"IsDenyForAllUser"` // false
			Pattern          string `json:"Pattern"`          // *.js
			Privileges       []struct {
				IsDeny bool `json:"IsDeny"` // true
				IsRole bool `json:"IsRole"` // true
				IsUser bool `json:"IsUser"` // false
				Role   struct {
					Name   string `json:"Name"`   // 项目管理员
					RoleID int64  `json:"RoleId"` // 4
				} `json:"Role"`
				User struct {
					Avatar    string `json:"Avatar"`    //
					GlobalKey string `json:"GlobalKey"` //
					Name      string `json:"Name"`      //
				} `json:"User"`
			} `json:"Privileges"`
		} `json:"GitFilePushRule"`
		RequestID string `json:"RequestId"` // c2805697-49fd-41cd-99db-e698388706ec
	} `json:"Response"`
}

// CreateDepotFilePushRule 新增 git 仓库文件推送规则
func (g *Git) CreateDepotFilePushRule(req *CreateDepotFilePushRuleReq) (resp CreateDepotFilePushRuleResp, err error) {
	err = poster.Post(g.c, "CreateDepotFilePushRule", req, &resp)
	return
}
