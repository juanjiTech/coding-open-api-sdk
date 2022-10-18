package branch

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeGitBranchReq struct {
	Action     string `json:"Action"`     // DescribeGitBranch
	BranchName string `json:"BranchName"` // master
	DepotID    int64  `json:"DepotId"`    // 1001
}

type DescribeGitBranchResp struct {
	Response struct {
		GitBranch struct {
			BranchName      string `json:"BranchName"`      // master
			IsDefaultBranch bool   `json:"IsDefaultBranch"` // true
			IsProtected     bool   `json:"IsProtected"`     // false
			Sha             string `json:"Sha"`             // bf778a27fa30a889a30af6362ba1f16a48dd58dd
		} `json:"GitBranch"`
		RequestID string `json:"RequestId"` // fd1d65dd-5c9d-df50-e348-7d3bf0f31f53
	} `json:"Response"`
}

// DescribeGitBranch 查询分支信息
func (b *Branch) DescribeGitBranch(req *DescribeGitBranchReq) (resp DescribeGitBranchResp, err error) {
	err = poster.Post(b.c, "DescribeGitBranch", req, &resp)
	return
}
