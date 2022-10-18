package protectedBranch

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateProtectedBranchReq struct {
	Action     string `json:"Action"`     // CreateProtectedBranch
	BranchName string `json:"BranchName"` // master
	DepotID    int64  `json:"DepotId"`    // 5001
}

type CreateProtectedBranchResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 2cf6e3b6-7fa7-0433-50b1-73e6226ac15b
	} `json:"Response"`
}

// CreateProtectedBranch 设置保护分支
func (p *ProtectedBranch) CreateProtectedBranch(req *CreateProtectedBranchReq) (resp CreateProtectedBranchResp, err error) {
	err = poster.Post(p.c, "CreateProtectedBranch", req, &resp)
	return
}
