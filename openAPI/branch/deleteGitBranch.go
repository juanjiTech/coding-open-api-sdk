package branch

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteGitBranchReq struct {
	Action     string `json:"Action"`     // DeleteGitBranch
	BranchName string `json:"BranchName"` // branch-demo
	DepotID    int64  `json:"DepotId"`    // 1001
}

type DeleteGitBranchResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 30892813-210d-a1cf-10bd-ca11cb1e271a
	} `json:"Response"`
}

// DeleteGitBranch 删除分支
func (b *Branch) DeleteGitBranch(req *DeleteGitBranchReq) (resp DeleteGitBranchResp, err error) {
	err = poster.Post(b.c, "DeleteGitBranch", req, &resp)
	return
}
