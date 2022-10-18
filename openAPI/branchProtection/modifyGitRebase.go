package branchProtection

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyGitRebaseReq struct {
	Action         string `json:"Action"`         // ModifyGitRebase
	BaseBranchName string `json:"BaseBranchName"` // master
	DepotID        int64  `json:"DepotId"`        // 1
	SrcBranchName  string `json:"SrcBranchName"`  // master-patch-1
}

type ModifyGitRebaseResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // c96a77a6-e366-f8b8-9d96-96d413d55414
	} `json:"Response"`
}

// ModifyGitRebase git 变基操作
func (b *BranchProtection) ModifyGitRebase(req *ModifyGitRebaseReq) (resp ModifyGitRebaseResp, err error) {
	err = poster.Post(b.c, "ModifyGitRebase", req, &resp)
	return
}
