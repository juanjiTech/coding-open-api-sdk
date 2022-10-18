package branchProtection

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateBranchProtectionReq struct {
	Action  string `json:"Action"`  // CreateBranchProtection
	DepotID int64  `json:"DepotId"` // 809883
	Rule    string `json:"Rule"`    // dev/*
}

type CreateBranchProtectionResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // c20cbf85-844c-6a7d-8f4d-99400db1e9e1
	} `json:"Response"`
}

// CreateBranchProtection 新建保护分支规则
func (b *BranchProtection) CreateBranchProtection(req *CreateBranchProtectionReq) (resp CreateBranchProtectionResp, err error) {
	err = poster.Post(b.c, "CreateBranchProtection", req, &resp)
	return
}
