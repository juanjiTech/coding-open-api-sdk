package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeGitTreeReq struct {
	Action      string `json:"Action"`      // DescribeGitTree
	DepotID     int64  `json:"DepotId"`     // 2
	IsRecursive bool   `json:"IsRecursive"` // true
	Path        string `json:"Path"`        // test
	Ref         string `json:"Ref"`         // master
}

type DescribeGitTreeResp struct {
	Response struct {
		RequestID string `json:"RequestId"` // d8cec7e5-8a0e-f384-d6a8-d5fbc9c2588c
		Trees     []struct {
			Mode int64  `json:"Mode,string"` // 040000
			Path string `json:"Path"`        // test
			Sha  string `json:"Sha"`         // 596ffc7d3af398763e14b173afbffaed901a2a3b
			Type string `json:"Type"`        // tree
		} `json:"Trees"`
	} `json:"Response"`
}

// DescribeGitTree 查询仓库tree信息
func (g *Git) DescribeGitTree(req *DescribeGitTreeReq) (resp DescribeGitTreeResp, err error) {
	err = poster.Post(g.c, "DescribeGitTree", req, &resp)
	return
}
