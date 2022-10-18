package tag

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeGitTagsByBranchReq struct {
	Action  string `json:"Action"`  // DescribeGitTagsByBranch
	Branch  string `json:"Branch"`  // master
	DepotID int64  `json:"DepotId"` // 46
}

type DescribeGitTagsByBranchResp struct {
	Response struct {
		RequestID string   `json:"RequestId"` // 78b0fa3f-c793-ebb1-822d-82cbc83848e6
		Tags      []string `json:"Tags"`      // v0.1.0
	} `json:"Response"`
}

// DescribeGitTagsByBranch 根据分支获取标签列表
func (t *Tag) DescribeGitTagsByBranch(req *DescribeGitTagsByBranchReq) (resp DescribeGitTagsByBranchResp, err error) {
	err = poster.Post(t.c, "DescribeGitTagsByBranch", req, &resp)
	return
}
