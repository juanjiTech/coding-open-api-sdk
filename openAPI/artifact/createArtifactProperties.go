package artifact

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateArtifactPropertiesReq struct {
	Action         string `json:"Action"`         // CreateArtifactProperties
	Package        string `json:"Package"`        // Pro
	PackageVersion string `json:"PackageVersion"` // latest
	ProjectID      int64  `json:"ProjectId"`      // 450
	PropertySet    []struct {
		Name  string `json:"Name"`  // name1
		Value string `json:"Value"` // value1
	} `json:"PropertySet"`
	Repository string `json:"Repository"` // generic
}

type CreateArtifactPropertiesResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// CreateArtifactProperties 新增制品属性
func (a *Artifact) CreateArtifactProperties(req *CreateArtifactPropertiesReq) (resp CreateArtifactPropertiesResp, err error) {
	err = poster.Post(a.c, "CreateArtifactProperties", req, &resp)
	return
}
