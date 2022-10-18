package protectedBranch

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateProtectedBranchMemberReq struct {
	Action     string `json:"Action"`     // CreateProtectedBranchMember
	BranchName string `json:"BranchName"` // master
	DepotID    int64  `json:"DepotId"`    // 5001
	GlobalKey  string `json:"GlobalKey"`  // GlobalKeyDemo
}

type CreateProtectedBranchMemberResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 6f2811dc-dd28-f099-86cc-f15dbeb8fd0e
	} `json:"Response"`
}

// CreateProtectedBranchMember 添加保护分支成员
func (p *ProtectedBranch) CreateProtectedBranchMember(req *CreateProtectedBranchMemberReq) (resp CreateProtectedBranchMemberResp, err error) {
	err = poster.Post(p.c, "CreateProtectedBranchMember", req, &resp)
	return
}
