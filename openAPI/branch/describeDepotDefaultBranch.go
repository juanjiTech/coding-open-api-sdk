package branch

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeDepotDefaultBranchReq struct {
	Action    string `json:"Action"`    // DescribeDepotDefaultBranch
	DepotPath string `json:"DepotPath"` // codingcorp/project-d/depot-d
}

type DescribeDepotDefaultBranchResp struct {
	Response struct {
		BranchName string `json:"BranchName"` // master
		RequestID  string `json:"RequestId"`  // 20b379d2-5ad8-4b70-ac63-e6cfbc4b354a
	} `json:"Response"`
}

// DescribeDepotDefaultBranch 查询仓库的默认分支
func (b *Branch) DescribeDepotDefaultBranch(req *DescribeDepotDefaultBranchReq) (resp DescribeDepotDefaultBranchResp, err error) {
	err = poster.Post(b.c, "DescribeDepotDefaultBranch", req, &resp)
	return
}
