package artifact

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type ModifyArtifactPropertiesReq struct {
	Action         string `json:"Action"`         // ModifyArtifactProperties
	Package        string `json:"Package"`        // Pro
	PackageVersion string `json:"PackageVersion"` // latest
	ProjectID      int64  `json:"ProjectId"`      // 450
	PropertySet    []struct {
		Name  string `json:"Name"`  // name1
		Value string `json:"Value"` // newValue
	} `json:"PropertySet"`
	Repository string `json:"Repository"` // generic
}

type ModifyArtifactPropertiesResp struct {
	Response struct {
		RequestID int64 `json:"RequestId,string"` // 1
	} `json:"Response"`
}

// ModifyArtifactProperties 修改制品属性
func (a *Artifact) ModifyArtifactProperties(req *ModifyArtifactPropertiesReq) (resp ModifyArtifactPropertiesResp, err error) {
	err = poster.Post(a.c, "ModifyArtifactProperties", req, &resp)
	return
}
