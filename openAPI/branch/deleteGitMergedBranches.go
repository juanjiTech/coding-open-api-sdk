package branch

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteGitMergedBranchesReq struct {
	Action  string `json:"Action"`  // DeleteGitMergedBranches
	DepotID int64  `json:"DepotId"` // 2
}

type DeleteGitMergedBranchesResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 8b048fe5-eb93-e8aa-a239-b864d9abd660
	} `json:"Response"`
}

// DeleteGitMergedBranches 删除已合并到默认分支的分支（此操作不会删除受保护的分支）
func (b *Branch) DeleteGitMergedBranches(req *DeleteGitMergedBranchesReq) (resp DeleteGitMergedBranchesResp, err error) {
	err = poster.Post(b.c, "DeleteGitMergedBranches", req, &resp)
	return
}
