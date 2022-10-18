package branchProtection

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateBranchProtectionMemberReq struct {
	Action             string `json:"Action"`             // CreateBranchProtectionMember
	AllowPush          bool   `json:"AllowPush"`          // true
	BranchProtectionID int64  `json:"BranchProtectionId"` // 28234
	DepotID            int64  `json:"DepotId"`            // 809883
	UserGlobalKey      string `json:"UserGlobalKey"`      // coding
}

type CreateBranchProtectionMemberResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 51e123c8-48af-72f9-ac98-dc32f393355f
	} `json:"Response"`
}

// CreateBranchProtectionMember 新建保护分支规则的成员
func (b *BranchProtection) CreateBranchProtectionMember(req *CreateBranchProtectionMemberReq) (resp CreateBranchProtectionMemberResp, err error) {
	err = poster.Post(b.c, "CreateBranchProtectionMember", req, &resp)
	return
}
