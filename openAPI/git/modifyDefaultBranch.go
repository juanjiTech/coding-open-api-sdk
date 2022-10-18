package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyDefaultBranchReq struct {
	Action     string `json:"Action"`     // ModifyDefaultBranch
	BranchName string `json:"BranchName"` // release
	DepotID    int64  `json:"DepotId"`    // 1
	UserID     int64  `json:"UserId"`     // 1
}

type ModifyDefaultBranchResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // 3732053f-a016-d452-a8f6-bb80cf8d74ed
	} `json:"Response"`
}

// ModifyDefaultBranch 修改仓库默认分支
func (g *Git) ModifyDefaultBranch(req *ModifyDefaultBranchReq) (resp ModifyDefaultBranchResp, err error) {
	err = poster.Post(g.c, "ModifyDefaultBranch", req, &resp)
	return
}
