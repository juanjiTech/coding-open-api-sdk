package branch

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateGitBranchReq struct {
	Action     string `json:"Action"`     // CreateGitBranch
	BranchName string `json:"BranchName"` // branch-demo
	DepotID    int64  `json:"DepotId"`    // 1001
	StartPoint string `json:"StartPoint"` // master
}

type CreateGitBranchResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // be249809-0852-bd61-1420-c7d66e5038a9
	} `json:"Response"`
}

// CreateGitBranch 创建分支
func (b *Branch) CreateGitBranch(req *CreateGitBranchReq) (resp CreateGitBranchResp, err error) {
	err = poster.Post(b.c, "CreateGitBranch", req, &resp)
	return
}
