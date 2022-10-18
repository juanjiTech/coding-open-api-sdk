package git

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeGitFilesReq struct {
	Action  string `json:"Action"`  // DescribeGitFiles
	DepotID int64  `json:"DepotId"` // 809883
	Path    string `json:"Path"`    // src
	Ref     string `json:"Ref"`     // master
}

type DescribeGitFilesResp struct {
	Response struct {
		Items []struct {
			Mode string `json:"Mode"` // tree
			Name string `json:"Name"` // main
			Path string `json:"Path"` // src/main
			Sha  string `json:"Sha"`  // 2346422600241eda05736401950d65cfd22a3cd
		} `json:"Items"`
		RequestID string `json:"RequestId"` // 14bba471-18d8-1c7e-5cee-a7454585bbd2
	} `json:"Response"`
}

// DescribeGitFiles 查询仓库目录下文件和文件夹名字
func (g *Git) DescribeGitFiles(req *DescribeGitFilesReq) (resp DescribeGitFilesResp, err error) {
	err = poster.Post(g.c, "DescribeGitFiles", req, &resp)
	return
}
