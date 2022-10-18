package artifact

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ReleaseArtifactVersionReq struct {
	Action         string `json:"Action"`         // ReleaseArtifactVersion
	Package        string `json:"Package"`        // Pro
	PackageVersion string `json:"PackageVersion"` // latest
	ProjectID      int64  `json:"ProjectId"`      // 450
	Repository     string `json:"Repository"`     // generic
}

type ReleaseArtifactVersionResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ReleaseArtifactVersion 发布制品版本
func (a *Artifact) ReleaseArtifactVersion(req *ReleaseArtifactVersionReq) (resp ReleaseArtifactVersionResp, err error) {
	err = poster.Post(a.c, "ReleaseArtifactVersion", req, &resp)
	return
}
