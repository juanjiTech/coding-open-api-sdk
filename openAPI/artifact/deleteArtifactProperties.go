package artifact

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DeleteArtifactPropertiesReq struct {
	Action          string   `json:"Action"`          // DeleteArtifactProperties
	Package         string   `json:"Package"`         // Pro
	PackageVersion  string   `json:"PackageVersion"`  // latest
	ProjectID       int64    `json:"ProjectId"`       // 450
	PropertyNameSet []string `json:"PropertyNameSet"` // name1
	Repository      string   `json:"Repository"`      // generic
}

type DeleteArtifactPropertiesResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// DeleteArtifactProperties 删除制品属性
func (a *Artifact) DeleteArtifactProperties(req *DeleteArtifactPropertiesReq) (resp DeleteArtifactPropertiesResp, err error) {
	err = poster.Post(a.c, "DeleteArtifactProperties", req, &resp)
	return
}
