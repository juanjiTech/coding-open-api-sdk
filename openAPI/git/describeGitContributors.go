package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeGitContributorsReq struct {
	Action  string `json:"Action"`  // DescribeGitContributors
	DepotID int64  `json:"DepotId"` // 2
	Ref     string `json:"Ref"`     // dev
}

type DescribeGitContributorsResp struct {
	Response struct {
		Contributors []struct {
			Commits int64  `json:"Commits"` // 1
			Email   string `json:"Email"`   // coding@coding.com
			Name    string `json:"Name"`    // coding
		} `json:"Contributors"`
		RequestID string `json:"RequestId"` // dd604373-7f1b-56c6-54e6-c3653eaee11b
	} `json:"Response"`
}

// DescribeGitContributors 查询 git 仓库的贡献者
func (g *Git) DescribeGitContributors(req *DescribeGitContributorsReq) (resp DescribeGitContributorsResp, err error) {
	err = poster.Post(g.c, "DescribeGitContributors", req, &resp)
	return
}
