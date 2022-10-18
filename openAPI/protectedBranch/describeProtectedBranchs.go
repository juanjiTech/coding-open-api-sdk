package protectedBranch

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeProtectedBranchsReq struct {
	Action  string `json:"Action"`  // DescribeProtectedBranchs
	DepotID int64  `json:"DepotId"` // 5001
}

type DescribeProtectedBranchsResp struct {
	Response struct {
		ProtectedBranchs []struct {
			CommitDate    int64  `json:"CommitDate"`    // 1.602668567e+12
			DenyForcePush bool   `json:"DenyForcePush"` // true
			ForceSquash   bool   `json:"ForceSquash"`   // false
			Name          string `json:"Name"`          // master
			StatusCheck   bool   `json:"StatusCheck"`   // false
		} `json:"ProtectedBranchs"`
		RequestID string `json:"RequestId"` // 80ca5636-b61b-2419-15a2-0ca80e8f1acd
	} `json:"Response"`
}

// DescribeProtectedBranchs 查询所有保护分支
func (p *ProtectedBranch) DescribeProtectedBranchs(req *DescribeProtectedBranchsReq) (resp DescribeProtectedBranchsResp, err error) {
	err = poster.Post(p.c, "DescribeProtectedBranchs", req, &resp)
	return
}
