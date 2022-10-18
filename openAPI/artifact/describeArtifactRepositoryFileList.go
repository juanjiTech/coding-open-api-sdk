package artifact

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeArtifactRepositoryFileListReq struct {
	Action     string `json:"Action"`     // DescribeArtifactRepositoryFileList
	PageSize   int64  `json:"PageSize"`   // 1000
	Project    string `json:"Project"`    // testing
	Repository string `json:"Repository"` // maven
}

type DescribeArtifactRepositoryFileListResp struct {
	Response struct {
		Data struct {
			ContinuationToken string `json:"ContinuationToken"` // 5e739c1e1f8bdb20cdb0db2ab2215903
			InstanceSet       []struct {
				ArtifactType int64  `json:"ArtifactType"` // 3
				DownloadURL  string `json:"DownloadUrl"`  // https://coding.net/repository/project/maven/pkg/vers/folder/object-1.jar
				Host         string `json:"Host"`         // coding.net
				Path         string `json:"Path"`         // folder/object-1.jpg
				Project      string `json:"Project"`      // project
				Repository   string `json:"Repository"`   // maven
			} `json:"InstanceSet"`
		} `json:"Data"`
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DescribeArtifactRepositoryFileList 查询制品仓库可下载文件列表
func (a *Artifact) DescribeArtifactRepositoryFileList(req *DescribeArtifactRepositoryFileListReq) (resp DescribeArtifactRepositoryFileListResp, err error) {
	err = poster.Post(a.c, "DescribeArtifactRepositoryFileList", req, &resp)
	return
}
