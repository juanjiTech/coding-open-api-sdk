package artifact

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeArtifactVersionFileListReq struct {
	Action         string `json:"Action"`         // DescribeArtifactVersionFileList
	Package        string `json:"Package"`        // Pro
	PackageVersion string `json:"PackageVersion"` // latest
	ProjectID      int64  `json:"ProjectId"`      // 450
	Repository     string `json:"Repository"`     // generic
}

type DescribeArtifactVersionFileListResp struct {
	Response struct {
		InstanceSet []struct {
			Name string  `json:"Name"` // LocalTest-2.0-sources.jar
			Size float64 `json:"Size"` // 1.04
		} `json:"InstanceSet"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeArtifactVersionFileList 查询制品版本可下载文件列表
func (a *Artifact) DescribeArtifactVersionFileList(req *DescribeArtifactVersionFileListReq) (resp DescribeArtifactVersionFileListResp, err error) {
	err = poster.Post(a.c, "DescribeArtifactVersionFileList", req, &resp)
	return
}
