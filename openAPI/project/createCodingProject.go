package project

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type CreateCodingProjectReq struct {
	Action           string `json:"Action"`           // CreateCodingProject
	CreateSvnLayout  bool   `json:"CreateSvnLayout"`  // false
	Description      string `json:"Description"`      // project
	DisplayName      string `json:"DisplayName"`      // project
	GitReadmeEnabled bool   `json:"GitReadmeEnabled"` // true
	Name             string `json:"Name"`             // project
	ProjectTemplate  string `json:"ProjectTemplate"`  // DEV_OPS
	Shared           int64  `json:"Shared"`           // 1
	VcsType          string `json:"VcsType"`          // git
}

type CreateCodingProjectResp struct {
	Response struct {
		ProjectID int64  `json:"ProjectId"` // 1
		RequestID string `json:"RequestId"` // ae8e2d5f-569b-443e-8c61-440ea3a7562a
	} `json:"Response"`
}

// CreateCodingProject 创建项目
func (p *Project) CreateCodingProject(req *CreateCodingProjectReq) (resp CreateCodingProjectResp, err error) {
	err = poster.Post(p.c, "CreateCodingProject", req, &resp)
	return
}
